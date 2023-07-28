package model_entity

import "github.com/marqstree/gstep/util/db"

type TemplateGroup struct {
	util_db.BaseModel
	Title string
}
