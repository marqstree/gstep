package TaskDao

import (
	"fmt"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"gorm.io/gorm"
)

func QueryTaskByStepId(stepId int, processId int, tx *gorm.DB) *entity.Task {
	var detail entity.Task
	err := tx.Table(detail.TableName()).Where("step_id=? and process_id=?", stepId, processId).First(&detail).Error
	if nil != err {
		msg := fmt.Sprintf("找不到流程步骤对应的任务: %s", err)
		panic(ServerError.New(msg))
	}
	return &detail
}
