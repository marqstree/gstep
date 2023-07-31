package entity

import "github.com/marqstree/gstep/util/db/entity"

type TaskAssignee struct {
	entity.BaseEntity
	TaskId int
	UserId string
	State  string
}

func (e TaskAssignee) TableName() string {
	return "task_assignee"
}

func (e TaskAssignee) GetId() any {
	return e.Id
}
