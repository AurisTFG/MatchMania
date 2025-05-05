package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"
)

type PlayerService interface {
	GetAllPlayers() ([]models.User, error)
}

type playerService struct {
	repo repositories.PlayerRepository
}

func NewPlayerService(repo repositories.PlayerRepository) PlayerService {
	return &playerService{repo: repo}
}

func (s *playerService) GetAllPlayers() ([]models.User, error) {
	players, err := s.repo.GetAll()

	return players, err
}
