package StepService

import (
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
)

func GetStepByProcess(pPrcess *entity.Process, stepId int) *entity.Step {
	pStep := findStep(&pPrcess.RootStep, stepId)

	if nil == pStep {
		panic(ServerError.New("无效的步骤id"))
	}

	return pStep
}

func GetStepByTemplateId(templateId int, stepId int, tx *gorm.DB) *entity.Step {
	template := dao.CheckById[entity.Template](templateId, tx)

	pStep := findStep(&template.RootStep, stepId)
	return pStep
}

func findStep(pRootStep *entity.Step, stepId int) *entity.Step {
	if nil == pRootStep {
		return nil
	}

	if pRootStep.Id == stepId {
		return pRootStep
	}

	if len(pRootStep.NextSteps) == 0 {
		return nil
	}
	
	for _, v := range pRootStep.NextSteps {
		if v.Id == stepId {
			return &v
		}

		pFindOne := findStep(&v, stepId)
		if nil != pFindOne {
			return pFindOne
		}
	}

	return nil
}
