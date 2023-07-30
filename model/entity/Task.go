package entity

import "github.com/marqstree/gstep/util/db/entity"

type Task struct {
	entity.BaseEntity
	ProcessId int
	Form      *map[string]any `json:"form,omitempty"`
	StepId    int
	State     string
}
