package dao_template

import (
	"github.com/marqstree/gstep/model/entity"
	util_db "github.com/marqstree/gstep/util/db"
	"gorm.io/gorm"
)

func SaveOrUpdate(entity *model_entity.Template, tx *gorm.DB) error {
	d := util_db.Db
	if nil != tx {
		d = tx
	}
	if entity.Id < 1 {
		result := d.Create(entity)
		if nil != result.Error {
			return result.Error
		}
	} else {
		d.Save(entity)
	}

	return nil
}

func GetById(id int) (user *model_entity.Template, err error) {
	var entity model_entity.Template
	util_db.Db.First(&entity, id)

	return &entity, nil
}

func GetLatestVersionByGroupId(id int) (*model_entity.Template, error) {
	var entities []*model_entity.Template
	err := util_db.Db.Where("group_id=?", id).Order("version desc").Find(&entities).Error
	if nil != err {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, nil
	}

	return entities[0], nil
}
