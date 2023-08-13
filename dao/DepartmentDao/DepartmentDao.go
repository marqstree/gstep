package DepartmentDao

import (
	"fmt"
	"github.com/marqstree/gstep/config"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"gorm.io/gorm"
)

func GetChildDepartments(parentId string, tx *gorm.DB) []entity.Department {
	var departments []entity.Department

	pId := parentId
	if len(parentId) == 0 {
		pId = config.Config.Department.RootDepartmentId
	}
	err := tx.Raw("select * from department a "+
		" where a.parent_id=?"+
		" order by a.id asc ", pId).Scan(&departments).Error
	if nil != err {
		msg := fmt.Sprintf("找不到子部门: %s", err)
		panic(ServerError.New(msg))
	}
	return departments
}

func GetChildDepartmentCount(parentId string, tx *gorm.DB) int {
	cnt := 0
	pId := parentId
	if len(parentId) == 0 {
		pId = config.Config.Department.RootDepartmentId
	}
	err := tx.Raw("select count(1) from department a "+
		" where a.parent_id=?", pId).Scan(&cnt).Error
	if nil != err {
		msg := fmt.Sprintf("找不到子部门数量失败: %s", err)
		panic(ServerError.New(msg))
	}
	return cnt
}
