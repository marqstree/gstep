package TaskCandidateDao

import (
	"fmt"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/service/StepService"
	"github.com/marqstree/gstep/util/CONSTANT"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
)

func CheckCandidate(userId string, taskId int, tx *gorm.DB) {
	var detail entity.TaskCandidate
	err := tx.Table(detail.TableName()).Where("value=? and task_id=?", userId, taskId).First(&detail).Error
	if nil != err {
		panic(ServerError.New("无效的候选人"))
	}

	err = dao.CheckId(detail.Id)
	if nil != err {
		msg := fmt.Sprintf("用户(%s)无权限", userId)
		panic(ServerError.NewByCode(CONSTANT.NEED_LOGIN_CODE, msg))
	}
}

// 候选人条数
func CandidateCount(taskId int, tx *gorm.DB) int {
	pTask := dao.CheckById[entity.Task](taskId, tx)
	pProcess := dao.CheckById[entity.Process](pTask.ProcessId, tx)
	pTemplate := dao.CheckById[entity.Template](pProcess.TemplateId, tx)
	pStep := StepService.FindStep(&pTemplate.RootStep, pTask.StepId)
	return len(pStep.Candidates)
}
