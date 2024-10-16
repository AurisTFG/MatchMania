package models

type Player struct {
	ID       uint `gorm:"primaryKey"`
	Nickname string
	Country  string

	TeamID uint
}
