package TemplateDao

import (
	"github.com/marqstree/gstep/model/entity"
	"gorm.io/gorm"
)

func GetLatestVersionByGroupId(id int, tx *gorm.DB) *entity.Template {
	var entities []*entity.Template
	err := tx.Where("group_id=?", id).Order("version desc").Find(&entities).Error
	if nil != err {
		panic(err)
	}

	if len(entities) == 0 {
		return nil
	}

	return entities[0]
}
