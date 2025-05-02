package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type League struct {
	BaseModel

	Name      string    `gorm:"not null"`
	LogoUrl   string    `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`

	UserId uuid.UUID `gorm:"not null"`
	User   User

	Teams  []Team            `gorm:"many2many:league_teams;"`
	Tracks []TrackmaniaTrack `gorm:"many2many:league_tracks;"`
}

func (s *League) BeforeDelete(tx *gorm.DB) error {
	results := []Result{}
	tx.Where("league_id = ?", s.Id).Find(&results)

	for _, result := range results {
		if err := tx.Delete(&result).Error; err != nil {
			return err
		}
	}

	teams := []Team{}
	tx.Where("league_id = ?", s.Id).Find(&teams)

	for _, team := range teams {
		if err := tx.Delete(&team).Error; err != nil {
			return err
		}
	}

	return nil
}
