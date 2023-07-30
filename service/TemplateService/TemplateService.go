package TemplateService

import (
	"github.com/jinzhu/copier"
	"github.com/marqstree/gstep/dao/TemplateDao"
	"github.com/marqstree/gstep/enum"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/DaoUtil"
	"gorm.io/gorm"
)

func SaveOrUpdate(dto *entity.Template, tx *gorm.DB) int {
	entity := &entity.Template{}
	copier.Copy(entity, dto)

	old := TemplateDao.GetLatestVersionByGroupId(dto.GroupId)
	if nil != old {
		entity.Version = old.Version + 1
	} else {
		entity.Version = 1
	}
	DaoUtil.SaveOrUpdate(&entity.BaseEntity, tx)

	return entity.Id
}

func Start(dto *dto.ProcessStartDto, tx *gorm.DB) int {
	entity := &entity.Process{}
	copier.Copy(entity, dto)

	template := TemplateDao.GetLatestVersionByGroupId(dto.TemplateGroupId)
	if nil == template {
		panic(ServerError.New("无效的模板"))
	}
	entity.TemplateId = template.Id

	entity.State = enum.STARTED.Code

	DaoUtil.SaveOrUpdate(&entity.BaseEntity, tx)

	return entity.Id
}
