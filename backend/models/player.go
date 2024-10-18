package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Nickname string
	Country  string

	TeamID uint
}
