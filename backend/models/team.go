package models

import (
	"github.com/google/uuid"
)

type Team struct {
	BaseModel

	Name string `gorm:"not null"`
	Elo  uint   `gorm:"not null"`

	UserId   uuid.UUID `gorm:"not null"`
	User     User
	SeasonId uuid.UUID `gorm:"not null"`
	Season   Season

	HomeResults []Result `gorm:"foreignKey:TeamId;constraint:OnDelete:CASCADE"`
	AwayResults []Result `gorm:"foreignKey:OpponentTeamId;constraint:OnDelete:CASCADE"`
}
