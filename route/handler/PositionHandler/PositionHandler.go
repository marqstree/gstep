package PositionHandler

import (
	"github.com/marqstree/gstep/dao/PositionDao"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"net/http"
)

func GetPositions(writer http.ResponseWriter, request *http.Request) {
	positions := PositionDao.GetPositions(DbUtil.Db)
	AjaxJson.SuccessByData(positions).Response(writer)
}
