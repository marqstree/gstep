package TemplateService

import (
	"github.com/marqstree/gstep/dao/TemplateDao"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
	"strconv"
)

func Query(dto *dto.TemplateQueryDto, tx *gorm.DB) []entity.Template {
	list := []entity.Template{}
	sql := "select * from template " +
		" where 1=1 "
	if dto.TemplateId > 0 {
		sql = sql + " and id=" + strconv.Itoa(dto.TemplateId)
	}
	sql = sql + " limit " + strconv.Itoa(dto.Limit)
	sql = sql + " offset " + strconv.Itoa((dto.Page-1)*dto.Limit)

	err := tx.Raw(sql).Scan(&list).Error
	if nil != err {
		panic(err)
	}

	return list
}

func QueryDetail(dto *dto.TemplateQueryDetailDto, tx *gorm.DB) *entity.Template {
	template := entity.Template{}
	if dto.TemplateId > 0 {
		template = *dao.CheckById[entity.Template](dto.TemplateId, tx)
	}

	if dto.GroupId > 0 {
		template = *TemplateDao.GetLatestVersionByGroupId(dto.GroupId, tx)
	}

	return &template
}
