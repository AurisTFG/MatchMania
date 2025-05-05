package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/teams"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"
	"fmt"

	"github.com/google/uuid"
)

type TeamService interface {
	GetAllTeams() ([]models.Team, error)
	GetAllTeamsByLeagueId(uuid.UUID) ([]models.Team, error)
	GetTeamById(uuid.UUID) (*models.Team, error)
	GetTeamByIdAndLeagueId(uuid.UUID, uuid.UUID) (*models.Team, error)
	CreateTeam(*requests.CreateTeamDto, uuid.UUID) error
	UpdateTeam(*models.Team, *requests.UpdateTeamDto) error
	DeleteTeam(*models.Team) error
}

type teamService struct {
	teamRepository repositories.TeamRepository
}

func NewTeamService(teamRepository repositories.TeamRepository) TeamService {
	return &teamService{teamRepository: teamRepository}
}

func (s *teamService) GetAllTeams() ([]models.Team, error) {
	return s.teamRepository.FindAll()
}

func (s *teamService) GetAllTeamsByLeagueId(leagueId uuid.UUID) ([]models.Team, error) {
	return s.teamRepository.FindAllByLeagueID(leagueId)
}

func (s *teamService) GetTeamById(teamId uuid.UUID) (*models.Team, error) {
	return s.teamRepository.FindById(teamId)
}

func (s *teamService) GetTeamByIdAndLeagueId(leagueId uuid.UUID, teamId uuid.UUID) (*models.Team, error) {
	return s.teamRepository.FindByIdAndLeagueID(leagueId, teamId)
}

func (s *teamService) CreateTeam(
	teamDto *requests.CreateTeamDto,
	userId uuid.UUID,
) error {
	newTeam := utils.MustCopy[models.Team](teamDto)
	newTeam.Elo = 1000
	newTeam.UserId = userId

	if err := s.addLeaguesToTeam(newTeam, teamDto.LeagueIds); err != nil {
		return err
	}

	if err := s.addPlayersToTeam(newTeam, teamDto.PlayerIds); err != nil {
		return err
	}

	return s.teamRepository.Create(newTeam)
}

func (s *teamService) UpdateTeam(
	currentTeam *models.Team,
	updatedTeamDto *requests.UpdateTeamDto,
) error {
	if err := s.teamRepository.ClearAssociations(currentTeam, []string{"Leagues", "Players"}); err != nil {
		return err
	}

	updatedTeam := utils.MustCopy[models.Team](updatedTeamDto)

	if err := s.addLeaguesToTeam(currentTeam, updatedTeamDto.LeagueIds); err != nil {
		return err
	}

	if err := s.addPlayersToTeam(currentTeam, updatedTeamDto.PlayerIds); err != nil {
		return err
	}

	return s.teamRepository.Update(currentTeam, updatedTeam)
}

func (s *teamService) DeleteTeam(team *models.Team) error {
	return s.teamRepository.Delete(team)
}

func (s *teamService) addLeaguesToTeam(team *models.Team, leagueIds []string) error {
	for _, leagueId := range leagueIds {
		parsedLeagueId, err := uuid.Parse(leagueId)
		if err != nil {
			return fmt.Errorf("invalid league ID '%s': %w", leagueId, err)
		}

		team.Leagues = append(team.Leagues, models.League{
			BaseModel: models.BaseModel{Id: parsedLeagueId},
		})
	}
	return nil
}

func (s *teamService) addPlayersToTeam(team *models.Team, playerIds []string) error {
	for _, playerId := range playerIds {
		parsedPlayerId, err := uuid.Parse(playerId)
		if err != nil {
			return fmt.Errorf("invalid player ID '%s': %w", playerId, err)
		}

		team.Players = append(team.Players, models.User{
			BaseModel: models.BaseModel{Id: parsedPlayerId},
		})
	}
	return nil
}
