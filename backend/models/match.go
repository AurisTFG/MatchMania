package models

import "github.com/google/uuid"

type Match struct {
	BaseModel
	GameMode                string `gorm:"not null"`
	TrackmaniaCompetitionId int    `gorm:"not null"`

	LeagueId uuid.UUID `gorm:"not null"`

	League League
	Teams  []Team
}
