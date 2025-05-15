package models

import (
	"time"

	"github.com/google/uuid"
)

type Result struct {
	BaseModel
	StartDate       time.Time `gorm:"not null"`
	EndDate         time.Time `gorm:"not null"`
	Score           uint      `gorm:"not null"`
	OpponentScore   uint      `gorm:"not null"`
	EloDiff         int       `gorm:"not null"`
	OpponentEloDiff int       `gorm:"not null"`
	NewElo          uint      `gorm:"not null"`
	NewOpponentElo  uint      `gorm:"not null"`

	LeagueId       uuid.UUID `gorm:"not null"`
	TeamId         uuid.UUID `gorm:"not null"`
	OpponentTeamId uuid.UUID `gorm:"not null"`
	UserId         *uuid.UUID

	League       League
	Team         Team
	OpponentTeam Team
	User         User
}
