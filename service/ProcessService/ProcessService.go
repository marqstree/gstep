package ProcessService

import (
	"github.com/jinzhu/copier"
	"github.com/marqstree/gstep/dao/TemplateDao"
	"github.com/marqstree/gstep/enum/ProcessState"
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
	template := TemplateDao.GetLatestVersionByGroupId(dto.TemplateGroupId, tx)
	if nil == template {
		panic(ServerError.New("无效的模板"))
	}
	process.TemplateId = template.Id
	process.StartUserId = dto.StartUserId
	process.RootStep = template.RootStep
	process.State = ProcessState.STARTED.Code
	dao.SaveOrUpdate(&process, tx)

	TaskService.NewStartTask(&process, dto.StartUserId, tx)

	return process.Id
}
