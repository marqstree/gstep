package TemplateHandler

import (
	"github.com/marqstree/gstep/dao/TemplateDao"
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
	if 0 == template.GroupId {
		template.GroupId = TemplateDao.NewGroupId(tx)
		template.Version = 1
	}
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
	dto := dto.TemplateQueryDetailDto{}
	RequestParsUtil.Body2dto(request, &dto)
	pDetail := TemplateService.QueryDetail(&dto, DbUtil.Db)
	if nil == pDetail {
		AjaxJson.Fail("查不到模板数据").Response(writer)
		return
	}
	AjaxJson.SuccessByData(*pDetail).Response(writer)
}
