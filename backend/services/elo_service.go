package services

import (
	"MatchManiaAPI/constants"
	"math"
)

type EloService interface {
	CalculateElo(eloA, eloB, scoreA, scoreB uint) (uint, uint)
}

type eloService struct{}

func NewEloService() EloService {
	return &eloService{}
}

func (s *eloService) CalculateElo(eloA, eloB, scoreA, scoreB uint) (uint, uint) {
	if eloA == 0 && eloB == 0 {
		return eloA, eloB
	}

	expectedScoreA := expectedScore(eloA, eloB)
	expectedScoreB := expectedScore(eloB, eloA)

	actualScoreA := actualScore(scoreA, scoreB)
	actualScoreB := actualScore(scoreB, scoreA)

	newEloA := newElo(eloA, expectedScoreA, actualScoreA)
	newEloB := newElo(eloB, expectedScoreB, actualScoreB)

	return newEloA, newEloB
}

func expectedScore(elo, opponentElo uint) float64 {
	return 1.0 / (1.0 + math.Pow(10, (float64(opponentElo)-float64(elo))/400.0))
}

func actualScore(score, opponentScore uint) float64 {
	if score > opponentScore {
		return 1.0
	} else if score < opponentScore {
		return 0.0
	}

	return 0.5
}

func newElo(elo uint, expectedScore, actualScore float64) uint {
	return uint(float64(elo) + float64(constants.KFactor)*(actualScore-expectedScore))
}
