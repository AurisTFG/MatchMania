package models

import (
	"MatchManiaAPI/models/enums"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel

	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	// Country   string // TODO: Add country to user
	Role     enums.Role `gorm:"type:role;default:'user'"`
	Password string

	Sessions []Session `gorm:"constraint:OnDelete:CASCADE;"`
	Seasons  []Season  `gorm:"constraint:OnDelete:CASCADE;"`
	Teams    []Team    `gorm:"constraint:OnDelete:CASCADE;"`
	Results  []Result  `gorm:"constraint:OnDelete:CASCADE;"`
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
