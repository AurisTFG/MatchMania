package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/dtos/requests/results"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"
	"errors"
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
	matchRepository repositories.MatchRepository
	resultService   ResultService
}

func NewMatchService(
	matchRepository repositories.MatchRepository,
	resultService ResultService,
) MatchService {
	return &matchService{
		matchRepository: matchRepository,
		resultService:   resultService,
	}
}

func (s *matchService) EndMatch(playerId uuid.UUID) error {
	matches, err := s.matchRepository.GetAll()
	if err != nil {
		return err
	}

	var match *models.Match
	for _, m := range matches {
		for _, team := range m.Teams {
			for _, player := range team.Players {
				if player.Id == playerId {
					match = m
					break
				}
			}
			if match != nil {
				break
			}
		}
		if match != nil {
			break
		}
	}

	if match == nil {
		return errors.New("match not found")
	}

	if len(match.Teams) != 2 {
		return errors.New("wrong number of teams, expected 2, got " + strconv.Itoa(len(match.Teams)))
	}

	createResultDto := &results.CreateResultDto{
		LeagueId:       match.LeagueId,
		StartDate:      match.CreatedAt,
		EndDate:        time.Now(),
		TeamId:         match.Teams[0].Id,
		OpponentTeamId: match.Teams[1].Id,
		Score:          "3",
		OpponentScore:  "4",
	}

	if err := s.resultService.CreateResult(createResultDto, nil); err != nil {
		return err
	}

	if err := s.matchRepository.Delete(match.Id); err != nil {
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
