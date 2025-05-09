package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type RoleRepository interface {
	GetAll() ([]models.Role, error)
	GetById(uuid.UUID) (*models.Role, error)
	GetByName(string) (*models.Role, error)
	Create(*models.Role) error
	Update(*models.Role, *models.Role) error
	Delete(*models.Role) error
}

type roleRepository struct {
	db *config.DB
}

func NewRoleRepository(db *config.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetAll() ([]models.Role, error) {
	var roles []models.Role

	result := r.db.
		Preload("Permissions").
		Find(&roles)

	return roles, result.Error
}

func (r *roleRepository) GetById(roleId uuid.UUID) (*models.Role, error) {
	var role models.Role

	result := r.db.First(&role, "id = ?", roleId)

	return &role, result.Error
}

func (r *roleRepository) GetByName(roleName string) (*models.Role, error) {
	var role models.Role

	result := r.db.
		Where("name = ?", roleName).
		First(&role)

	return &role, result.Error
}

func (r *roleRepository) Create(newRole *models.Role) error {
	result := r.db.Create(newRole)

	return result.Error
}

func (r *roleRepository) Update(currentRole *models.Role, updatedRole *models.Role) error {
	result := r.db.
		Model(currentRole).
		Updates(updatedRole)

	return result.Error
}

func (r *roleRepository) Delete(roleModel *models.Role) error {
	result := r.db.Delete(roleModel)

	return result.Error
}
