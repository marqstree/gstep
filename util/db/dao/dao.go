package dao

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
	id := reflect.Indirect(value).FieldByName("Id").Interface()

	switch id.(type) {
	case int:
		if id.(int) < 1 {
			result := d.Create(e)
			if nil != result.Error {
				panic(result.Error)
			}
		} else {
			d.Save(e)
		}
	case string:
		if len(id.(string)) < 1 {
			result := d.Create(e)
			if nil != result.Error {
				panic(result.Error)
			}
		} else {
			d.Save(e)
		}
	}

}

func CheckById[T entity.CommonEntity, I int | string](id I) T {
	var detail T

	err := DbUtil.Db.Table(detail.TableName()).Where("id=?", id).First(&detail).Error
	if nil != err {
		panic(err)
	}

	newId := detail.GetId()
	switch newId.(type) {
	case int:
		if newId == 0 {
			panic(ServerError.New("无效的表id"))
		}
	case string:
		if len(newId.(string)) == 0 {
			panic(ServerError.New("无效的表id"))
		}
	}

	return detail
}
