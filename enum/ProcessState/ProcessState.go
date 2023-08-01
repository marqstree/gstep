package ProcessState

import "github.com/marqstree/gstep/util/enum"

type ProcessState struct {
	enum.BaseEnum[string]
}

var STARTED = ProcessState{}
var FINISH_PASS = ProcessState{}
var FINISH_REFUSE = ProcessState{}

func init() {
	STARTED.Code = "started"
	STARTED.Title = "开始"

	FINISH_PASS.Code = "finish_pass"
	FINISH_PASS.Title = "已通过"

	FINISH_REFUSE.Code = "finish_refuse"
	FINISH_REFUSE.Title = "已驳回"
}
