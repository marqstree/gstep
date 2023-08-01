package TaskService

import (
	"github.com/marqstree/gstep/config"
	"github.com/marqstree/gstep/dao/TaskAssigneeDao"
	"github.com/marqstree/gstep/dao/TaskCandidateDao"
	"github.com/marqstree/gstep/dao/TaskDao"
	"github.com/marqstree/gstep/enum/AuditMethod"
	"github.com/marqstree/gstep/enum/ProcessState"
	"github.com/marqstree/gstep/enum/StepCat"
	"github.com/marqstree/gstep/enum/TaskState"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/StepService"
	"github.com/marqstree/gstep/util/CollectionUtil"
	"github.com/marqstree/gstep/util/ExpressionUtil"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestUtil"
	"gorm.io/gorm"
	"log"
)

func Pass(pDto *dto.TaskPassDto, tx *gorm.DB) {
	pTask := dao.CheckById[entity.Task](pDto.TaskId, tx)
	pProcess := dao.CheckById[entity.Process](pTask.ProcessId, tx)

	//保存任务提交人
	assignee := entity.TaskAssignee{}
	assignee.TaskId = pTask.Id
	assignee.UserId = pDto.UserId
	assignee.State = TaskState.PASS.Code
	dao.SaveOrUpdate(&assignee, tx)

	//保存任务表单
	pTask.Form = pDto.Form
	//更新任务状态
	if CanPass(pTask, pProcess, tx) {
		pTask.State = TaskState.PASS.Code
	}
	dao.SaveOrUpdate(pTask, tx)

	//启动下一步
	if pTask.State == TaskState.PASS.Code {
		pNextStep := GetNextStep(pTask.StepId, pProcess, &pTask.Form, tx)
		if nil != pNextStep {
			NewTaskByStep(pNextStep, pProcess, tx)
		} else {
			FinishPassProcess(pProcess, tx)
		}
	}
}

func Refuse(pDto *dto.TaskRefuseDto, tx *gorm.DB) {
	pTask := dao.CheckById[entity.Task](pDto.TaskId, tx)
	pProcess := dao.CheckById[entity.Process](pTask.ProcessId, tx)

	//保存任务提交人
	assignee := entity.TaskAssignee{}
	assignee.TaskId = pTask.Id
	assignee.UserId = pDto.UserId
	assignee.State = TaskState.REFUSE.Code
	dao.SaveOrUpdate(&assignee, tx)

	//保存任务表单
	pTask.Form = pDto.Form
	//更新任务状态
	pTask.State = TaskState.REFUSE.Code
	dao.SaveOrUpdate(pTask, tx)

	//撤销上一步
	pPrevSteps := StepService.FindPrevEndAuditSteps(&pProcess.RootStep, pTask.StepId, pDto.PrevStepId)
	for i, v := range pPrevSteps {
		pPrevTask := TaskDao.QueryTaskByStepId(v.Id, pProcess.Id, tx)

		if i == len(pPrevSteps)-1 {
			pPrevTask.State = TaskState.STARTED.Code
		} else {
			pPrevTask.State = TaskState.WITHDRAW.Code
		}
		dao.SaveOrUpdate(pPrevTask, tx)
	}
}

func Cease(pDto *dto.TaskCeaseDto, tx *gorm.DB) {
	pTask := dao.CheckById[entity.Task](pDto.TaskId, tx)
	pProcess := dao.CheckById[entity.Process](pTask.ProcessId, tx)

	//保存任务提交人
	assignee := entity.TaskAssignee{}
	assignee.TaskId = pTask.Id
	assignee.UserId = pDto.UserId
	assignee.State = TaskState.REFUSE.Code
	dao.SaveOrUpdate(&assignee, tx)

	//保存任务表单
	pTask.Form = pDto.Form
	//更新任务状态
	pTask.State = TaskState.REFUSE.Code
	dao.SaveOrUpdate(pTask, tx)

	//撤销上一步
	pProcess.State = ProcessState.FINISH_REFUSE.Code
	dao.SaveOrUpdate(pProcess, tx)
}

func CanPass(pTask *entity.Task, pProcess *entity.Process, tx *gorm.DB) bool {
	if pTask.Category == StepCat.CONDITION.Code {
		pStep := StepService.GetStepByProcess(pProcess, pTask.StepId)
		exp := ExpressionUtil.Template2jsExpression(pStep.Expression, &pTask.Form)
		isPass := ExpressionUtil.RunJsExpression(exp)
		return isPass
	} else {
		passCount := TaskAssigneeDao.PassCount(pTask.Id, tx)

		if pTask.AuditMethod == AuditMethod.OR.Code {
			return passCount > 0
		} else {
			candidateCount := TaskCandidateDao.CandidateCount(pTask.Id, tx)
			return passCount >= candidateCount
		}
	}
}

func GetNextStep(currentStepId int, pProcess *entity.Process, pForm *map[string]any, tx *gorm.DB) *entity.Step {
	pStep := StepService.GetStepByProcess(pProcess, currentStepId)
	if len(pStep.NextSteps) == 0 {
		return nil
	}

	for _, v := range pStep.NextSteps {
		if v.Category == StepCat.CONDITION.Code {
			isPass := ExpressionUtil.ExecuteExpression(v.Expression, pForm)
			if isPass {
				nextStep := GetNextStep(v.Id, pProcess, pForm, tx)
				return nextStep
			}
		} else {
			return &v
		}
	}

	return nil
}

func NewTaskByStep(pStep *entity.Step, pProcess *entity.Process, tx *gorm.DB) *entity.Task {
	if nil == pStep {
		panic(ServerError.New("流程步骤不能为空"))
	}

	if pStep.Category == StepCat.CONDITION.Code {
		panic(ServerError.New("无法用条件步骤创建流程任务"))
	} else if pStep.Category != StepCat.START.Code {
		if len(pStep.Candidates) == 0 {
			panic(ServerError.New("不能用无候选人的流程步骤创建流程任务"))
		}
	}

	task := entity.Task{}
	task.ProcessId = pProcess.Id
	task.StepId = pStep.Id
	task.Title = pStep.Title
	task.Category = pStep.Category
	task.AuditMethod = pStep.AuditMethod
	task.State = TaskState.STARTED.Code
	dao.SaveOrUpdate(&task, tx)

	//创建任务的候选人
	if pStep.Category != StepCat.START.Code {
		for _, v := range pStep.Candidates {
			candidate := entity.TaskCandidate{}
			candidate.TaskId = task.Id
			candidate.UserId = v
			dao.SaveOrUpdate(&candidate, tx)

			assignee := entity.TaskAssignee{}
			assignee.TaskId = task.Id
			assignee.UserId = v
			assignee.State = TaskState.STARTED.Code
			dao.SaveOrUpdate(&assignee, tx)
		}
	}

	return &task
}

func NewStartTask(pProcess *entity.Process, startUserId string, tx *gorm.DB) *entity.Task {
	//创建启动任务
	task := entity.Task{}
	task.ProcessId = pProcess.Id
	task.Form = pProcess.RootStep.Form
	task.StepId = pProcess.RootStep.Id
	task.Title = pProcess.RootStep.Title
	task.Category = StepCat.START.Code
	task.State = TaskState.STARTED.Code
	if len(pProcess.RootStep.AuditMethod) == 0 {
		task.AuditMethod = AuditMethod.OR.Code
	}
	dao.SaveOrUpdate(&task, tx)

	//创建启动任务的候选人
	candidate := entity.TaskCandidate{}
	candidate.TaskId = task.Id
	candidate.UserId = startUserId
	dao.SaveOrUpdate(&candidate, tx)

	assignee := entity.TaskAssignee{}
	assignee.TaskId = task.Id
	assignee.UserId = startUserId
	assignee.State = TaskState.STARTED.Code
	dao.SaveOrUpdate(&assignee, tx)

	return &task
}

// 审核通过流程
func FinishPassProcess(pProcess *entity.Process, tx *gorm.DB) {
	pProcess.State = ProcessState.FINISH_PASS.Code
	dao.SaveOrUpdate(pProcess, tx)
}

// 调用流程任务状态变更通知回调接口
func NotifyTaskStateChange(pTask *entity.Task) {
	url := config.Config.Notify.TaskStateChange
	if len(url) == 0 {
		return
	}

	m, err := CollectionUtil.Obj2map(*pTask)
	if nil != err {
		return
	}
	result := AjaxJson.AjaxJson{}
	RequestUtil.PostJson(url, m, &result)
	log.Println("接收通知服务端返回: %v", result)
}

// 检查流程任务是否重复审核
func CheckTaskCanChange(pTask *entity.Task) {
	if pTask.State == TaskState.PASS.Code {
		panic(ServerError.New("重复审核任务"))
	}
}
