package models

import (
	"time"

	"github.com/google/uuid"
)

type Result struct {
	BaseModel

	StartDate     time.Time `gorm:"not null"`
	EndDate       time.Time `gorm:"not null"`
	Score         string    `gorm:"not null"`
	OpponentScore string    `gorm:"not null"`

	TeamId         uuid.UUID `gorm:"not null"`
	Team           Team
	OpponentTeamId uuid.UUID `gorm:"not null"`
	OpponentTeam   Team
	LeagueId       uuid.UUID `gorm:"not null"`
	League         League
	UserId         uuid.UUID `gorm:"not null"`
	User           User
}
