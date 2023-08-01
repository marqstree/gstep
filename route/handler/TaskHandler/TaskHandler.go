package TaskHandler

import (
	"github.com/marqstree/gstep/dao/TaskCandidateDao"
	"github.com/marqstree/gstep/dao/TaskDao"
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
	pTask := dao.CheckById[entity.Task](dto.TaskId, tx)
	TaskService.CheckTaskCanChange(pTask)
	//校验提交人在候选人列表中
	TaskCandidateDao.CheckCandidate(dto.UserId, dto.TaskId, tx)
	TaskService.Pass(&dto, tx)
	tx.Commit()

	//发送任务状态变更通知
	pTask = dao.CheckById[entity.Task](dto.TaskId, DbUtil.Db)
	TaskService.NotifyTaskStateChange(pTask)

	AjaxJson.Success().Response(writer)
}

func Refuse(writer http.ResponseWriter, request *http.Request) {
	dto := dto.TaskRefuseDto{}
	RequestParsUtil.Body2dto(request, &dto)

	tx := DbUtil.GetTx()
	//校验taskid
	pTask := dao.CheckById[entity.Task](dto.TaskId, tx)
	TaskService.CheckTaskCanChange(pTask)
	//校验提交人在候选人列表中
	TaskCandidateDao.CheckCandidate(dto.UserId, dto.TaskId, tx)
	TaskService.Refuse(&dto, tx)
	tx.Commit()

	AjaxJson.Success().Response(writer)
}

func Cease(writer http.ResponseWriter, request *http.Request) {
	dto := dto.TaskCeaseDto{}
	RequestParsUtil.Body2dto(request, &dto)

	tx := DbUtil.GetTx()
	//校验taskid
	pTask := dao.CheckById[entity.Task](dto.TaskId, tx)
	TaskService.CheckTaskCanChange(pTask)

	TaskService.Cease(&dto, tx)
	tx.Commit()

	AjaxJson.Success().Response(writer)
}

func Pending(writer http.ResponseWriter, request *http.Request) {
	dto := dto.TaskPendingDto{}
	RequestParsUtil.Body2dto(request, &dto)

	tasks, total := TaskDao.QueryMyPendingTasks(dto.UserId, DbUtil.Db)
	AjaxJson.SuccessByPagination(*tasks, total).Response(writer)
}
