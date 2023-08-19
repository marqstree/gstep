package vo

import "github.com/marqstree/gstep/model/entity"

type TaskStateChangeNotifyTaskVo struct {
	Task      entity.Task           `json:"task"`
	Assignees []entity.TaskAssignee `json:"assignees"`
}
