package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/results"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"

	"github.com/google/uuid"
)

type ResultService interface {
	GetAllResults(leagueId uuid.UUID, teamId uuid.UUID) ([]models.Result, error)
	GetResultById(leagueId uuid.UUID, teamId uuid.UUID, resultId uuid.UUID) (*models.Result, error)
	CreateResult(*requests.CreateResultDto, uuid.UUID, uuid.UUID, uuid.UUID) error
	UpdateResult(*models.Result, *requests.UpdateResultDto) error
	DeleteResult(*models.Result) error
}

type resultService struct {
	repo repositories.ResultRepository
}

func NewResultService(repo repositories.ResultRepository) ResultService {
	return &resultService{repo: repo}
}

func (s *resultService) GetAllResults(leagueId uuid.UUID, teamId uuid.UUID) ([]models.Result, error) {
	return s.repo.FindAllByLeagueIDAndTeamID(leagueId, teamId)
}

func (s *resultService) GetResultById(
	leagueId uuid.UUID,
	teamId uuid.UUID,
	resultId uuid.UUID,
) (*models.Result, error) {
	return s.repo.FindByIdAndLeagueIDAndTeamID(leagueId, teamId, resultId)
}

func (s *resultService) CreateResult(
	resultDto *requests.CreateResultDto,
	leagueId uuid.UUID,
	teamId uuid.UUID,
	userId uuid.UUID,
) error {
	newResult := utils.MustCopy[models.Result](resultDto)
	newResult.LeagueId = leagueId
	newResult.TeamId = teamId
	newResult.UserId = userId

	return s.repo.Create(newResult)
}

func (s *resultService) UpdateResult(
	currentResult *models.Result,
	updatedResultDto *requests.UpdateResultDto,
) error {
	updatedResult := utils.MustCopy[models.Result](updatedResultDto)

	return s.repo.Update(currentResult, updatedResult)
}

func (s *resultService) DeleteResult(resultModel *models.Result) error {
	return s.repo.Delete(resultModel)
}
