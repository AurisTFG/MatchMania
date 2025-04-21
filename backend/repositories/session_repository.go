package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type SessionRepository interface {
	FindAll() ([]models.Session, error)
	FindById(uuid.UUID) (*models.Session, error)
	Create(*models.Session) error
	Update(*models.Session) error
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

func (r *sessionRepository) FindById(sessionId uuid.UUID) (*models.Session, error) {
	var session models.Session

	result := r.db.First(&session, "id = ?", sessionId)

	return &session, result.Error
}

func (r *sessionRepository) Create(session *models.Session) error {
	result := r.db.Create(session)

	return result.Error
}

func (r *sessionRepository) Update(session *models.Session) error {
	result := r.db.Model(&models.Session{}).Where("id = ?", session.Id).Updates(session)

	return result.Error
}

func (r *sessionRepository) Delete(session *models.Session) error {
	result := r.db.Delete(session)

	return result.Error
}
