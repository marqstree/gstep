package enum

import "github.com/marqstree/gstep/util/enum"

type ProcessState struct {
	enum.BaseEnum[string]
}

var STARTED = ProcessState{}
var FINISHED = ProcessState{}

func init() {
	STARTED.Code = "started"
	STARTED.Title = "开始"

	FINISHED.Code = "finished"
	FINISHED.Title = "结束"
}
