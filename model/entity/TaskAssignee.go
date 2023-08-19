package entity

import "github.com/marqstree/gstep/util/db/entity"

type TaskAssignee struct {
	entity.BaseEntity
	TaskId      int             `json:"taskId"`
	UserId      string          `json:"userId"`
	State       string          `json:"state"`
	SubmitIndex int             `json:"submitIndex"`
	Form        *map[string]any `json:"form" gorm:"serializer:json"`
}

func (e TaskAssignee) TableName() string {
	return "task_assignee"
}

func (e TaskAssignee) GetId() any {
	return e.Id
}
