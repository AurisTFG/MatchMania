package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"
)

type TeamService interface {
	GetAllTeams(uint) ([]models.Team, error)
	GetTeamByID(uint, uint) (*models.Team, error)
	CreateTeam(*models.CreateTeamDto, uint) (*models.Team, error)
	UpdateTeam(uint, *models.UpdateTeamDto) (*models.Team, error)
	DeleteTeam(*models.Team) error
}

type teamService struct {
	repo repositories.TeamRepository
}

func NewTeamService(repo repositories.TeamRepository) TeamService {
	return &teamService{repo: repo}
}

func (s *teamService) GetAllTeams(seasonID uint) ([]models.Team, error) {
	return s.repo.FindAllBySeasonID(seasonID)
}

func (s *teamService) GetTeamByID(seasonID uint, teamID uint) (*models.Team, error) {
	return s.repo.FindByIDAndSeasonID(teamID, seasonID)
}

func (s *teamService) CreateTeam(teamDto *models.CreateTeamDto, seasonID uint) (*models.Team, error) {
	newTeam := teamDto.ToTeam()
	newTeam.SeasonID = seasonID
	newTeam.Elo = 1000

	return s.repo.Create(&newTeam)
}

func (s *teamService) UpdateTeam(teamID uint, teamDto *models.UpdateTeamDto) (*models.Team, error) {
	updatedTeam := teamDto.ToTeam()
	updatedTeam.ID = teamID

	return s.repo.Update(&updatedTeam)
}

func (s *teamService) DeleteTeam(team *models.Team) error {
	return s.repo.Delete(team)
}