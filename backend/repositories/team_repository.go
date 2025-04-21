package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type TeamRepository interface {
	FindAll() ([]models.Team, error)
	FindAllBySeasonID(uuid.UUID) ([]models.Team, error)
	FindById(uuid.UUID) (*models.Team, error)
	FindByIdAndSeasonID(uuid.UUID, uuid.UUID) (*models.Team, error)
	Create(*models.Team) error
	Update(*models.Team, *models.Team) error
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

	result := r.db.Joins("User").Find(&teams)

	return teams, result.Error
}

func (r *teamRepository) FindAllBySeasonID(seasonId uuid.UUID) ([]models.Team, error) {
	var teams []models.Team

	result := r.db.Joins("User").Where("season_id = ?", seasonId).Find(&teams)

	return teams, result.Error
}

func (r *teamRepository) FindById(teamId uuid.UUID) (*models.Team, error) {
	var team models.Team

	result := r.db.Joins("User").First(&team, teamId)

	return &team, result.Error
}

func (r *teamRepository) FindByIdAndSeasonID(seasonId uuid.UUID, teamId uuid.UUID) (*models.Team, error) {
	var team models.Team

	result := r.db.Joins("User").Where("season_id = ? AND \"teams\".\"id\" = ?", seasonId, teamId).First(&team)

	return &team, result.Error
}

func (r *teamRepository) Create(team *models.Team) error {
	result := r.db.Create(team)

	return result.Error
}

func (r *teamRepository) Update(currentTeam *models.Team, updatedTeam *models.Team) error {
	result := r.db.Model(currentTeam).Updates(updatedTeam)

	return result.Error
}

func (r *teamRepository) Delete(team *models.Team) error {
	result := r.db.Select(clause.Associations).Delete(team)

	return result.Error
}
