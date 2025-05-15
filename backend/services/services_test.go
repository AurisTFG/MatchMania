// nolint
package services_test

import (
	"testing"

	"MatchManiaAPI/config"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/services"

	"github.com/stretchr/testify/assert"
)

func TestNewServices(t *testing.T) {
	env := &config.Env{
		IsDev:  true,
		IsProd: false,
	}

	repos := &repositories.Repositories{
		AppSettingRepository:           nil,
		UserRepository:                 nil,
		RoleRepository:                 nil,
		TrackmaniaTrackRepository:      nil,
		PlayerRepository:               nil,
		SessionRepository:              nil,
		LeagueRepository:               nil,
		TeamRepository:                 nil,
		ResultRepository:               nil,
		MatchRepository:                nil,
		QueueRepository:                nil,
		TrackmaniaOAuthStateRepository: nil,
	}

	s := services.NewServices(env, repos)

	assert.NotNil(t, s)
	assert.NotNil(t, s.AppSettingService)
	assert.NotNil(t, s.UserService)
	assert.NotNil(t, s.PlayerService)
	assert.NotNil(t, s.AuthService)
	assert.NotNil(t, s.LeagueService)
	assert.NotNil(t, s.TeamService)
	assert.NotNil(t, s.ResultService)
	assert.NotNil(t, s.EloService)
	assert.NotNil(t, s.QueueService)
	assert.NotNil(t, s.MatchService)
	assert.NotNil(t, s.UbisoftApiService)
	assert.NotNil(t, s.NadeoApiService)
	assert.NotNil(t, s.TrackmaniaApiService)
	assert.NotNil(t, s.TrackmaniaOAuthService)
}
