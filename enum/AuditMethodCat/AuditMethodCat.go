package AuditMethodCat

import (
	"github.com/marqstree/gstep/util/enum"
)

type AuditMethodCat struct {
	enum.BaseEnum[string]
}

var AND = AuditMethodCat{}
var OR = AuditMethodCat{}

func init() {
	AND.Code = "and"
	AND.Title = "会签"

	OR.Code = "or"
	OR.Title = "或签"
}
