package entity

import (
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	Id        int            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
