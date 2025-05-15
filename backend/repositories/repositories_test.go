// nolint
package repositories_test

import (
	"testing"

	"MatchManiaAPI/config"
	"MatchManiaAPI/repositories"

	"github.com/stretchr/testify/assert"
)

func TestNewRepositories(t *testing.T) {
	db := &config.DB{}

	r := repositories.NewRepositories(db)

	assert.NotNil(t, r)
	assert.NotNil(t, r.AppSettingRepository)
	assert.NotNil(t, r.SessionRepository)
	assert.NotNil(t, r.UserRepository)
	assert.NotNil(t, r.RoleRepository)
	assert.NotNil(t, r.PlayerRepository)
	assert.NotNil(t, r.LeagueRepository)
	assert.NotNil(t, r.TeamRepository)
	assert.NotNil(t, r.ResultRepository)
	assert.NotNil(t, r.TrackmaniaOAuthStateRepository)
	assert.NotNil(t, r.TrackmaniaTrackRepository)
	assert.NotNil(t, r.QueueRepository)
	assert.NotNil(t, r.MatchRepository)
}
