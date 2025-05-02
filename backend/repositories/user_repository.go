package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"slices"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(uuid.UUID) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	GetDistinctPermissionsByUserId(uuid.UUID) ([]string, error)
	Create(*models.User) error
	Update(*models.User, *models.User) error
	Save(*models.User) error
	Delete(*models.User) error
}

type userRepository struct {
	db *config.DB
}

func NewUserRepository(db *config.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User

	result := r.db.Find(&users)

	return users, result.Error
}

func (r *userRepository) FindById(userId uuid.UUID) (*models.User, error) {
	var user models.User

	result := r.db.Preload("TrackmaniaTracks").Preload("Roles").First(&user, "id = ?", userId)

	return &user, result.Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	result := r.db.Where("email = ?", email).First(&user)

	return &user, result.Error
}

func (r *userRepository) GetDistinctPermissionsByUserId(userId uuid.UUID) ([]string, error) {
	var user models.User

	result := r.db.Preload("Roles.Permissions").First(&user, "id = ?", userId)

	if result.Error != nil {
		return nil, result.Error
	}

	permissions := []string{}
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			if !slices.Contains(permissions, permission.Name) {
				permissions = append(permissions, permission.Name)
			}
		}
	}

	return permissions, result.Error
}

func (r *userRepository) Create(user *models.User) error {
	if err := user.HashPassword(); err != nil {
		return err
	}

	result := r.db.Create(user)

	return result.Error
}

func (r *userRepository) Update(currentUser *models.User, updatedUser *models.User) error {
	if err := updatedUser.HashPassword(); err != nil {
		return err
	}

	result := r.db.Model(currentUser).Updates(updatedUser)

	return result.Error
}

func (r *userRepository) Save(user *models.User) error {
	result := r.db.Save(user)

	return result.Error
}

func (r *userRepository) Delete(user *models.User) error {
	result := r.db.Delete(user)

	return result.Error
}
