package entity

import (
	"github.com/marqstree/gstep/util/db/entity"
)

type Template struct {
	entity.BaseEntity
	GroupId  int    `json:"groupId"`
	Title    string `json:"title"`
	Version  int    `json:"version"`
	RootStep Step   `json:"rootStep" gorm:"serializer:json"`
}

func (e Template) TableName() string {
	return "template"
}

func (e Template) GetId() any {
	return e.Id
}
