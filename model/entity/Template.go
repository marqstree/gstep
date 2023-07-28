package model_entity

import "github.com/marqstree/gstep/util/db"

type Template struct {
	util_db.BaseModel
	GroupId int64
	Title   string
	Version int64
	Content string
}
