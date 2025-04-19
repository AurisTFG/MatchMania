package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/teams"
	"MatchManiaAPI/repositories"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type TeamService interface {
	GetAllTeams(uuid.UUID) ([]models.Team, error)
	GetTeamById(uuid.UUID, uuid.UUID) (*models.Team, error)
	CreateTeam(*requests.CreateTeamDto, uuid.UUID, uuid.UUID) (*models.Team, error)
	UpdateTeam(*models.Team, *requests.UpdateTeamDto) (*models.Team, error)
	DeleteTeam(*models.Team) error
}

type teamService struct {
	repo repositories.TeamRepository
}

func NewTeamService(repo repositories.TeamRepository) TeamService {
	return &teamService{repo: repo}
}

func (s *teamService) GetAllTeams(seasonId uuid.UUID) ([]models.Team, error) {
	return s.repo.FindAllBySeasonID(seasonId)
}

func (s *teamService) GetTeamById(seasonId uuid.UUID, teamId uuid.UUID) (*models.Team, error) {
	return s.repo.FindByIdAndSeasonID(teamId, seasonId)
}

func (s *teamService) CreateTeam(
	teamDto *requests.CreateTeamDto,
	seasonId uuid.UUID,
	userId uuid.UUID,
) (*models.Team, error) {
	var newTeam models.Team

	copier.Copy(&newTeam, teamDto)
	newTeam.SeasonId = seasonId
	newTeam.Elo = 1000
	newTeam.UserId = userId

	return s.repo.Create(&newTeam)
}

func (s *teamService) UpdateTeam(currentTeam *models.Team, updatedTeamDto *requests.UpdateTeamDto) (*models.Team, error) {
	var updatedTeam models.Team

	copier.Copy(&updatedTeam, updatedTeamDto)

	return s.repo.Update(currentTeam, &updatedTeam)
}

func (s *teamService) DeleteTeam(team *models.Team) error {
	return s.repo.Delete(team)
}
