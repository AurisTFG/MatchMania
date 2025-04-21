package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/seasons"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"

	"github.com/google/uuid"
)

type SeasonService interface {
	GetAllSeasons() ([]models.Season, error)
	GetSeasonById(uuid.UUID) (*models.Season, error)
	CreateSeason(*requests.CreateSeasonDto, uuid.UUID) error
	UpdateSeason(*models.Season, *requests.UpdateSeasonDto) error
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

func (s *seasonService) GetSeasonById(seasonId uuid.UUID) (*models.Season, error) {
	return s.repo.FindById(seasonId)
}

func (s *seasonService) CreateSeason(seasonDto *requests.CreateSeasonDto, userId uuid.UUID) error {
	newSeason := utils.CopyOrPanic[models.Season](seasonDto)
	newSeason.UserId = userId

	return s.repo.Create(newSeason)
}

func (s *seasonService) UpdateSeason(
	currentSeason *models.Season,
	updatedSeasonDto *requests.UpdateSeasonDto,
) error {
	updatedSeason := utils.CopyOrPanic[models.Season](updatedSeasonDto)

	return s.repo.Update(currentSeason, updatedSeason)
}

func (s *seasonService) DeleteSeason(season *models.Season) error {
	return s.repo.Delete(season)
}
