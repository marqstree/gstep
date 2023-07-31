package TaskCandidateDao

import (
	"fmt"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/CONSTANT"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/dao"
	"gorm.io/gorm"
)

func CheckCandidate(userId string, taskId int, tx *gorm.DB) {
	var detail entity.TaskCandidate
	err := tx.Table(detail.TableName()).Where("user_id=? and task_id=?", userId, taskId).First(&detail).Error
	if nil != err {
		panic(ServerError.New("无效的候选人"))
	}

	err = dao.CheckId(detail.Id)
	if nil != err {
		msg := fmt.Sprintf("用户(%s)无权限", userId)
		panic(ServerError.NewByCode(CONSTANT.NEED_LOGIN_CODE, msg))
	}
}

func CandidateCount(taskId int, tx *gorm.DB) int64 {
	var count int64
	err := tx.Table("task_candiddate").Where("task_id=?", taskId).Count(&count)
	if nil != err {
		panic(err)
	}
	return count
}
