package models

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

type Session struct {
	UUID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Device           string    `gorm:"not null"`
	LastRefreshToken string    `gorm:"not null"`
	InitiatedAt      time.Time `gorm:"not null"`
	ExpiresAt        time.Time `gorm:"not null"`
	IsRevoked        bool      `gorm:"not null"`

	UserUUID uuid.UUID `gorm:"type:uuid;not null"`
}

func (s *Session) HashToken() error {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}

	hash := argon2.IDKey([]byte(s.LastRefreshToken), salt, 1, 64*1024, 4, 32)

	s.LastRefreshToken = base64.StdEncoding.EncodeToString(append(salt, hash...))

	return nil
}

func (s *Session) CompareToken(newToken string) bool {
	data, err := base64.StdEncoding.DecodeString(s.LastRefreshToken)
	if err != nil {
		log.Println("Error decoding token:", err)
		return false
	}

	salt := data[:16]
	storedHash := data[16:]

	newTokenHash := argon2.IDKey([]byte(newToken), salt, 1, 64*1024, 4, 32)

	return string(newTokenHash) == string(storedHash)
}
