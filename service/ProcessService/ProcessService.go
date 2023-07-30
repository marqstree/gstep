package ProcessService

import (
	"github.com/jinzhu/copier"
	"github.com/marqstree/gstep/dao/TemplateDao"
	"github.com/marqstree/gstep/enum/AuditMethod"
	"github.com/marqstree/gstep/enum/ProcessState"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
)

func Start(dto *dto.ProcessStartDto, tx *gorm.DB) int {
	process := entity.Process{}
	copier.Copy(process, dto)

	template := TemplateDao.GetLatestVersionByGroupId(dto.TemplateGroupId)
	if nil == template {
		panic(ServerError.New("无效的模板"))
	}
	process.TemplateId = template.Id
	process.StartUserId = dto.StartUserId
	process.State = ProcessState.STARTED.Code
	dao.SaveOrUpdate(&process, tx)

	task := entity.Task{}
	task.ProcessId = process.Id
	task.Form = template.RootStep.Form
	task.StepId = template.RootStep.Id
	task.Title = template.RootStep.Title
	task.Category = template.RootStep.Category
	if len(template.RootStep.AuditMethod) == 0 {
		task.AuditMethod = AuditMethod.OR.Code
	}

	dao.SaveOrUpdate(&task, tx)

	return process.Id
}
