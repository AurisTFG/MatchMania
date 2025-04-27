package services

import (
	"MatchManiaAPI/models"
	requests "MatchManiaAPI/models/dtos/requests/auth"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"

	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(uuid.UUID) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetDistincPermissionsByUserId(uuid.UUID) ([]string, error)
	CreateUser(signUpDto *requests.SignupDto) error
	DeleteUser(*models.User) error
	UpdateUserWithTrackmaniaUser(userId uuid.UUID, trackmaniaUser *responses.TrackmaniaOAuthUserDto) error
	UpdateUserWithTrackmaniaTracks(userId uuid.UUID, tracksDto []responses.TrackmaniaOAuthTracksDto) error
}

type userService struct {
	repo      repositories.UserRepository
	trackRepo repositories.TrackmaniaOAuthTrackRepository
}

func NewUserService(
	repo repositories.UserRepository,
	trackRepo repositories.TrackmaniaOAuthTrackRepository,
) UserService {
	return &userService{repo: repo, trackRepo: trackRepo}
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

func (s *userService) GetDistincPermissionsByUserId(userId uuid.UUID) ([]string, error) {
	return s.repo.GetDistincPermissionsByUserId(userId)
}

func (s *userService) CreateUser(signUpDto *requests.SignupDto) error {
	newUser := utils.MustCopy[models.User](signUpDto)

	return s.repo.Create(newUser)
}

func (s *userService) DeleteUser(user *models.User) error {
	return s.repo.Delete(user)
}

func (s *userService) UpdateUserWithTrackmaniaUser(
	userId uuid.UUID,
	trackmaniaUser *responses.TrackmaniaOAuthUserDto,
) error {
	user, err := s.repo.FindById(userId)
	if err != nil {
		return err
	}

	user.TrackmaniaId = trackmaniaUser.AccountId
	user.TrackmaniaName = trackmaniaUser.DisplayName

	if err = s.repo.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUserWithTrackmaniaTracks(
	userId uuid.UUID,
	tracksDto []responses.TrackmaniaOAuthTracksDto,
) error {
	tracks := utils.MustCopy[[]models.TrackmaniaOauthTrack](tracksDto)

	if err := s.trackRepo.DeleteAllTracksByUserId(userId); err != nil {
		return err
	}

	if err := s.trackRepo.InsertAllTracksForUser(userId, *tracks); err != nil {
		return err
	}

	return nil
}
