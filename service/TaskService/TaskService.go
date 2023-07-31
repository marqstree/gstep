package TaskService

import (
	"github.com/marqstree/gstep/dao/TaskAssigneeDao"
	"github.com/marqstree/gstep/dao/TaskCandidateDao"
	"github.com/marqstree/gstep/enum/AuditMethod"
	"github.com/marqstree/gstep/enum/TaskState"
	"github.com/marqstree/gstep/model/dto"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
)

func Pass(pDto *dto.TaskPassDto, tx *gorm.DB) {
	pTask := dao.CheckById[entity.Task](pDto.TaskId, tx)

	assignee := entity.TaskAssignee{}
	assignee.TaskId = pTask.Id
	assignee.UserId = pDto.UserId
	assignee.State = TaskState.PASS.Code
	dao.SaveOrUpdate(&assignee, tx)

	pTask.Form = pDto.Form
	if isCanPass(pTask, tx) {
		pTask.State = TaskState.PASS.Code
	}
	dao.SaveOrUpdate(*pTask, tx)
}

func isCanPass(pTask *entity.Task, tx *gorm.DB) bool {
	passCount := TaskAssigneeDao.PassCount(pTask.Id, tx)

	if pTask.AuditMethod == AuditMethod.OR.Code {
		return passCount > 0
	} else {
		candidateCount := TaskCandidateDao.CandidateCount(pTask.Id, tx)
		return passCount >= candidateCount
	}
}
