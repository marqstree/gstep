package TaskHandler

import (
	"github.com/marqstree/gstep/dao/TaskCandidateDao"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/TaskService"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/db/dao"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestParsUtil"
	"net/http"
)

func Pass(writer http.ResponseWriter, request *http.Request) {
	dto := dto.TaskPassDto{}
	RequestParsUtil.Body2dto(request, &dto)

	tx := DbUtil.GetTx()
	//校验taskid
	dao.CheckById[entity.Task](dto.TaskId, tx)
	//校验提交人在候选人列表中
	TaskCandidateDao.CheckCandidate(dto.UserId, dto.TaskId, tx)
	TaskService.Pass(&dto, tx)
	tx.Commit()

	AjaxJson.Success().Response(writer)
}

func Refuse(writer http.ResponseWriter, request *http.Request) {
	dto := dto.TaskRefuseDto{}
	RequestParsUtil.Body2dto(request, &dto)

	tx := DbUtil.GetTx()
	//校验taskid
	dao.CheckById[entity.Task](dto.TaskId, tx)
	//校验提交人在候选人列表中
	TaskCandidateDao.CheckCandidate(dto.UserId, dto.TaskId, tx)
	TaskService.Refuse(&dto, tx)
	tx.Commit()

	AjaxJson.Success().Response(writer)
}
