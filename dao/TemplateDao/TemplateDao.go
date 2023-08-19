package TemplateDao

import (
	"fmt"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
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

func NewGroupId(tx *gorm.DB) int {
	maxGroudId := 0
	err := tx.Raw("select ifnull(max(group_id),0) from template").Scan(&maxGroudId).Error
	if nil != err {
		panic(ServerError.New(fmt.Sprintf("获取新groupId失败,%v", err)))
	}

	return maxGroudId + 1
}
