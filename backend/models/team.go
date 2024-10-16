package models

import (
	"time"
)

type Team struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time

	SeasonID uint

	Players     []Player `gorm:"foreignKey:TeamID"`
	HomeResults []Result `gorm:"foreignKey:TeamID"`
	AwayResults []Result `gorm:"foreignKey:OpponentTeamID"`
}
