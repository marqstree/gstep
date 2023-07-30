package entity

import (
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	Id        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
