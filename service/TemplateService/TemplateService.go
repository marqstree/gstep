package TemplateService

import (
	"github.com/jinzhu/copier"
	"github.com/marqstree/gstep/dao/TemplateDao"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/db/dao"
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
	dao.SaveOrUpdate(entity, tx)

	return entity.Id
}
