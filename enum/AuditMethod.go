package enum

import (
	"github.com/marqstree/gstep/util/enum"
)

type AuditMethod struct {
	enum.BaseEnum[string]
}

var AND = AuditMethod{}
var OR = AuditMethod{}

func init() {
	AND.Code = "and"
	AND.Title = "会签"

	OR.Code = "or"
	OR.Title = "或签"
}
