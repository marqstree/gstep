package TaskHandler

import (
	"github.com/marqstree/gstep/dao/TaskAssigneeDao"
	"github.com/marqstree/gstep/dao/TaskCandidateDao"
	"github.com/marqstree/gstep/dao/TaskDao"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/StepService"
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

	//检查任务重复审核
	TaskService.CheckTaskCanChange(pTask)

	//检查流程提交人是否是候选人
	pProcess := dao.CheckById[entity.Process](pTask.ProcessId, tx)
	pTemplate := dao.CheckById[entity.Template](pProcess.TemplateId, tx)
	StepService.CheckCandidate(dto.UserId, &pTemplate.RootStep, pTask.StepId, tx)

	//检查提交人重复提交
	TaskAssigneeDao.CheckAssigneeCanPass(pTask.Id, dto.UserId, tx)

	//审核通过
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
