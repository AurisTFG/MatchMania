package models

import (
	"time"
)

type Result struct {
	ID             uint `gorm:"primaryKey"`
	MatchStartDate time.Time
	MatchEndDate   time.Time
	Score          uint
	OpponentScore  uint

	TeamID         uint `gorm:"not null"`
	Team           Team `gorm:"foreignKey:TeamID"`
	OpponentTeamID uint `gorm:"not null"`
	OpponentTeam   Team `gorm:"foreignKey:OpponentTeamID"`
}
