package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/auth"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/models/enums"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"

	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(uuid.UUID) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetDistinctPermissionsByUserId(uuid.UUID) ([]string, error)
	CreateUser(signUpDto *requests.SignupDto) error
	DeleteUser(*models.User) error
	UpdateUserWithTrackmaniaUser(userId uuid.UUID, trackmaniaUser *responses.TrackmaniaOAuthUserDto) error
	UpdateUserWithTrackmaniaTracks(userId uuid.UUID, tracksDto []responses.TrackmaniaOAuthFavoritesDto) error
}

type userService struct {
	userRepository  repositories.UserRepository
	roleRepository  repositories.RoleRepository
	trackRepository repositories.TrackmaniaTrackRepository
}

func NewUserService(
	userRepository repositories.UserRepository,
	roleRepository repositories.RoleRepository,
	trackRepository repositories.TrackmaniaTrackRepository,
) UserService {
	return &userService{
		userRepository:  userRepository,
		roleRepository:  roleRepository,
		trackRepository: trackRepository,
	}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) GetUserById(userId uuid.UUID) (*models.User, error) {
	return s.userRepository.FindById(userId)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.FindByEmail(email)
}

func (s *userService) GetDistinctPermissionsByUserId(userId uuid.UUID) ([]string, error) {
	return s.userRepository.GetDistinctPermissionsByUserId(userId)
}

func (s *userService) CreateUser(signUpDto *requests.SignupDto) error {
	newUser := utils.MustCopy[models.User](signUpDto)

	return s.userRepository.Create(newUser)
}

func (s *userService) DeleteUser(user *models.User) error {
	return s.userRepository.Delete(user)
}

func (s *userService) UpdateUserWithTrackmaniaUser(
	userId uuid.UUID,
	trackmaniaUser *responses.TrackmaniaOAuthUserDto,
) error {
	user, err := s.userRepository.FindById(userId)
	if err != nil {
		return err
	}

	trackmaniaPlayerRole, err := s.roleRepository.GetByName(string(enums.TrackmaniaPlayerRole))
	if err != nil {
		return err
	}

	user.TrackmaniaId = trackmaniaUser.AccountId
	user.TrackmaniaName = trackmaniaUser.DisplayName
	user.Roles = append(user.Roles, *trackmaniaPlayerRole)

	if err = s.userRepository.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUserWithTrackmaniaTracks(
	userId uuid.UUID,
	tracksDto []responses.TrackmaniaOAuthFavoritesDto,
) error {
	tracks := utils.MustCopy[[]models.TrackmaniaTrack](tracksDto)

	if err := s.trackRepository.DeleteAllTracksByUserId(userId); err != nil {
		return err
	}

	if err := s.trackRepository.InsertAllTracksForUser(userId, *tracks); err != nil {
		return err
	}

	return nil
}
