// nolint
package controllers_test

import (
	"MatchManiaAPI/controllers"
	"MatchManiaAPI/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewControllers(t *testing.T) {
	s := &services.Services{
		AuthService:            nil,
		UserService:            nil,
		PlayerService:          nil,
		LeagueService:          nil,
		TeamService:            nil,
		ResultService:          nil,
		QueueService:           nil,
		MatchService:           nil,
		TrackmaniaOAuthService: nil,
	}

	ctrl := controllers.NewControllers(s)

	assert.NotNil(t, ctrl)
	assert.NotNil(t, ctrl.AuthController)
	assert.NotNil(t, ctrl.UserController)
	assert.NotNil(t, ctrl.PlayerController)
	assert.NotNil(t, ctrl.LeagueController)
	assert.NotNil(t, ctrl.TeamController)
	assert.NotNil(t, ctrl.ResultController)
	assert.NotNil(t, ctrl.TrackmaniaOAuthController)
	assert.NotNil(t, ctrl.QueueController)
	assert.NotNil(t, ctrl.MatchController)
}
