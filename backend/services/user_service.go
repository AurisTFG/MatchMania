package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(uint) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	CreateUser(signUpDto *models.SignUpDto) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(userID uint) (*models.User, error) {
	return s.repo.FindByID(userID)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) CreateUser(signUpDto *models.SignUpDto) (*models.User, error) {
	newUser := signUpDto.ToUser()

	return s.repo.Create(&newUser)
}
