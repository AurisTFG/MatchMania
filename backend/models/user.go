package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel

	Username       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	TrackmaniaId   uuid.UUID
	TrackmaniaName string
	Country        string
	Password       string

	Sessions              []Session
	Seasons               []Season
	Teams                 []Team
	Results               []Result
	TrackmaniaOauthTracks []TrackmaniaOauthTrack

	Roles []Role `gorm:"many2many:user_roles"`
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
