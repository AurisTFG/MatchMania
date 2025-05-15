// nolint
package services_test

import (
	"MatchManiaAPI/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateElo_ZeroElo(t *testing.T) {
	s := services.NewEloService()

	eloA, eloB := s.CalculateElo(0, 0, 1, 0)

	assert.Equal(t, uint(0), eloA)
	assert.Equal(t, uint(0), eloB)
}

func TestCalculateElo_PlayerAWins(t *testing.T) {
	s := services.NewEloService()

	eloAStart := uint(1500)
	eloBStart := uint(1500)
	scoreA := uint(3)
	scoreB := uint(1)

	eloAEnd, eloBEnd := s.CalculateElo(eloAStart, eloBStart, scoreA, scoreB)

	assert.Greater(t, eloAEnd, eloAStart)
	assert.Less(t, eloBEnd, eloBStart)
}

func TestCalculateElo_PlayerBWins(t *testing.T) {
	s := services.NewEloService()

	eloAStart := uint(1600)
	eloBStart := uint(1400)
	scoreA := uint(0)
	scoreB := uint(2)

	eloAEnd, eloBEnd := s.CalculateElo(eloAStart, eloBStart, scoreA, scoreB)

	assert.Less(t, eloAEnd, eloAStart)
	assert.Greater(t, eloBEnd, eloBStart)
}

func TestCalculateElo_Draw(t *testing.T) {
	s := services.NewEloService()
	
	eloAStart := uint(1800)
	eloBStart := uint(1700)
	scoreA := uint(2)
	scoreB := uint(2)

	eloAEnd, eloBEnd := s.CalculateElo(eloAStart, eloBStart, scoreA, scoreB)

	assert.Less(t, eloAEnd, eloAStart)
	assert.Greater(t, eloBEnd, eloBStart)
}
