package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type QueueRepository interface {
	GetAll() ([]*models.Queue, error)
	GetByID(id uuid.UUID) (*models.Queue, error)
	GetByLeagueID(id uuid.UUID) (*models.Queue, error)
	Create(queue *models.Queue) error
	Save(queue *models.Queue) error
	Delete(id uuid.UUID) error
}

type queueRepository struct {
	db *config.DB
}

func NewQueueRepository(db *config.DB) QueueRepository {
	return &queueRepository{db: db}
}

func (r *queueRepository) GetAll() ([]*models.Queue, error) {
	var queues []*models.Queue

	result := r.db.
		Joins("League").
		Preload("Teams.Players").
		Order("created_at DESC").
		Find(&queues)

	return queues, result.Error
}

func (r *queueRepository) GetByID(id uuid.UUID) (*models.Queue, error) {
	var queue models.Queue

	result := r.db.
		Joins("League").
		Preload("Teams.Players").
		First(&queue, "id = ?", id)

	return &queue, result.Error
}

func (r *queueRepository) GetByLeagueID(id uuid.UUID) (*models.Queue, error) {
	var queue models.Queue

	result := r.db.
		Joins("League").
		Preload("Teams.Players").
		First(&queue, "league_id = ?", id)

	return &queue, result.Error
}

func (r *queueRepository) Create(queue *models.Queue) error {
	result := r.db.Create(queue)

	return result.Error
}

func (r *queueRepository) Save(queue *models.Queue) error {
	result := r.db.Save(queue)

	return result.Error
}

func (r *queueRepository) Delete(id uuid.UUID) error {
	var queue models.Queue

	result := r.db.
		Where("id = ?", id).
		Delete(&queue)

	return result.Error
}
