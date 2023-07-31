package TaskService

import (
	"github.com/marqstree/gstep/dao/TaskAssigneeDao"
	"github.com/marqstree/gstep/dao/TaskCandidateDao"
	"github.com/marqstree/gstep/enum/AuditMethod"
	"github.com/marqstree/gstep/enum/ProcessState"
	"github.com/marqstree/gstep/enum/StepCat"
	"github.com/marqstree/gstep/enum/TaskState"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/StepService"
	"github.com/marqstree/gstep/util/ExpressionUtil"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
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

		pNextStep := GetNextStep(pTask.StepId, &pTask.Form, pProcess, tx)
		if nil != pNextStep {
			NewTaskByStep(pNextStep, pProcess, tx)
		} else {
			FinishProcess(pProcess, tx)
		}
	}
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

func GetNextStep(currentStepId int, pForm *map[string]any, pProcess *entity.Process, tx *gorm.DB) *entity.Step {
	pStep := StepService.GetStepByProcess(pProcess, currentStepId)
	if len(pStep.NextSteps) == 0 {
		return nil
	}

	for _, v := range pStep.NextSteps {
		if v.Category == StepCat.CONDITION.Code {
			isPass := ExpressionUtil.ExecuteExpression(v.Expression, pForm)
			if isPass {
				nextStep := GetNextStep(v.Id, pForm, pProcess, tx)
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
	} else if pStep.Category != StepCat.END.Code {
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
	if pStep.Category == StepCat.END.Code {
		task.State = TaskState.PASS.Code
	} else {
		task.State = TaskState.STARTED.Code
	}
	dao.SaveOrUpdate(&task, tx)

	//创建任务的候选人
	if pStep.Category != StepCat.END.Code {
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

func FinishProcess(pProcess *entity.Process, tx *gorm.DB) {
	pProcess.State = ProcessState.FINISHED.Code
	dao.SaveOrUpdate(pProcess, tx)
}
