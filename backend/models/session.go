package models

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

type Session struct {
	BaseModel

	LastRefreshToken string    `gorm:"not null"`
	InitiatedAt      time.Time `gorm:"not null"`
	ExpiresAt        time.Time `gorm:"not null"`
	IsRevoked        bool      `gorm:"not null"`

	UserId uuid.UUID `gorm:"not null"`
	User   User
}

func (s *Session) HashToken() error {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}

	hash := argon2.IDKey([]byte(s.LastRefreshToken), salt, 1, 64*1024, 4, 32)

	var buf bytes.Buffer
	buf.Write(salt)
	buf.Write(hash)
	s.LastRefreshToken = base64.StdEncoding.EncodeToString(buf.Bytes())

	return nil
}

func (s *Session) CompareToken(newToken string) bool {
	data, err := base64.StdEncoding.DecodeString(s.LastRefreshToken)
	if err != nil {
		return false
	}

	salt := data[:16]
	storedHash := data[16:]

	newTokenHash := argon2.IDKey([]byte(newToken), salt, 1, 64*1024, 4, 32)

	return string(newTokenHash) == string(storedHash)
}
