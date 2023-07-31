package TaskAssigneeDao

import (
	"github.com/marqstree/gstep/enum/TaskState"
	"gorm.io/gorm"
)

func PassCount(taskId int, tx *gorm.DB) int64 {
	var count int64
	tx.Table("task_assignee").Where("task_id=? and state=?", taskId, TaskState.PASS.Code).Count(&count)
	return count
}
