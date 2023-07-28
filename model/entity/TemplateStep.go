package model_entity

import "github.com/marqstree/gstep/util/db"

type TemplateStep struct {
	util_db.BaseModel
	Title       string
	NextStepIds []string
	PrevStepId  []string
	Category    string
	AuditMethod string
	Expression  string
}
