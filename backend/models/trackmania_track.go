package models

import (
	"time"

	"github.com/google/uuid"
)

type TrackmaniaTrack struct {
	Id           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt    time.Time
	Uid          string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Author       string
	ThumbnailUrl string `gorm:"not null"`

	UserId uuid.UUID `gorm:"not null"`

	User    User
	Leagues []League `gorm:"many2many:league_tracks;"`
}
