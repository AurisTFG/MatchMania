package models

import (
	"time"

	"github.com/google/uuid"
)

type League struct {
	BaseModel
	Name      string `gorm:"not null"`
	LogoUrl   string
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`

	UserId uuid.UUID `gorm:"not null"`

	User    User
	Queue   *Queue
	Results []Result
	Matches []Match           `gorm:"foreignKey:LeagueId"`
	Teams   []Team            `gorm:"many2many:league_teams;"`
	Tracks  []TrackmaniaTrack `gorm:"many2many:league_tracks;"`
}
