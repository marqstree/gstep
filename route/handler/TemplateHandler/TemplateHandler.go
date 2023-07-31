package TemplateHandler

import (
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/db/dao"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestParsUtil"
	"net/http"
)

func Save(writer http.ResponseWriter, request *http.Request) {
	template := entity.Template{}
	RequestParsUtil.Body2dto(request, &template)

	tx := DbUtil.GetTx()
	dao.SaveOrUpdate(&template, tx)

	tx.Commit()
	AjaxJson.SuccessByData(template.Id).Response(writer)
}
