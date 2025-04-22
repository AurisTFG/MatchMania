package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type TrackmaniaOAuthTrackRepository interface {
	GetAllTracks() ([]models.TrackmaniaOauthTrack, error)
	InsertAllTracksForUser(userId uuid.UUID, tracks []models.TrackmaniaOauthTrack) error
	DeleteAllTracksByUserId(userId uuid.UUID) error
}
type trackmaniaOAuthTrackRepository struct {
	db *config.DB
}

func NewTrackmaniaOAuthTrackRepository(db *config.DB) TrackmaniaOAuthTrackRepository {
	return &trackmaniaOAuthTrackRepository{db: db}
}

func (r *trackmaniaOAuthTrackRepository) GetAllTracks() ([]models.TrackmaniaOauthTrack, error) {
	var tracks []models.TrackmaniaOauthTrack

	result := r.db.Find(&tracks)

	return tracks, result.Error
}

func (r *trackmaniaOAuthTrackRepository) InsertAllTracksForUser(
	userId uuid.UUID,
	tracks []models.TrackmaniaOauthTrack,
) error {
	for i := range tracks {
		tracks[i].UserId = userId
	}

	result := r.db.Create(&tracks)

	return result.Error
}

func (r *trackmaniaOAuthTrackRepository) DeleteAllTracksByUserId(userId uuid.UUID) error {
	result := r.db.Where("user_id = ?", userId).Delete(&models.TrackmaniaOauthTrack{})

	return result.Error
}
