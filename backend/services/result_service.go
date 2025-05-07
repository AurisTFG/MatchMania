//nolint:gosec // G115: suppress uint to int conversion warning for this file
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
	CreateResult(*requests.CreateResultDto, *uuid.UUID) error
	UpdateResult(*models.Result, *requests.UpdateResultDto) error
	DeleteResult(*models.Result) error
}

type resultService struct {
	resultRepository repositories.ResultRepository
	teamRepository   repositories.TeamRepository
	eloService       EloService
}

func NewResultService(
	resultRepository repositories.ResultRepository,
	teamRepository repositories.TeamRepository,
	eloService EloService,
) ResultService {
	return &resultService{
		resultRepository: resultRepository,
		teamRepository:   teamRepository,
		eloService:       eloService,
	}
}

func (s *resultService) GetAllResults() ([]models.Result, error) {
	return s.resultRepository.GetAll()
}

func (s *resultService) GetResultById(resultId uuid.UUID) (*models.Result, error) {
	return s.resultRepository.GetById(resultId)
}

func (s *resultService) CreateResult(
	resultDto *requests.CreateResultDto,
	userId *uuid.UUID,
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

	team, err := s.teamRepository.FindById(resultDto.TeamId)
	if err != nil {
		return err
	}

	opponentTeam, err := s.teamRepository.FindById(resultDto.OpponentTeamId)
	if err != nil {
		return err
	}

	newElo, newOpponentElo := s.eloService.CalculateElo(
		team.Elo,
		opponentTeam.Elo,
		uint(scoreUint),
		uint(opponentScoreUint),
	)

	eloDiff := int(newElo) - int(team.Elo)
	opponentEloDiff := int(newOpponentElo) - int(opponentTeam.Elo)

	team.Elo = newElo
	opponentTeam.Elo = newOpponentElo

	if err = s.teamRepository.Save(team); err != nil {
		return err
	}

	if err = s.teamRepository.Save(opponentTeam); err != nil {
		return err
	}

	newResult.Score = uint(scoreUint)
	newResult.OpponentScore = uint(opponentScoreUint)
	newResult.EloDiff = eloDiff
	newResult.OpponentEloDiff = opponentEloDiff
	newResult.NewElo = newElo
	newResult.NewOpponentElo = newOpponentElo
	newResult.UserId = userId

	return s.resultRepository.Create(newResult)
}

func (s *resultService) UpdateResult(
	currentResult *models.Result,
	updatedResultDto *requests.UpdateResultDto,
) error {
	updatedResult := utils.MustCopy[models.Result](updatedResultDto)

	return s.resultRepository.Update(currentResult, updatedResult)
}

func (s *resultService) DeleteResult(resultModel *models.Result) error {
	return s.resultRepository.Delete(resultModel)
}
