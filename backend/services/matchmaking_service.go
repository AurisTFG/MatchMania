package services

import (
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Matchup struct {
	TeamA uuid.UUID
	TeamB uuid.UUID
}

type MatchmakingService interface {
	JoinQueue(leagueId uuid.UUID, teamId uuid.UUID) error
	LeaveQueue(leagueId uuid.UUID, teamId uuid.UUID) error
	GetQueueTeamsCount(leagueId uuid.UUID) (uint, error)
	GetOngoingMatches() []Matchup
	IsInMatch(teamId uuid.UUID) bool
	StartMatchmakingWorker()
}

type matchmakingService struct {
	queues         map[uuid.UUID][]uuid.UUID
	ongoingMatches []Matchup
	mu             sync.Mutex
}

func NewMatchmakingService() MatchmakingService {
	return &matchmakingService{
		queues:         make(map[uuid.UUID][]uuid.UUID),
		ongoingMatches: make([]Matchup, 0),
	}
}
func (ms *matchmakingService) JoinQueue(leagueId uuid.UUID, teamId uuid.UUID) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, exists := ms.queues[leagueId]; !exists {
		ms.queues[leagueId] = make([]uuid.UUID, 0)
	}

	if slices.ContainsFunc(ms.queues[leagueId], func(id uuid.UUID) bool { return id == teamId }) {
		return errors.New("team already in queue for this league")
	}

	ms.queues[leagueId] = append(ms.queues[leagueId], teamId)
	return nil
}

func (ms *matchmakingService) LeaveQueue(leagueId uuid.UUID, teamId uuid.UUID) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	queue, exists := ms.queues[leagueId]
	if !exists {
		return errors.New("no queue found for this league")
	}

	for i, id := range queue {
		if id == teamId {
			ms.queues[leagueId] = slices.Delete(queue, i, i+1)
			return nil
		}
	}

	return errors.New("team not found in the queue for this league")
}

func (ms *matchmakingService) GetQueueTeamsCount(leagueId uuid.UUID) (uint, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	queue, exists := ms.queues[leagueId]
	if !exists {
		return 0, errors.New("no queue found for this league")
	}

	return uint(len(queue)), nil
}

func (ms *matchmakingService) GetOngoingMatches() []Matchup {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	matches := make([]Matchup, len(ms.ongoingMatches))
	copy(matches, ms.ongoingMatches)
	return matches
}

func (ms *matchmakingService) IsInMatch(teamId uuid.UUID) bool {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	for _, match := range ms.ongoingMatches {
		if match.TeamA == teamId || match.TeamB == teamId {
			return true
		}
	}

	return false
}

func (ms *matchmakingService) StartMatchmakingWorker() {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			ms.processQueues()
		}
	}()
}

func (ms *matchmakingService) processQueues() {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if len(ms.queues) == 0 {
		fmt.Println("No queues to process")
		return
	}

	for leagueId, queue := range ms.queues {
		for len(queue) >= 2 {
			matchedTeams := queue[:2]
			ms.queues[leagueId] = queue[2:]

			matchup := Matchup{
				TeamA: matchedTeams[0],
				TeamB: matchedTeams[1],
			}
			ms.ongoingMatches = append(ms.ongoingMatches, matchup)

			fmt.Printf("Match created: %s vs %s in league %s\n", matchup.TeamA, matchup.TeamB, leagueId)
		}

		fmt.Printf("Remaining teams in queue for league %s: %v\n", leagueId, queue)
	}
}
