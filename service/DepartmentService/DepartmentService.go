package DepartmentService

import (
	"github.com/marqstree/gstep/dao/DepartmentDao"
	"github.com/marqstree/gstep/dao/UserDao"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/model/vo"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"gorm.io/gorm"
)

func ToVo(bean entity.Department, tx *gorm.DB) vo.DepartmentVo {
	aVo := vo.DepartmentVo{}
	aVo.Department = bean
	userCnt := UserDao.GetDepartmentUserCount(bean.Id, tx)
	aVo.UserCount = userCnt

	cnt := DepartmentDao.GetChildDepartmentCount(bean.Id, tx)
	aVo.HasSubDepartments = cnt > 0

	return aVo
}

func GetChildDepartments(dto dto.DepartmentQueryChildDto, tx *gorm.DB) []vo.DepartmentVo {
	childDepartments := DepartmentDao.GetChildDepartments(dto.ParentId, DbUtil.Db)
	vos := []vo.DepartmentVo{}
	for _, v := range childDepartments {
		aVo := ToVo(v, tx)
		vos = append(vos, aVo)
	}
	return vos
}

func GetDepartmentUsers(dto dto.DepartmentQueryUsersDto, tx *gorm.DB) []entity.User {
	users := UserDao.GetDepartmentUsers(dto.DepartmentId, DbUtil.Db)
	return users
}
