package TaskService

import (
	"github.com/marqstree/gstep/config"
	"github.com/marqstree/gstep/dao/ProcessDao"
	"github.com/marqstree/gstep/dao/TaskAssigneeDao"
	"github.com/marqstree/gstep/dao/TaskCandidateDao"
	"github.com/marqstree/gstep/dao/TaskDao"
	"github.com/marqstree/gstep/dao/UserDao"
	"github.com/marqstree/gstep/enum/AuditMethodCat"
	"github.com/marqstree/gstep/enum/CandidateCat"
	"github.com/marqstree/gstep/enum/ProcessState"
	"github.com/marqstree/gstep/enum/StepCat"
	"github.com/marqstree/gstep/enum/TaskState"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/StepService"
	"github.com/marqstree/gstep/util/CollectionUtil"
	"github.com/marqstree/gstep/util/ExpressionUtil"
	"github.com/marqstree/gstep/util/LocalTime"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestUtil"
	"gorm.io/gorm"
	"log"
	"time"
)

// 审核通过
func Pass(pDto *dto.TaskPassDto, tx *gorm.DB) {
	pTask := dao.CheckById[entity.Task](pDto.TaskId, tx)
	pProcess := dao.CheckById[entity.Process](pTask.ProcessId, tx)
	pTemplate := dao.CheckById[entity.Template](pProcess.TemplateId, tx)
	pStep := StepService.FindStep(&pTemplate.RootStep, pTask.StepId)

	if nil == pStep {
		panic(ServerError.New("无效的流程步骤"))
	}
	if pStep.Category == StepCat.END.Code {
		panic(ServerError.New("结束步骤不用提交"))
	}

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
		pNextStep := GetNextStep(pTask.StepId, pTemplate, &pTask.Form, tx)
		if nil == pNextStep {
			panic(ServerError.New("找不到流程的下一个步骤"))
		}

		//创建新任务
		if pNextStep.Category != StepCat.END.Code {
			NewTaskByStep(pNextStep, pProcess, tx)
		}

		//结束步骤,结束流程
		if pNextStep.Category == StepCat.END.Code {
			FinishPassProcess(pProcess, tx)
		}
	}
}

func Refuse(pDto *dto.TaskRefuseDto, tx *gorm.DB) {
	pTask := dao.CheckById[entity.Task](pDto.TaskId, tx)
	pProcess := dao.CheckById[entity.Process](pTask.ProcessId, tx)
	pProcessVo := ProcessDao.ToVo(pProcess, tx)

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
	pPrevSteps := StepService.FindPrevEndAuditSteps(&pProcessVo.Template.RootStep, pTask.StepId, pDto.PrevStepId)
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

// 终止流程
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

// 判断是否可通过
func CanPass(pTask *entity.Task, pProcess *entity.Process, tx *gorm.DB) bool {
	if pTask.Category == StepCat.CONDITION.Code {
		pStep := StepService.GetStepByProcess(pProcess, pTask.StepId, tx)
		exp := ExpressionUtil.Template2jsExpression(pStep.Expression, &pTask.Form)
		isPass := ExpressionUtil.RunJsExpression(exp)
		return isPass
	} else {
		passCount := TaskAssigneeDao.PassCount(pTask.Id, tx)

		//或签
		if pTask.AuditMethod == AuditMethodCat.OR.Code {
			return passCount > 0
		} else { //会签
			candidateCount := TaskCandidateDao.CandidateCount(pTask.Id, tx)
			return passCount >= candidateCount
		}
	}
}

// 查找流程的下一个步骤
// 碰到条件步骤,取条件步骤的下一个步骤
func GetNextStep(currentStepId int, pTemplate *entity.Template, pForm *map[string]any, tx *gorm.DB) *entity.Step {
	pStep := StepService.FindStep(&pTemplate.RootStep, currentStepId)

	if nil == pStep {
		panic(ServerError.New("找不到流程步骤"))
	}

	//分支步骤,找满足条件的子条件步骤的下一步
	if pStep.Category == StepCat.BRANCH.Code {
		if len(pStep.BranchSteps) < 2 {
			panic(ServerError.New("分支步骤的分支数量小于2"))
		}

		var pDefaultConditionStep = &entity.Step{}
		//先找满足的非默认条件
		for _, v := range pStep.BranchSteps {
			if nil == v {
				panic(ServerError.New("无效流程步骤"))
			}

			if v.Category != StepCat.CONDITION.Code {
				panic(ServerError.New("流程分支的首个步骤不是条件类型步骤"))
			}

			if v.Title == "默认条件" {
				pDefaultConditionStep = v
				continue
			}

			isPass := ExpressionUtil.ExecuteExpression(v.Expression, pForm)
			if isPass {
				nextStep := GetNextStep(v.Id, pTemplate, pForm, tx)
				return nextStep
			}
		}
		//所有条件都不满足,走默认条件步骤
		nextStep := GetNextStep(pDefaultConditionStep.Id, pTemplate, pForm, tx)
		return nextStep
	}

	if nil != pStep.NextStep && pStep.NextStep.Id != 0 {
		return pStep.NextStep
	} else {
		//从父分支步骤开始往前递归查找有下一步的分支步骤
		pPrevBranchStep := StepService.FindPrevBranchStepWithNextStep(&pTemplate.RootStep, pStep.Id)
		if nil != pPrevBranchStep.NextStep {
			return pPrevBranchStep.NextStep
		}
	}

	return nil
}

func NewTaskByStep(pStep *entity.Step, pProcess *entity.Process, tx *gorm.DB) *entity.Task {
	if nil == pStep {
		panic(ServerError.New("流程步骤不能为空"))
	}

	if pStep.Category == StepCat.BRANCH.Code {
		panic(ServerError.New("无法用分支步骤创建流程任务"))
	} else if pStep.Category == StepCat.CONDITION.Code {
		panic(ServerError.New("无法用条件步骤创建流程任务"))
	} else if pStep.Category == StepCat.END.Code {
		panic(ServerError.New("无法用结束步骤创建流程任务"))
	}

	if len(pStep.Candidates) == 0 {
		panic(ServerError.New("不能用无候选人的流程步骤创建流程任务"))
	}

	task := entity.Task{}
	task.ProcessId = pProcess.Id
	task.StepId = pStep.Id
	task.Title = pStep.Title
	task.Category = pStep.Category
	task.AuditMethod = pStep.AuditMethod
	if pStep.Category == StepCat.END.Code {
		task.State = TaskState.PASS.Code
	} else {
		task.State = TaskState.STARTED.Code
	}
	dao.SaveOrUpdate(&task, tx)

	//创建任务的候选人
	if pStep.Category != StepCat.END.Code {
		for _, c := range pStep.Candidates {
			if c.Category == CandidateCat.USER.Code {
				assignee := entity.TaskAssignee{}
				assignee.TaskId = task.Id
				assignee.UserId = c.Value
				assignee.State = TaskState.STARTED.Code
				dao.SaveOrUpdate(&assignee, tx)
			} else if c.Category == CandidateCat.DEPARTMENT.Code {
				users := UserDao.GetGrandsonDepartmentUsers(c.Value, tx)
				for _, user := range users {
					assignee := entity.TaskAssignee{}
					assignee.TaskId = task.Id
					assignee.UserId = user.Id
					assignee.State = TaskState.STARTED.Code
					dao.SaveOrUpdate(&assignee, tx)
				}
			}
		}
	}

	return &task
}

// 创建启动任务
func NewStartTask(pProcess *entity.Process, startUserId string, tx *gorm.DB) *entity.Task {
	//创建启动任务
	task := entity.Task{}
	task.ProcessId = pProcess.Id

	processVo := ProcessDao.ToVo(pProcess, tx)
	rootStep := processVo.Template.RootStep

	//检查流程提交人是否是候选人
	StepService.CheckCandidate(startUserId, &rootStep, task.StepId, tx)

	//创建启动任务
	task.StepId = rootStep.Id
	task.Title = rootStep.Title
	task.Category = rootStep.Category
	task.State = TaskState.PASS.Code
	if len(rootStep.AuditMethod) == 0 {
		task.AuditMethod = AuditMethodCat.OR.Code
	} else {
		task.AuditMethod = rootStep.AuditMethod
	}
	dao.SaveOrUpdate(&task, tx)

	//创建启动任务的候选人
	assignee := entity.TaskAssignee{}
	assignee.TaskId = task.Id
	assignee.UserId = startUserId
	assignee.State = TaskState.PASS.Code
	dao.SaveOrUpdate(&assignee, tx)

	return &task
}

// 审核通过流程
func FinishPassProcess(pProcess *entity.Process, tx *gorm.DB) {
	pProcess.State = ProcessState.FINISH_PASS.Code
	finishTime := LocalTime.LocalTime(time.Now())
	pProcess.FinishedAt = &finishTime
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
