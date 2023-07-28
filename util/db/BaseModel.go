package util_db

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
