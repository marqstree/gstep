package CandidateCat

import "github.com/marqstree/gstep/util/enum"

type CandidateCat struct {
	enum.BaseEnum[string]
}

var USER = CandidateCat{}
var DEPARTMENT = CandidateCat{}
var POSITION = CandidateCat{}

// 需要手动处理的步骤类型列表
var Cats = [3]CandidateCat{}

func init() {
	USER.Code = "user"
	USER.Title = "用户"

	DEPARTMENT.Code = "department"
	DEPARTMENT.Title = "部门"

	POSITION.Code = "position"
	POSITION.Title = "职位"

	Cats = [3]CandidateCat{USER, DEPARTMENT, POSITION}
}

func IsContain(code string) bool {
	for _, v := range Cats {
		if v.Code == code {
			return true
		}
	}

	return false
}
