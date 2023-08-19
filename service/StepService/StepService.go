package StepService

import (
	"github.com/marqstree/gstep/dao/DepartmentDao"
	"github.com/marqstree/gstep/dao/ProcessDao"
	"github.com/marqstree/gstep/dao/UserDao"
	"github.com/marqstree/gstep/enum/CandidateCat"
	"github.com/marqstree/gstep/enum/StepCat"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
)

func GetStepByProcess(pPrcess *entity.Process, stepId int, tx *gorm.DB) *entity.Step {
	processVo := ProcessDao.ToVo(pPrcess, tx)
	pStep := FindStep(&processVo.Template.RootStep, stepId)

	if nil == pStep {
		panic(ServerError.New("无效的步骤id"))
	}

	return pStep
}

func GetStepByTemplateId(templateId int, stepId int, tx *gorm.DB) *entity.Step {
	template := dao.CheckById[entity.Template](templateId, tx)

	pStep := FindStep(&template.RootStep, stepId)
	return pStep
}

func FindStep(pParentStep *entity.Step, stepId int) *entity.Step {
	if nil == pParentStep {
		return nil
	}

	if pParentStep.Id == stepId {
		return pParentStep
	}

	if nil != pParentStep.NextStep && pParentStep.NextStep.Id == stepId {
		return pParentStep.NextStep
	}

	for _, v := range pParentStep.BranchSteps {
		if v.Id == stepId {
			return v
		}
	}

	if nil != pParentStep {
		aStep := FindStep(pParentStep.NextStep, stepId)
		if nil != aStep {
			return aStep
		}
	}

	for _, v := range pParentStep.BranchSteps {
		pFindOne := FindStep(v, stepId)
		if nil != pFindOne {
			return pFindOne
		}
	}

	return nil
}

func FindPrevStep(pParentStep *entity.Step, beginStepId int) *entity.Step {
	if nil == pParentStep {
		return nil
	}

	if nil != pParentStep.NextStep && pParentStep.NextStep.Id == beginStepId {
		return pParentStep
	}

	for _, v := range pParentStep.BranchSteps {
		if v.Id == beginStepId {
			return pParentStep
		}
	}

	if nil != pParentStep.NextStep {
		pNextStep := FindPrevStep(pParentStep.NextStep, beginStepId)
		if nil != pNextStep {
			return pNextStep
		}
	}

	for _, v := range pParentStep.BranchSteps {
		pFindOne := FindPrevStep(v, beginStepId)
		if nil != pFindOne {
			return pFindOne
		}
	}

	return nil
}

func FindPrevBranchStepWithNextStep(pRootStep *entity.Step, beginStepId int) *entity.Step {
	pPrevStep := FindPrevStep(pRootStep, beginStepId)

	if nil == pPrevStep {
		return nil
	}

	if pPrevStep.Category == StepCat.BRANCH.Code && nil != pPrevStep.NextStep && pPrevStep.NextStep.Id != 0 {
		return pPrevStep
	}

	pPrevPrevStep := FindPrevBranchStepWithNextStep(pRootStep, pPrevStep.Id)
	return pPrevPrevStep
}

// 查找前一个审核步骤
func FindPrevAuditStep(pRootStep *entity.Step, beginStepId int) *entity.Step {
	fromStepId := beginStepId
	for {
		pPrevStep := FindPrevStep(pRootStep, fromStepId)
		if nil == pPrevStep || pPrevStep.Id == 0 {
			return nil
		}
		if StepCat.IsContainAudit(pPrevStep.Category) {
			return pPrevStep
		}

		fromStepId = pPrevStep.Id
	}
}

// 查询前面所有审核步骤列表
func FindPrevAuditSteps(pRootStep *entity.Step, beginStepId int) []entity.Step {
	auditSteps := []entity.Step{}
	fromStepId := beginStepId
	for {
		pPrevStep := FindPrevAuditStep(pRootStep, fromStepId)
		if nil == pPrevStep {
			return auditSteps
		}
		if StepCat.IsContainAudit(pPrevStep.Category) {
			auditSteps = append(auditSteps, *pPrevStep)
		}

		fromStepId = pPrevStep.Id
	}
}

// 查询之前到指定步骤的审核列表
func FindPrevAuditStepsByEndId(pRootStep *entity.Step, beginStepId int, endStepId int) []entity.Step {
	auditpSteps := []entity.Step{}
	fromStepId := beginStepId
	for {
		pPrevStep := FindPrevAuditStep(pRootStep, fromStepId)

		if nil == pPrevStep {
			return auditpSteps
		}

		if StepCat.IsContainAudit(pPrevStep.Category) {
			auditpSteps = append(auditpSteps, *pPrevStep)
		}

		if endStepId == pPrevStep.Id {
			return auditpSteps
		}

		fromStepId = pPrevStep.Id
	}
}

// 检查指定步骤的候选人
func CheckCandidate(userId string, pRootStep *entity.Step, stepId int, tx *gorm.DB) {
	pStep := FindStep(pRootStep, stepId)
	if nil == pStep {
		panic(ServerError.New("找不到流程步骤"))
	}
	if len(pStep.Candidates) == 0 {
		return
	}

	for _, v := range pStep.Candidates {
		if v.Category == CandidateCat.USER.Code {
			if userId == v.Value {
				return
			}
		} else if v.Category == CandidateCat.DEPARTMENT.Code {
			departments := DepartmentDao.GetGrandsonDepartments(v.Value, tx)
			isIn := UserDao.IsUserInDepartments(userId, departments, tx)
			if isIn {
				return
			}
		}
	}

	panic(ServerError.New("流程提交人不在候选人列表中"))
}

// 候选人条数
func CandidateCount(taskId int, tx *gorm.DB) int {
	pTask := dao.CheckById[entity.Task](taskId, tx)
	pProcess := dao.CheckById[entity.Process](pTask.ProcessId, tx)
	pTemplate := dao.CheckById[entity.Template](pProcess.TemplateId, tx)
	pStep := FindStep(&pTemplate.RootStep, pTask.StepId)
	return len(pStep.Candidates)
}
