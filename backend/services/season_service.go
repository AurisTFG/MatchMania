package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"

	"github.com/google/uuid"
)

type SeasonService interface {
	GetAllSeasons() ([]models.Season, error)
	GetSeasonByID(uint) (*models.Season, error)
	CreateSeason(*models.CreateSeasonDto, uuid.UUID) (*models.Season, error)
	UpdateSeason(*models.Season, *models.UpdateSeasonDto) (*models.Season, error)
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

func (s *seasonService) CreateSeason(seasonDto *models.CreateSeasonDto, userUUID uuid.UUID) (*models.Season, error) {
	newSeason := seasonDto.ToSeason()
	newSeason.UserUUID = userUUID

	return s.repo.Create(&newSeason)
}

func (s *seasonService) UpdateSeason(
	currentSeason *models.Season,
	updatedSeasonDto *models.UpdateSeasonDto,
) (*models.Season, error) {
	updatedSeason := updatedSeasonDto.ToSeason()

	return s.repo.Update(currentSeason, &updatedSeason)
}

func (s *seasonService) DeleteSeason(season *models.Season) error {
	return s.repo.Delete(season)
}
