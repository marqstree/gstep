package StepCat

import "github.com/marqstree/gstep/util/enum"

type StepCat struct {
	enum.BaseEnum[string]
}

var AUDIT = StepCat{}
var CONDITION = StepCat{}
var NOTIFY = StepCat{}
var START = StepCat{}
var END = StepCat{}

func init() {
	AUDIT.Code = "audit"
	AUDIT.Title = "审核"

	CONDITION.Code = "condition"
	CONDITION.Title = "条件"

	NOTIFY.Code = "notify"
	NOTIFY.Title = "抄送"

	START.Code = "start"
	START.Title = "开始"

	END.Code = "end"
	END.Title = "结束"
}
