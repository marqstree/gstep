package StepService

import (
	"github.com/marqstree/gstep/enum/StepCat"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
)

func GetStepByProcess(pPrcess *entity.Process, stepId int) *entity.Step {
	pStep := FindStep(&pPrcess.RootStep, stepId)

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
