package StepCat

import "github.com/marqstree/gstep/util/enum"

type StepCat struct {
	enum.BaseEnum[string]
}

var AUDIT = StepCat{}
var CONDITION = StepCat{}
var NOTIFY = StepCat{}
var START = StepCat{}

// 需要手动处理的步骤类型列表
var AuditStepCats = [2]StepCat{}
var StepCats = [4]StepCat{}

func init() {
	AUDIT.Code = "audit"
	AUDIT.Title = "审核"

	CONDITION.Code = "condition"
	CONDITION.Title = "条件"

	NOTIFY.Code = "notify"
	NOTIFY.Title = "抄送"

	START.Code = "start"
	START.Title = "开始"

	StepCats = [4]StepCat{START, NOTIFY, CONDITION, AUDIT}
	AuditStepCats = [2]StepCat{AUDIT, START}
}

func IsContain(code string) bool {
	for _, v := range StepCats {
		if v.Code == code {
			return true
		}
	}

	return false
}

func IsContainAudit(code string) bool {
	for _, v := range AuditStepCats {
		if v.Code == code {
			return true
		}
	}

	return false
}
