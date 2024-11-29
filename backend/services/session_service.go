package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"
	"time"

	"github.com/google/uuid"
)

type SessionService interface {
	CreateSession(sessionUUID uuid.UUID, userUUID uuid.UUID, refreshToken string, expiresAt time.Time) error
	ExtendSession(sessionUUID string, refreshToken string, expiresAt time.Time) error
	InvalidateSession(sessionUUID string) error
	IsSessionValid(sessionUUID string, refreshToken string) bool
}

type sessionService struct {
	repo repositories.SessionRepository
}

func NewSessionService(repo repositories.SessionRepository) SessionService {
	return &sessionService{repo: repo}
}

func (s *sessionService) CreateSession(sessionUUID uuid.UUID, userUUID uuid.UUID, refreshToken string, expiresAt time.Time) error {
	session := &models.Session{
		UUID:             sessionUUID,
		UserUUID:         userUUID,
		LastRefreshToken: refreshToken,
		ExpiresAt:        expiresAt,
		InitiatedAt:      time.Now(),
	}

	if err := session.HashToken(); err != nil {
		return err
	}

	_, err := s.repo.Create(session)
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionService) ExtendSession(sessionUUID string, refreshToken string, expiresAt time.Time) error {
	session, err := s.repo.FindByID(sessionUUID)
	if err != nil {
		return err
	}

	session.LastRefreshToken = refreshToken
	session.ExpiresAt = expiresAt

	if err := session.HashToken(); err != nil {
		return err
	}

	_, err = s.repo.Update(session)
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionService) InvalidateSession(sessionUUID string) error {
	session, err := s.repo.FindByID(sessionUUID)
	if err != nil {
		return err
	}

	session.IsRevoked = true

	_, err = s.repo.Update(session)
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionService) IsSessionValid(sessionUUID string, refreshToken string) bool {
	session, err := s.repo.FindByID(sessionUUID)
	if err != nil {
		return false
	}

	isTokenValid := session.CompareToken(refreshToken)
	if !isTokenValid {
		return false
	}

	if session.IsRevoked {
		return false
	}

	if session.ExpiresAt.Before(time.Now()) {
		return false
	}

	return true
}
