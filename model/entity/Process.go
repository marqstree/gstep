package entity

import (
	"github.com/marqstree/gstep/util/db/entity"
)

type Process struct {
	entity.BaseEntity
	TemplateId  int
	StartUserId string
	State       string
}
