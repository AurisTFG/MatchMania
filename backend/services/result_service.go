package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/results"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"

	"github.com/google/uuid"
)

type ResultService interface {
	GetAllResults(seasonId uuid.UUID, teamId uuid.UUID) ([]models.Result, error)
	GetResultById(seasonId uuid.UUID, teamId uuid.UUID, resultId uuid.UUID) (*models.Result, error)
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

func (s *resultService) GetAllResults(seasonId uuid.UUID, teamId uuid.UUID) ([]models.Result, error) {
	return s.repo.FindAllBySeasonIDAndTeamID(seasonId, teamId)
}

func (s *resultService) GetResultById(
	seasonId uuid.UUID,
	teamId uuid.UUID,
	resultId uuid.UUID,
) (*models.Result, error) {
	return s.repo.FindByIdAndSeasonIDAndTeamID(seasonId, teamId, resultId)
}

func (s *resultService) CreateResult(
	resultDto *requests.CreateResultDto,
	seasonId uuid.UUID,
	teamId uuid.UUID,
	userId uuid.UUID,
) error {
	newResult := utils.CopyOrPanic[models.Result](resultDto)
	newResult.SeasonId = seasonId
	newResult.TeamId = teamId
	newResult.UserId = userId

	return s.repo.Create(newResult)
}

func (s *resultService) UpdateResult(
	currentResult *models.Result,
	updatedResultDto *requests.UpdateResultDto,
) error {
	updatedResult := utils.CopyOrPanic[models.Result](updatedResultDto)

	return s.repo.Update(currentResult, updatedResult)
}

func (s *resultService) DeleteResult(resultModel *models.Result) error {
	return s.repo.Delete(resultModel)
}
