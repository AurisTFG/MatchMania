package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/leagues"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"

	"github.com/google/uuid"
)

type LeagueService interface {
	GetAllLeagues() ([]models.League, error)
	GetLeagueById(uuid.UUID) (*models.League, error)
	CreateLeague(*requests.CreateLeagueDto, uuid.UUID) error
	UpdateLeague(*models.League, *requests.UpdateLeagueDto) error
	DeleteLeague(*models.League) error
}

type leagueService struct {
	repo repositories.LeagueRepository
}

func NewLeagueService(repo repositories.LeagueRepository) LeagueService {
	return &leagueService{repo: repo}
}

func (s *leagueService) GetAllLeagues() ([]models.League, error) {
	return s.repo.FindAll()
}

func (s *leagueService) GetLeagueById(leagueId uuid.UUID) (*models.League, error) {
	return s.repo.FindById(leagueId)
}

func (s *leagueService) CreateLeague(leagueDto *requests.CreateLeagueDto, userId uuid.UUID) error {
	newLeague := utils.MustCopy[models.League](leagueDto)
	newLeague.UserId = userId

	return s.repo.Create(newLeague)
}

func (s *leagueService) UpdateLeague(
	currentLeague *models.League,
	updatedLeagueDto *requests.UpdateLeagueDto,
) error {
	updatedLeague := utils.MustCopy[models.League](updatedLeagueDto)

	return s.repo.Update(currentLeague, updatedLeague)
}

func (s *leagueService) DeleteLeague(league *models.League) error {
	return s.repo.Delete(league)
}
