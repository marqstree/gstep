package TaskAssigneeDao

import (
	"github.com/marqstree/gstep/enum/TaskState"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
)

func PassCount(taskId int, tx *gorm.DB) int {
	var count int64
	tx.Table("task_assignee").Where("task_id=? and state=?", taskId, TaskState.PASS.Code).Count(&count)
	return int(count)
}

func CheckAssigneeCanSubmit(taskId int, userId string, state string, tx *gorm.DB) {
	task := dao.CheckById[entity.Task](taskId, tx)
	if task.State == TaskState.PASS.Code {
		panic(ServerError.New("任务已完成,不可重复提交"))
	}

	record := entity.TaskAssignee{}
	tx.Table("task_assignee").Where("task_id=? and user_id=?", taskId, userId).Order("create_at desc").First(&record)
	if record.Id > 0 && record.State == state {
		panic(ServerError.New("重复提交任务"))
	}
}

func GetMaxSubmitIndex(processId int, tx *gorm.DB) int {
	var submitIndex int64
	tx.Raw("select ifnull(max(submit_index),0) from task_assignee a "+
		" where exists(select 1 from task b where b.id=a.task_id and b.process_id=?)", processId).Scan(&submitIndex)
	return int(submitIndex)
}

func GetLastSubmitAssigneesByTask(processId int, taskId int, tx *gorm.DB) []entity.TaskAssignee {
	maxSubmitIndex := GetMaxSubmitIndex(processId, tx)

	assignees := []entity.TaskAssignee{}
	tx.Raw("select * from task_assignee a "+
		" where exists(select 1 from task b where b.id=a.task_id and b.process_id=?) "+
		" and a.submit_index=?"+
		" and a.task_id=? "+
		" order by a.id desc ", processId, maxSubmitIndex, taskId).Scan(&assignees)

	return assignees
}

func GetTasksByLastSubmitIndex(processId int, tx *gorm.DB) []entity.Task {
	maxSubmitIndex := GetMaxSubmitIndex(processId, tx)

	taskIds := []int{}
	tx.Raw("select distinct task_id from task_assignee a "+
		" where exists(select 1 from task b where b.id=a.task_id and b.process_id=?) "+
		" and a.submit_index=?"+
		" order by a.task_id desc ", processId, maxSubmitIndex).Scan(&taskIds)

	tasks := []entity.Task{}
	for _, id := range taskIds {
		pTask := dao.CheckById[entity.Task](id, tx)
		tasks = append(tasks, *pTask)
	}

	return tasks
}
