package DaoUtil

import (
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/db/entity"
	"gorm.io/gorm"
)

func SaveOrUpdate(entity *entity.BaseEntity, tx *gorm.DB) {
	d := DbUtil.Db
	if nil != tx {
		d = tx
	}
	if entity.Id < 1 {
		result := d.Create(entity)
		if nil != result.Error {
			panic(result.Error)
		}
	} else {
		d.Save(entity)
	}
}

func CheckById[T entity.BaseEntity](id any) *T {
	var entity *T
	DbUtil.Db.First(entity, id)

	if nil == entity {
		panic(ServerError.New("无效的表id"))
	}

	return entity
}
