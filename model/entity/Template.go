package entity

import (
	"github.com/marqstree/gstep/util/db/entity"
)

type Template struct {
	entity.BaseEntity
	TemplateId int    `json:"templateId"`
	Title      string `json:"title"`
	Version    int    `json:"version"`
	RootStep   Step   `json:"rootStep" gorm:"serializer:json"`
}

func (e Template) TableName() string {
	return "template"
}

func (e Template) GetId() any {
	return e.Id
}
