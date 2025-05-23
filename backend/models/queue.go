package models

import "github.com/google/uuid"

type Queue struct {
	BaseModel
	GameMode string `gorm:"not null"`

	LeagueId uuid.UUID `gorm:"not null"`

	League League
	Teams  []Team
}
