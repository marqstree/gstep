package entity

import "github.com/marqstree/gstep/util/db/entity"

type TaskCandidate struct {
	entity.BaseEntity
	TaskId int
	UserId string
}

func (e TaskCandidate) TableName() string {
	return "task_candidate"
}

func (e TaskCandidate) GetId() any {
	return e.Id
}
