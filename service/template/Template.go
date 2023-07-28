package service_template

import (
	"github.com/jinzhu/copier"
	dao_template "github.com/marqstree/gstep/dao/template"
	model_dto "github.com/marqstree/gstep/model/dto"
	model_entity "github.com/marqstree/gstep/model/entity"
	util_json "github.com/marqstree/gstep/util/json"
	"gorm.io/gorm"
)

func SaveOrUpdate(dto *model_dto.TemplateDto, tx *gorm.DB) (int, error) {
	entity := &model_entity.Template{}
	copier.Copy(entity, dto)

	jsonStr, err := util_json.Obj2json(dto)
	if nil != err {
		return 0, err
	}
	entity.Content = jsonStr

	old, err := dao_template.GetLatestVersionByGroupId(dto.GroupId)
	if nil != err {
		return 0, err
	}

	if nil != old {
		entity.Version = old.Version + 1
	} else {
		entity.Version = 1
	}
	err = dao_template.SaveOrUpdate(entity, tx)
	if nil != err {
		return 0, err
	}

	return entity.Id, nil
}
