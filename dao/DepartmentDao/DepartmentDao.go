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
		pId = config.Config.Department.RootParentDepartmentId
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

func GetGrandsonDepartments(parentId string, tx *gorm.DB) []entity.Department {
	var departments []entity.Department

	err := tx.Raw("WITH recursive temp AS ( "+
		" SELECT a.id, a.name, a.parent_id FROM department a "+
		" WHERE a.id = ? "+
		" UNION ALL "+
		" SELECT b.id,b.name,b.parent_id "+
		" FROM department b JOIN temp ON (b.parent_id = temp.id) ) "+
		" SELECT temp.id, temp.name, temp.parent_id FROM temp", parentId).Scan(&departments).Error
	if nil != err {
		msg := fmt.Sprintf("找不到部门及其子部门: %s", err)
		panic(ServerError.New(msg))
	}
	return departments
}

func GetGrandsonDepartmentIds(parentId string, tx *gorm.DB) []string {
	departments := GetGrandsonDepartments(parentId, tx)
	ids := []string{}
	for _, v := range departments {
		ids = append(ids, v.Id)
	}
	return ids
}

func GetChildDepartmentCount(parentId string, tx *gorm.DB) int {
	cnt := 0
	pId := parentId
	if len(parentId) == 0 {
		pId = config.Config.Department.RootParentDepartmentId
	}
	err := tx.Raw("select count(1) from department a "+
		" where a.parent_id=?", pId).Scan(&cnt).Error
	if nil != err {
		msg := fmt.Sprintf("找不到子部门数量失败: %s", err)
		panic(ServerError.New(msg))
	}
	return cnt
}
