package handler

import (
	model_dto "github.com/marqstree/gstep/model/dto"
	service_template "github.com/marqstree/gstep/service/template"
	util_db "github.com/marqstree/gstep/util/db"
	util_net_ajax "github.com/marqstree/gstep/util/net/ajax"
	util_net_parse "github.com/marqstree/gstep/util/net/parse"
	"net/http"
)

func SaveWorkflowTemplate(writer http.ResponseWriter, request *http.Request) {
	dto := model_dto.TemplateDto{}
	err := util_net_parse.Body2dto(request, &dto)
	if err != nil {
		util_net_ajax.ResponseAjaxJson(writer, util_net_ajax.FailByError(err))
		return
	}

	tx := util_db.GetTx()
	id, err := service_template.SaveOrUpdate(&dto, tx)
	if nil != err {
		tx.Rollback()
		util_net_ajax.ResponseAjaxJson(writer, util_net_ajax.FailByError(err))
		return
	}

	tx.Commit()
	util_net_ajax.ResponseAjaxJson(writer, util_net_ajax.SuccessByData(id))
}
