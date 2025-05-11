package models

import (
	"time"

	"github.com/google/uuid"
)

type TrackmaniaOauthState struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time `gorm:"not null"`
	State     string    `gorm:"not null;unique"`

	UserId uuid.UUID `gorm:"not null;unique"`

	User User
}
