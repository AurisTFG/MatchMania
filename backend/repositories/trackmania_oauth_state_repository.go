package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type TrackmaniaOAuthStateRepository interface {
	SaveState(state *models.TrackmaniaOauthState) error
	DoesStateExist(state string) bool
	GetUserIdByState(state string) (uuid.UUID, error)
	DeleteStateByUserId(userId uuid.UUID) error
}

type trackmaniaOAuthStateRepository struct {
	db *config.DB
}

func NewTrackmaniaOAuthStateRepository(db *config.DB) TrackmaniaOAuthStateRepository {
	return &trackmaniaOAuthStateRepository{db: db}
}

func (r *trackmaniaOAuthStateRepository) SaveState(state *models.TrackmaniaOauthState) error {
	return r.db.Create(state).Error
}

func (r *trackmaniaOAuthStateRepository) DoesStateExist(state string) bool {
	var oauthState models.TrackmaniaOauthState

	err := r.db.
		Where("state = ?", state).
		First(&oauthState).
		Error

	return err == nil
}

func (r *trackmaniaOAuthStateRepository) GetUserIdByState(state string) (uuid.UUID, error) {
	var oauthState models.TrackmaniaOauthState

	err := r.db.
		Where("state = ?", state).
		First(&oauthState).
		Error

	if err != nil {
		return uuid.Nil, err
	}

	return oauthState.UserId, nil
}

func (r *trackmaniaOAuthStateRepository) DeleteStateByUserId(userId uuid.UUID) error {
	var oauthState models.TrackmaniaOauthState

	err := r.db.
		Where("user_id = ?", userId).
		Delete(&oauthState).
		Error

	if err != nil {
		return err
	}

	return nil
}
