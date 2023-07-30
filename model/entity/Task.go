package entity

import "github.com/marqstree/gstep/util/db/entity"

type Task struct {
	entity.BaseEntity
	ProcessId   int
	Form        map[string]any `gorm:"serializer:json"`
	AuditMethod string
	StepId      int
	Title       string
	Category    string
	State       string
}

func (e Task) TableName() string {
	return "task"
}

func (e Task) GetId() any {
	return e.Id
}
