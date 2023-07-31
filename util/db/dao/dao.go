package dao

import (
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/db/entity"
	"gorm.io/gorm"
	"reflect"
)

func SaveOrUpdate(pEntity any, tx *gorm.DB) {
	//反射获取id
	value := reflect.ValueOf(pEntity)
	id := reflect.Indirect(value).FieldByName("Id").Interface()

	switch id.(type) {
	case int:
		if id.(int) < 1 {
			result := tx.Create(pEntity)
			if nil != result.Error {
				panic(result.Error)
			}
		} else {
			tx.Save(pEntity)
		}
	case string:
		if len(id.(string)) < 1 {
			result := tx.Create(pEntity)
			if nil != result.Error {
				panic(result.Error)
			}
		} else {
			tx.Save(pEntity)
		}
	}
}

func CheckById[T entity.CommonEntity, I int | string](id I, tx *gorm.DB) *T {
	var detail T

	err := tx.Table(detail.TableName()).Where("id=?", id).First(&detail).Error
	if nil != err {
		panic(err)
	}

	newId := detail.GetId()
	err = CheckId(newId)
	if nil != err {
		panic(err)
	}

	return &detail
}

func CheckId(id any) error {
	switch id.(type) {
	case int:
		if id == 0 {
			return ServerError.New("无效的表id")
		}
	case string:
		if len(id.(string)) == 0 {
			return ServerError.New("无效的表id")
		}
	}

	return nil
}
