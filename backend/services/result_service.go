package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"

	"github.com/google/uuid"
)

type ResultService interface {
	GetAllResults(uint, uint) ([]models.Result, error)
	GetResultByID(uint, uint, uint) (*models.Result, error)
	CreateResult(*models.CreateResultDto, uint, uint, uuid.UUID) (*models.Result, error)
	UpdateResult(*models.Result, *models.UpdateResultDto) (*models.Result, error)
	DeleteResult(*models.Result) error
}

type resultService struct {
	repo repositories.ResultRepository
}

func NewResultService(repo repositories.ResultRepository) ResultService {
	return &resultService{repo: repo}
}

func (s *resultService) GetAllResults(seasonID uint, teamID uint) ([]models.Result, error) {
	return s.repo.FindAllBySeasonIDAndTeamID(seasonID, teamID)
}

func (s *resultService) GetResultByID(seasonID uint, teamID uint, resultID uint) (*models.Result, error) {
	return s.repo.FindByIDAndSeasonIDAndTeamID(seasonID, teamID, resultID)
}

func (s *resultService) CreateResult(resultDto *models.CreateResultDto, seasonID uint, teamID uint, userUUID uuid.UUID) (*models.Result, error) {
	newResult := resultDto.ToResult()
	newResult.SeasonID = seasonID
	newResult.TeamID = teamID
	newResult.UserUUID = userUUID

	return s.repo.Create(&newResult)
}

func (s *resultService) UpdateResult(currentResult *models.Result, updatedResultDto *models.UpdateResultDto) (*models.Result, error) {
	updatedResult := updatedResultDto.ToResult()

	return s.repo.Update(currentResult, &updatedResult)
}

func (s *resultService) DeleteResult(resultModel *models.Result) error {
	return s.repo.Delete(resultModel)
}
