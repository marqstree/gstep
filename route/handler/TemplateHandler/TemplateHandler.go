package TemplateHandler

import (
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/TemplateService"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestParsUtil"
	"net/http"
)

func Save(writer http.ResponseWriter, request *http.Request) {
	entity := entity.Template{}
	RequestParsUtil.Body2dto(request, &entity)

	tx := DbUtil.GetTx()
	id := TemplateService.SaveOrUpdate(&entity, tx)

	tx.Commit()
	AjaxJson.SuccessByData(id).Response(writer)
}
