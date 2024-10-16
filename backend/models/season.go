package models

import (
	"time"
)

type Season struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	StartDate time.Time
	EndDate   time.Time

	Teams []Team `gorm:"foreignKey:SeasonID"`
}
