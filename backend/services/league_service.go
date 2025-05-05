package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/leagues"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"
	"fmt"

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

	if err := s.addTracksToLeague(newLeague, leagueDto.TrackIds); err != nil {
		return err
	}

	return s.repo.Create(newLeague)
}

func (s *leagueService) UpdateLeague(
	currentLeague *models.League,
	updatedLeagueDto *requests.UpdateLeagueDto,
) error {
	if err := s.repo.ClearAssociations(currentLeague, []string{"Tracks"}); err != nil {
		return err
	}

	updatedLeague := utils.MustCopy[models.League](updatedLeagueDto)

	if err := s.addTracksToLeague(currentLeague, updatedLeagueDto.TrackIds); err != nil {
		return err
	}

	return s.repo.Update(currentLeague, updatedLeague)
}

func (s *leagueService) DeleteLeague(league *models.League) error {
	return s.repo.Delete(league)
}

func (s *leagueService) addTracksToLeague(league *models.League, trackIds []string) error {
	for _, trackId := range trackIds {
		parsedTrackId, err := uuid.Parse(trackId)
		if err != nil {
			return fmt.Errorf("invalid track ID '%s': %w", trackId, err)
		}

		league.Tracks = append(league.Tracks, models.TrackmaniaTrack{
			Id: parsedTrackId,
		})
	}
	return nil
}
