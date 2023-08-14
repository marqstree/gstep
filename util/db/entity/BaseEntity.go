package entity

import (
	"github.com/marqstree/gstep/util/LocalTime"
)

type BaseEntity struct {
	Id        int                  `json:"id" gorm:"primarykey"`
	CreatedAt *LocalTime.LocalTime `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt *LocalTime.LocalTime `json:"updatedAt" gorm:"autoUpdateTime:false"`
	DeletedAt *LocalTime.LocalTime `json:"deletedAt" gorm:"softDelete"`
}
