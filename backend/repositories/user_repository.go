package repositories

import (
	"MatchManiaAPI/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(uint) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	Create(*models.User) (*models.User, error)
	Update(*models.User) (*models.User, error)
	Delete(*models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User

	result := r.db.Find(&users)

	return users, result.Error
}

func (r *userRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User

	result := r.db.First(&user, userID)

	return &user, result.Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	result := r.db.Where("email = ?", email).First(&user)

	return &user, result.Error
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	result := r.db.Create(user)

	return user, result.Error
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	result := r.db.Save(user)

	return user, result.Error
}

func (r *userRepository) Delete(user *models.User) error {
	result := r.db.Delete(user)

	return result.Error
}
