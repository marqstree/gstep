package ProcessService

import (
	"github.com/jinzhu/copier"
	"github.com/marqstree/gstep/dao/TemplateDao"
	"github.com/marqstree/gstep/enum/ProcessState"
	"github.com/marqstree/gstep/enum/StepCat"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/TaskService"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
)

func Start(dto *dto.ProcessStartDto, tx *gorm.DB) int {
	process := entity.Process{}
	copier.Copy(process, dto)

	//创建流程
	pTemplate := TemplateDao.GetLatestVersionByGroupId(dto.TemplateGroupId, tx)
	if nil == pTemplate {
		panic(ServerError.New("无效的模板"))
	}
	process.TemplateId = pTemplate.Id
	process.StartUserId = dto.StartUserId
	process.State = ProcessState.STARTED.Code
	dao.SaveOrUpdate(&process, tx)

	//创建启动任务
	pStartTask := TaskService.NewStartTask(&process, dto.StartUserId, tx)

	//启动下一步
	pNextStep := TaskService.GetNextStep(pStartTask.StepId, pTemplate, &pStartTask.Form, tx)
	if nil == pNextStep {
		panic(ServerError.New("找不到流程启动步骤的下一个步骤"))
	}
	//创建新任务
	if pNextStep.Category != StepCat.END.Code {
		TaskService.NewTaskByStep(pNextStep, &process, tx)
	}
	//结束步骤,结束流程
	if pNextStep.Category == StepCat.END.Code {
		TaskService.FinishPassProcess(&process, tx)
	}

	return process.Id
}
