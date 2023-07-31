package ProcessHandler

import (
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/ProcessService"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/db/dao"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestParsUtil"
	"net/http"
)

func Start(writer http.ResponseWriter, request *http.Request) {
	dto := dto.ProcessStartDto{}
	RequestParsUtil.Body2dto(request, &dto)

	tx := DbUtil.GetTx()
	dao.CheckById[entity.User](dto.StartUserId, tx)
	id := ProcessService.Start(&dto, tx)
	tx.Commit()

	AjaxJson.SuccessByData(id).Response(writer)
}
