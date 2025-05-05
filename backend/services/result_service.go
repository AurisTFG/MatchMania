package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/results"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"
	"strconv"

	"github.com/google/uuid"
)

type ResultService interface {
	GetAllResults() ([]models.Result, error)
	GetResultById(uuid.UUID) (*models.Result, error)
	CreateResult(*requests.CreateResultDto, uuid.UUID) error
	UpdateResult(*models.Result, *requests.UpdateResultDto) error
	DeleteResult(*models.Result) error
}

type resultService struct {
	repo repositories.ResultRepository
}

func NewResultService(repo repositories.ResultRepository) ResultService {
	return &resultService{repo: repo}
}

func (s *resultService) GetAllResults() ([]models.Result, error) {
	return s.repo.GetAll()
}

func (s *resultService) GetResultById(resultId uuid.UUID) (*models.Result, error) {
	return s.repo.GetById(resultId)
}

func (s *resultService) CreateResult(
	resultDto *requests.CreateResultDto,
	userId uuid.UUID,
) error {
	newResult := utils.MustCopy[models.Result](resultDto)

	scoreUint, err := strconv.ParseUint(resultDto.Score, 10, 32)
	if err != nil {
		return err
	}
	opponentScoreUint, err := strconv.ParseUint(resultDto.OpponentScore, 10, 32)
	if err != nil {
		return err
	}

	newResult.Score = uint(scoreUint)
	newResult.OpponentScore = uint(opponentScoreUint)
	newResult.UserId = &userId

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
