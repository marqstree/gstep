package entity

import (
	"github.com/marqstree/gstep/util/db/entity"
)

type Process struct {
	entity.BaseEntity
	TemplateId  int
	StartUserId string
	RootStep    Step `gorm:"serializer:json"`
	State       string
}

func (e Process) TableName() string {
	return "process"
}

func (e Process) GetId() any {
	return e.Id
}
