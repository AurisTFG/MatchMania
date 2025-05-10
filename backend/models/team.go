package models

import (
	"github.com/google/uuid"
)

type Team struct {
	BaseModel
	Name    string `gorm:"not null"`
	LogoUrl string
	Elo     uint `gorm:"not null"`

	MatchId *uuid.UUID
	QueueId *uuid.UUID
	UserId  uuid.UUID `gorm:"not null"`

	User        User
	Match       Match
	Queue       Queue
	Players     []User   `gorm:"many2many:team_players;"`
	Leagues     []League `gorm:"many2many:league_teams;"`
	HomeResults []Result `gorm:"foreignKey:TeamId"`
	AwayResults []Result `gorm:"foreignKey:OpponentTeamId"`
}
