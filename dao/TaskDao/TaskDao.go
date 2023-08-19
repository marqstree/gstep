package TaskDao

import (
	"fmt"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"gorm.io/gorm"
)

func QueryTaskByStepId(stepId int, processId int, tx *gorm.DB) *entity.Task {
	var detail entity.Task
	tx.Table(detail.TableName()).Where("step_id=? and process_id=?", stepId, processId).First(&detail)
	if 0 == detail.Id {
		return nil
	} else {
		return &detail
	}
}

func QueryMyPendingTasks(userId string, tx *gorm.DB) (*[]entity.Task, int) {
	total := 0
	var details []entity.Task

	err := tx.Raw("select count(1) from task "+
		" where state='started' "+
		" and exists(select 1 from task_assignee ta"+
		" where ta.task_id=task.id "+
		" and ta.user_id=?)", userId).Scan(&total).Error
	if nil != err {
		msg := fmt.Sprintf("找不到待处理任务: %s", err)
		panic(ServerError.New(msg))
	}
	err = tx.Raw("select * from task "+
		" where state='started' "+
		" and exists(select 1 from task_assignee ta"+
		" where ta.task_id=task.id "+
		" and ta.user_id=?)", userId).Scan(&details).Error
	if nil != err {
		msg := fmt.Sprintf("找不到待处理任务: %s", err)
		panic(ServerError.New(msg))
	}
	return &details, total
}
