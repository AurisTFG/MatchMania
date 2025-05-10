package models

import (
	"time"

	"gorm.io/gorm"
)

type AppSetting struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Key       string         `gorm:"primaryKey;size:100"`
	Value     string         `gorm:"type:text"`
}
