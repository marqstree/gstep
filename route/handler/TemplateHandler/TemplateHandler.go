package TemplateHandler

import (
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/TemplateService"
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

func Query(writer http.ResponseWriter, request *http.Request) {
	dto := dto.TemplateQueryDto{}
	RequestParsUtil.Body2dto(request, &dto)
	list := TemplateService.Query(&dto, DbUtil.Db)
	AjaxJson.SuccessByData(list).Response(writer)
}

func Detail(writer http.ResponseWriter, request *http.Request) {
	//dto := dto.TemplateQueryDetailDto{}
	//RequestParsUtil.Body2dto(request, &dto)
	//detail := TemplateService.QueryDetail(&dto, DbUtil.Db)
	//AjaxJson.SuccessByData(*detail).Response(writer)
	AjaxJson.Success().Response(writer)
}
