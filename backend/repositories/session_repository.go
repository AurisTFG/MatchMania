package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
)

type SessionRepository interface {
	FindAll() ([]models.Session, error)
	FindByID(string) (*models.Session, error)
	Create(*models.Session) (*models.Session, error)
	Update(*models.Session) (*models.Session, error)
	Delete(*models.Session) error
}

type sessionRepository struct {
	db *config.DB
}

func NewSessionRepository(db *config.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) FindAll() ([]models.Session, error) {
	var sessions []models.Session

	result := r.db.Find(&sessions)

	return sessions, result.Error
}

func (r *sessionRepository) FindByID(sessionID string) (*models.Session, error) {
	var session models.Session

	result := r.db.First(&session, "uuid = ?", sessionID)

	return &session, result.Error
}

func (r *sessionRepository) Create(session *models.Session) (*models.Session, error) {
	result := r.db.Create(session)

	return session, result.Error
}

func (r *sessionRepository) Update(session *models.Session) (*models.Session, error) {
	result := r.db.Model(&models.Session{}).Where("uuid = ?", session.UUID).Updates(session)

	return session, result.Error
}

func (r *sessionRepository) Delete(session *models.Session) error {
	result := r.db.Delete(session)

	return result.Error
}
