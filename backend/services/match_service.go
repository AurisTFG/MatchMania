package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/dtos/requests/results"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type MatchService interface {
	EndMatch(playerId uuid.UUID) error
	GetAllMatches() ([]*responses.MatchDto, error)
	GetMatchByID(uuid.UUID) (*responses.MatchDto, error)
	CreateMatch(*models.Match) error
	SaveMatch(*models.Match) error
	DeleteMatch(uuid.UUID) error
}

type matchService struct {
	matchRepository      repositories.MatchRepository
	resultService        ResultService
	trackmaniaApiService TrackmaniaApiService
}

func NewMatchService(
	matchRepository repositories.MatchRepository,
	resultService ResultService,
	trackmaniaApiService TrackmaniaApiService,
) MatchService {
	return &matchService{
		matchRepository:      matchRepository,
		resultService:        resultService,
		trackmaniaApiService: trackmaniaApiService,
	}
}

func (s *matchService) EndMatch(playerId uuid.UUID) error {
	matches, err := s.matchRepository.GetAll()
	if err != nil {
		return err
	}

	match, err := findMatchByPlayerId(matches, playerId)
	if err != nil {
		return err
	}

	if len(match.Teams) != 2 {
		return errors.New("wrong number of teams, expected 2, got " + strconv.Itoa(len(match.Teams)))
	}

	teamsResultDto, err := s.trackmaniaApiService.GetTeamsResults(match.TrackmaniaCompetitionId)
	if err != nil {
		return fmt.Errorf("getting teams results: %w", err)
	}

	if err = validateTeamsResultDto(teamsResultDto, match); err != nil {
		return fmt.Errorf("validating teams results: %w", err)
	}

	createResultDto := &results.CreateResultDto{
		LeagueId:       match.LeagueId,
		StartDate:      match.CreatedAt,
		EndDate:        time.Now(),
		TeamId:         uuid.MustParse(teamsResultDto.Teams[0].TeamId),
		OpponentTeamId: uuid.MustParse(teamsResultDto.Teams[1].TeamId),
		Score:          strconv.Itoa(teamsResultDto.Teams[0].Score),
		OpponentScore:  strconv.Itoa(teamsResultDto.Teams[1].Score),
	}

	if err = s.resultService.CreateResult(createResultDto, nil); err != nil {
		return err
	}

	err = s.trackmaniaApiService.DeleteCompetition(match.TrackmaniaCompetitionId)
	if err != nil {
		return fmt.Errorf("deleting competition: %w", err)
	}

	if err = s.matchRepository.Delete(match.Id); err != nil {
		return err
	}

	return nil
}

func (s *matchService) GetAllMatches() ([]*responses.MatchDto, error) {
	matches, err := s.matchRepository.GetAll()
	if err != nil {
		return nil, err
	}

	matchDtos := utils.MustCopy[[]*responses.MatchDto](matches)

	return *matchDtos, nil
}

func (s *matchService) GetMatchByID(id uuid.UUID) (*responses.MatchDto, error) {
	match, err := s.matchRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	matchDto := utils.MustCopy[responses.MatchDto](match)

	return matchDto, nil
}

func (s *matchService) CreateMatch(match *models.Match) error {
	err := s.matchRepository.Create(match)
	if err != nil {
		return err
	}

	return nil
}

func (s *matchService) SaveMatch(match *models.Match) error {
	err := s.matchRepository.Save(match)
	if err != nil {
		return err
	}

	return nil
}

func (s *matchService) DeleteMatch(id uuid.UUID) error {
	match, err := s.matchRepository.GetByID(id)
	if err != nil {
		return err
	}

	err = s.matchRepository.Delete(match.Id)
	if err != nil {
		return err
	}

	return nil
}

func findMatchByPlayerId(matches []*models.Match, playerId uuid.UUID) (*models.Match, error) {
	for _, match := range matches {
		for _, team := range match.Teams {
			for _, player := range team.Players {
				if player.Id == playerId {
					return match, nil
				}
			}
		}
	}

	return nil, errors.New("match not found")
}

func validateTeamsResultDto(teamsResultDto *responses.TeamsResultsDto, match *models.Match) error {
	if len(teamsResultDto.Teams) != 2 {
		return errors.New("wrong number of teams results, expected 2, got " + strconv.Itoa(len(teamsResultDto.Teams)))
	}

	if match.Teams[0].Id.String() != teamsResultDto.Teams[0].TeamId &&
		match.Teams[0].Id.String() != teamsResultDto.Teams[1].TeamId {
		return errors.New("team " + match.Teams[0].Name + " not found in teams results")
	}

	if match.Teams[1].Id.String() != teamsResultDto.Teams[0].TeamId &&
		match.Teams[1].Id.String() != teamsResultDto.Teams[1].TeamId {
		return errors.New("team " + match.Teams[1].Name + " not found in teams results")
	}

	if teamsResultDto.Teams[0].Score == 0 && teamsResultDto.Teams[1].Score == 0 {
		return errors.New("both teams have 0 score")
	}

	if teamsResultDto.Teams[0].Score == teamsResultDto.Teams[1].Score {
		return errors.New("both teams have the same score")
	}

	return nil
}
