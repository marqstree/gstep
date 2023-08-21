package TemplateDao

import (
	"fmt"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"gorm.io/gorm"
)

func GetLatestVersionByTemplateId(templateId int, tx *gorm.DB) *entity.Template {
	var entities []*entity.Template
	err := tx.Where("template_id=?", templateId).Order("version desc").Find(&entities).Error
	if nil != err {
		panic(err)
	}

	if len(entities) == 0 {
		return nil
	}

	return entities[0]
}

func NewTemplateId(tx *gorm.DB) int {
	maxTemplateId := 0
	err := tx.Raw("select ifnull(max(template_id),0) from template").Scan(&maxTemplateId).Error
	if nil != err {
		panic(ServerError.New(fmt.Sprintf("获取新templateId失败,%v", err)))
	}

	return maxTemplateId + 1
}

func NewVersion(templateId int, tx *gorm.DB) int {
	maxVersion := 0
	err := tx.Raw("select ifnull(max(version),0) from template where template_id=?", templateId).Scan(&maxVersion).Error
	if nil != err {
		panic(ServerError.New(fmt.Sprintf("查询最近版本号失败,%v", err)))
	}

	return maxVersion + 1
}
