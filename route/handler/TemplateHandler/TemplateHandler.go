package TemplateHandler

import (
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/TemplateService"
	"github.com/marqstree/gstep/util/db/DaoUtil"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestParsUtil"
	"net/http"
)

func Save(writer http.ResponseWriter, request *http.Request) {
	dto := entity.Template{}
	RequestParsUtil.Body2dto(request, &dto)

	tx := DbUtil.GetTx()
	id := TemplateService.SaveOrUpdate(&dto, tx)

	tx.Commit()
	AjaxJson.SuccessByData(id).Response(writer)
}

func Start(writer http.ResponseWriter, request *http.Request) {
	dto := dto.ProcessStartDto{}
	RequestParsUtil.Body2dto(request, &dto)

	DaoUtil.CheckById(dto.StartUserId)

	tx := DbUtil.GetTx()
	id := TemplateService.Start(&dto, tx)
	tx.Commit()

	AjaxJson.SuccessByData(id).Response(writer)
}
