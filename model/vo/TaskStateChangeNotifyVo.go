package vo

import "github.com/marqstree/gstep/model/entity"

type TaskStateChangeNotifyVo struct {
	Process entity.Process                `json:"process"`
	Tasks   []TaskStateChangeNotifyTaskVo `json:"tasks"`
}
