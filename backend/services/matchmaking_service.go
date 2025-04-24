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
	JoinQueue(seasonId uuid.UUID, teamId uuid.UUID) error
	LeaveQueue(seasonId uuid.UUID, teamId uuid.UUID) error
	GetQueueTeamsCount(seasonId uuid.UUID) (uint, error)
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
func (ms *matchmakingService) JoinQueue(seasonId uuid.UUID, teamId uuid.UUID) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, exists := ms.queues[seasonId]; !exists {
		ms.queues[seasonId] = make([]uuid.UUID, 0)
	}

	if slices.ContainsFunc(ms.queues[seasonId], func(id uuid.UUID) bool { return id == teamId }) {
		return errors.New("team already in queue for this season")
	}

	ms.queues[seasonId] = append(ms.queues[seasonId], teamId)
	return nil
}

func (ms *matchmakingService) LeaveQueue(seasonId uuid.UUID, teamId uuid.UUID) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	queue, exists := ms.queues[seasonId]
	if !exists {
		return errors.New("no queue found for this season")
	}

	for i, id := range queue {
		if id == teamId {
			ms.queues[seasonId] = slices.Delete(queue, i, i+1)
			return nil
		}
	}

	return errors.New("team not found in the queue for this season")
}

func (ms *matchmakingService) GetQueueTeamsCount(seasonId uuid.UUID) (uint, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	queue, exists := ms.queues[seasonId]
	if !exists {
		return 0, errors.New("no queue found for this season")
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

	for seasonId, queue := range ms.queues {
		for len(queue) >= 2 {
			matchedTeams := queue[:2]
			ms.queues[seasonId] = queue[2:]

			matchup := Matchup{
				TeamA: matchedTeams[0],
				TeamB: matchedTeams[1],
			}
			ms.ongoingMatches = append(ms.ongoingMatches, matchup)

			fmt.Printf("Match created: %s vs %s in season %s\n", matchup.TeamA, matchup.TeamB, seasonId)
		}

		fmt.Printf("Remaining teams in queue for season %s: %v\n", seasonId, queue)
	}
}
