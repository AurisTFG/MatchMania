package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/teams"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"

	"github.com/google/uuid"
)

type TeamService interface {
	GetAllTeams(uuid.UUID) ([]models.Team, error)
	GetTeamById(uuid.UUID, uuid.UUID) (*models.Team, error)
	CreateTeam(*requests.CreateTeamDto, uuid.UUID, uuid.UUID) error
	UpdateTeam(*models.Team, *requests.UpdateTeamDto) error
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
	return s.repo.FindByIdAndSeasonID(seasonId, teamId)
}

func (s *teamService) CreateTeam(
	teamDto *requests.CreateTeamDto,
	seasonId uuid.UUID,
	userId uuid.UUID,
) error {
	newTeam := utils.MustCopy[models.Team](teamDto)
	newTeam.Elo = 1000
	newTeam.SeasonId = seasonId
	newTeam.UserId = userId

	return s.repo.Create(newTeam)
}

func (s *teamService) UpdateTeam(
	currentTeam *models.Team,
	updatedTeamDto *requests.UpdateTeamDto,
) error {
	updatedTeam := utils.MustCopy[models.Team](updatedTeamDto)

	return s.repo.Update(currentTeam, updatedTeam)
}

func (s *teamService) DeleteTeam(team *models.Team) error {
	return s.repo.Delete(team)
}
