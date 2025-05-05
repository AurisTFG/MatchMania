package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type TrackmaniaTrackRepository interface {
	GetAllTracks() ([]models.TrackmaniaTrack, error)
	InsertAllTracksForUser(userId uuid.UUID, tracks []models.TrackmaniaTrack) error
	DeleteAllTracksByUserId(userId uuid.UUID) error
}
type trackmaniaOAuthTrackRepository struct {
	db *config.DB
}

func NewTrackmaniaTrackRepository(db *config.DB) TrackmaniaTrackRepository {
	return &trackmaniaOAuthTrackRepository{db: db}
}

func (r *trackmaniaOAuthTrackRepository) GetAllTracks() ([]models.TrackmaniaTrack, error) {
	var tracks []models.TrackmaniaTrack

	result := r.db.Find(&tracks)

	return tracks, result.Error
}

func (r *trackmaniaOAuthTrackRepository) InsertAllTracksForUser(
	userId uuid.UUID,
	tracks []models.TrackmaniaTrack,
) error {
	for i := range tracks {
		tracks[i].UserId = userId
	}

	result := r.db.Create(&tracks)

	return result.Error
}

func (r *trackmaniaOAuthTrackRepository) DeleteAllTracksByUserId(userId uuid.UUID) error {
	result := r.db.
		Where("user_id = ?", userId).
		Delete(&models.TrackmaniaTrack{})

	return result.Error
}
