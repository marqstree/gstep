package TaskAssigneeDao

import (
	"github.com/marqstree/gstep/enum/TaskState"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"gorm.io/gorm"
)

func PassCount(taskId int, tx *gorm.DB) int {
	var count int64
	tx.Table("task_assignee").Where("task_id=? and state=?", taskId, TaskState.PASS.Code).Count(&count)
	return int(count)
}

func CheckAssigneeCanPass(taskId int, userId string, tx *gorm.DB) {
	record := entity.TaskAssignee{}
	tx.Table("task_assignee").Where("task_id=? and user_id=?", taskId, userId).Order("create_at desc").First(&record)
	if record.Id > 0 && record.State == TaskState.PASS.Code {
		panic(ServerError.New("重复审核通过任务"))
	}
}
