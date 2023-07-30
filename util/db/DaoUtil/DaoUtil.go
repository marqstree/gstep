package DaoUtil

import (
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/DbUtil"
	"github.com/marqstree/gstep/util/db/entity"
	"gorm.io/gorm"
	"reflect"
)

func SaveOrUpdate(e any, tx *gorm.DB) {
	d := DbUtil.Db
	if nil != tx {
		d = tx
	}

	//反射获取id
	value := reflect.ValueOf(e)
	id := reflect.Indirect(value).FieldByName("Id").Int()

	if id < 1 {
		result := d.Create(e)
		if nil != result.Error {
			panic(result.Error)
		}
	} else {
		d.Save(e)
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
