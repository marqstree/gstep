package PositionDao

import (
	"fmt"
	"github.com/marqstree/gstep/model/entity"
	"github.com/marqstree/gstep/util/ServerError"
	"gorm.io/gorm"
)

func GetPositions(tx *gorm.DB) []entity.Position {
	positions := []entity.Position{}

	err := tx.Raw(" SELECT * from position a" +
		" order by a.title").Scan(&positions).Error
	if nil != err {
		msg := fmt.Sprintf("找不到职位: %s", err)
		panic(ServerError.New(msg))
	}
	return positions
}
