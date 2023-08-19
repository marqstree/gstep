package TaskState

import "github.com/marqstree/gstep/util/enum"

type TaskState struct {
	enum.BaseEnum[string]
}

var STARTED = TaskState{}
var PASS = TaskState{}
var REFUSE = TaskState{}
var RETREAT = TaskState{}

func init() {
	STARTED.Code = "started"
	STARTED.Title = "开始"

	PASS.Code = "pass"
	PASS.Title = "同意"

	REFUSE.Code = "refuse"
	REFUSE.Title = "拒绝"

	RETREAT.Code = "retreat"
	RETREAT.Title = "退回"
}
