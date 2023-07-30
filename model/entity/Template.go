package entity

import (
	"github.com/marqstree/gstep/util/db/entity"
)

type Template struct {
	entity.BaseEntity
	GroupId  int
	Title    string
	Version  int
	RootStep Step `gorm:"serializer:json"`
}

func (e Template) TableName() string {
	return "template"
}

func (e Template) GetId() any {
	return e.Id
}
