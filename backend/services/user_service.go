package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/auth"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"

	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(uuid.UUID) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	CreateUser(signUpDto *requests.SignupDto) error
	DeleteUser(*models.User) error
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

func (s *userService) GetUserById(userId uuid.UUID) (*models.User, error) {
	return s.repo.FindById(userId)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) CreateUser(signUpDto *requests.SignupDto) error {
	newUser := utils.CopyOrPanic[models.User](signUpDto)

	return s.repo.Create(newUser)
}

func (s *userService) DeleteUser(user *models.User) error {
	return s.repo.Delete(user)
}
