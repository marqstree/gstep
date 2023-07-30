package TemplateDao

import (
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/db/DbUtil"
)

func GetLatestVersionByGroupId(id int) *entity.Template {
	var entities []*entity.Template
	err := DbUtil.Db.Where("group_id=?", id).Order("version desc").Find(&entities).Error
	if nil != err {
		panic(err)
	}

	if len(entities) == 0 {
		return nil
	}

	return entities[0]
}
