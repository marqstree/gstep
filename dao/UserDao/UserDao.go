package UserDao

import (
	"fmt"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"gorm.io/gorm"
)

func GetDepartmentUsers(departmentId string, tx *gorm.DB) []entity.User {
	users := []entity.User{}

	err := tx.Raw("select * from user a "+
		" where a.department_id=?"+
		" order by a.id asc ", departmentId).Scan(&users).Error
	if nil != err {
		msg := fmt.Sprintf("找不到部门员工: %s", err)
		panic(ServerError.New(msg))
	}
	return users
}

func GetDepartmentUserCount(departmentId string, tx *gorm.DB) int {
	cnt := 0

	err := tx.Raw("select count(1) from user a "+
		" where a.department_id=?"+
		" order by a.id asc ", departmentId).Scan(&cnt).Error
	if nil != err {
		msg := fmt.Sprintf("找不到部门员工数量: %s", err)
		panic(ServerError.New(msg))
	}
	return cnt
}
