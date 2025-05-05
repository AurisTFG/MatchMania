package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
)

type PlayerRepository interface {
	GetAll() ([]models.User, error)
}

type playerRepository struct {
	db *config.DB
}

func NewPlayerRepository(db *config.DB) PlayerRepository {
	return &playerRepository{db: db}
}

func (r *playerRepository) GetAll() ([]models.User, error) {
	var players []models.User

	result := r.db.
		Where("trackmania_name IS NOT NULL AND trackmania_name <> ''").
		Find(&players)

	return players, result.Error
}
