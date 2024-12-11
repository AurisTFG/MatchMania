package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"gorm.io/gorm/clause"
)

type TeamRepository interface {
	FindAll() ([]models.Team, error)
	FindAllBySeasonID(uint) ([]models.Team, error)
	FindByID(uint) (*models.Team, error)
	FindByIDAndSeasonID(uint, uint) (*models.Team, error)
	Create(*models.Team) (*models.Team, error)
	Update(*models.Team, *models.Team) (*models.Team, error)
	Delete(*models.Team) error
}

type teamRepository struct {
	db *config.DB
}

func NewTeamRepository(db *config.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) FindAll() ([]models.Team, error) {
	var teams []models.Team

	result := r.db.Find(&teams)

	return teams, result.Error
}

func (r *teamRepository) FindAllBySeasonID(seasonID uint) ([]models.Team, error) {
	var teams []models.Team

	result := r.db.Where("season_id = ?", seasonID).Find(&teams)

	return teams, result.Error
}

func (r *teamRepository) FindByID(teamID uint) (*models.Team, error) {
	var team models.Team

	result := r.db.First(&team, teamID)

	return &team, result.Error
}

func (r *teamRepository) FindByIDAndSeasonID(teamID uint, seasonID uint) (*models.Team, error) {
	var team models.Team

	result := r.db.Where("season_id = ? AND id = ?", seasonID, teamID).First(&team)

	return &team, result.Error
}

func (r *teamRepository) Create(team *models.Team) (*models.Team, error) {
	result := r.db.Create(team)

	return team, result.Error
}

func (r *teamRepository) Update(currentTeam *models.Team, updatedTeam *models.Team) (*models.Team, error) {
	result := r.db.Model(currentTeam).Updates(updatedTeam)

	return currentTeam, result.Error
}

func (r *teamRepository) Delete(team *models.Team) error {
	result := r.db.Select(clause.Associations).Delete(team)

	return result.Error
}
