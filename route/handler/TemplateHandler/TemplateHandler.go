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
	//已存在版本
	if template.Id > 0 {
		oldTemplate := dao.CheckById[entity.Template](template.Id, tx)
		oldTemplate.RootStep = template.RootStep
		oldTemplate.Title = template.Title
		oldTemplate.Version = template.Version
	} else if template.TemplateId > 0 { //已存在模板,新版本
		latestTemplate := TemplateDao.GetLatestVersionByTemplateId(template.TemplateId, tx)
		if nil == latestTemplate {
			AjaxJson.Fail("无效的流程模板id").Response(writer)
			return
		}
		newVersion := TemplateDao.NewVersion(template.TemplateId, tx)
		template.Version = newVersion
	} else { //新模板
		template.TemplateId = TemplateDao.NewTemplateId(tx)
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

func Info(writer http.ResponseWriter, request *http.Request) {
	dto := dto.TemplateQueryInfoDto{}
	RequestParsUtil.Body2dto(request, &dto)
	pDetail := TemplateService.QueryInfo(&dto, DbUtil.Db)
	if nil == pDetail {
		AjaxJson.Fail("查不到模板数据").Response(writer)
		return
	}
	AjaxJson.SuccessByData(*pDetail).Response(writer)
}

func SaveInfo(writer http.ResponseWriter, request *http.Request) {
	template := entity.Template{}
	RequestParsUtil.Body2dto(request, &template)

	tx := DbUtil.GetTx()
	oldTemplate := dao.CheckById[entity.Template](template.Id, tx)
	oldTemplate.Title = template.Title
	dao.SaveOrUpdate(&template, tx)

	tx.Commit()
	AjaxJson.SuccessByData(template.Id).Response(writer)
}
