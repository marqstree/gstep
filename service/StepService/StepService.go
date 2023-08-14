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

	if pParentStep.NextStep.Id == stepId {
		return pParentStep.NextStep
	}

	for _, v := range pParentStep.BranchSteps {
		if v.Id == stepId {
			return v
		}
	}

	aStep := FindStep(pParentStep.NextStep, stepId)
	if nil != aStep {
		return aStep
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

	if pParentStep.NextStep.Id == beginStepId {
		return pParentStep
	}

	for _, v := range pParentStep.BranchSteps {
		if v.Id == beginStepId {
			return pParentStep
		}

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

func FindPrevAuditStep(pRootStep *entity.Step, beginStepId int) *entity.Step {
	fromStepId := beginStepId
	for {
		pPrevStep := FindPrevStep(pRootStep, fromStepId)
		if nil == pPrevStep {
			return nil
		}
		if StepCat.IsContainAudit(pPrevStep.Category) {
			return pPrevStep
		}

		fromStepId = pPrevStep.Id
	}
}

func FindPrevAuditSteps(pRootStep *entity.Step, beginStepId int) []*entity.Step {
	auditpSteps := []*entity.Step{}
	fromStepId := beginStepId
	for {
		pPrevStep := FindPrevAuditStep(pRootStep, fromStepId)
		if nil == pPrevStep {
			return []*entity.Step{}
		}
		if StepCat.IsContainAudit(pPrevStep.Category) {
			auditpSteps = append(auditpSteps, pPrevStep)
		}

		fromStepId = pPrevStep.Id
	}
}

func FindPrevEndAuditSteps(pRootStep *entity.Step, beginStepId int, endStepId int) []*entity.Step {
	auditpSteps := []*entity.Step{}
	fromStepId := beginStepId
	for {
		pPrevStep := FindPrevAuditStep(pRootStep, fromStepId)

		if nil == pPrevStep {
			return []*entity.Step{}
		}

		if StepCat.IsContainAudit(pPrevStep.Category) {
			auditpSteps = append(auditpSteps, pPrevStep)
		}

		if endStepId == pPrevStep.Id {
			return auditpSteps
		}

		fromStepId = pPrevStep.Id
	}
}

func CheckCandidate(userId string, pRootStep *entity.Step, stepId int, tx *gorm.DB) {
	pStep := FindStep(pRootStep, stepId)
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
