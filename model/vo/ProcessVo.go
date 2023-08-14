package vo

import "github.com/marqstree/gstep/model/entity"

type ProcessVo struct {
	entity.Process
	Template entity.Template `json:"template"`
}
