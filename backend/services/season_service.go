package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"
)

type SeasonService interface {
	GetAllSeasons() ([]models.Season, error)
	GetSeasonByID(uint) (*models.Season, error)
	CreateSeason(*models.CreateSeasonDto) (*models.Season, error)
	UpdateSeason(uint, *models.UpdateSeasonDto) (*models.Season, error)
	DeleteSeason(*models.Season) error
}

type seasonService struct {
	repo repositories.SeasonRepository
}

func NewSeasonService(repo repositories.SeasonRepository) SeasonService {
	return &seasonService{repo: repo}
}

func (s *seasonService) GetAllSeasons() ([]models.Season, error) {
	return s.repo.FindAll()
}

func (s *seasonService) GetSeasonByID(seasonID uint) (*models.Season, error) {
	return s.repo.FindByID(seasonID)
}

func (s *seasonService) CreateSeason(seasonDto *models.CreateSeasonDto) (*models.Season, error) {
	newSeason := seasonDto.ToSeason()

	return s.repo.Create(&newSeason)
}

func (s *seasonService) UpdateSeason(seasonID uint, seasonDto *models.UpdateSeasonDto) (*models.Season, error) {
	updatedSeason := seasonDto.ToSeason()

	return s.repo.Update(seasonID, &updatedSeason)
}

func (s *seasonService) DeleteSeason(season *models.Season) error {
	return s.repo.Delete(season)
}
