package DepartmentHandler

import (
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/service/DepartmentService"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestParsUtil"
	"net/http"
)

func GetChildDepartments(writer http.ResponseWriter, request *http.Request) {
	dto := dto.DepartmentQueryChildDto{}
	RequestParsUtil.Body2dto(request, &dto)

	childDepartments := DepartmentService.GetChildDepartments(dto, DbUtil.Db)

	AjaxJson.SuccessByData(childDepartments).Response(writer)
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	dto := dto.DepartmentQueryUsersDto{}
	RequestParsUtil.Body2dto(request, &dto)

	users := DepartmentService.GetDepartmentUsers(dto, DbUtil.Db)

	AjaxJson.SuccessByData(users).Response(writer)
}
