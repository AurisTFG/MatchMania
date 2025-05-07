package workers

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"
	"fmt"
	"time"
)

type MatchmakingWorker interface {
	Start()
}

type matchmakingWorker struct {
	queueRepository repositories.QueueRepository
	matchRepository repositories.MatchRepository
	teamRepository  repositories.TeamRepository
}

func NewMatchmakingWorker(
	queueRepository repositories.QueueRepository,
	matchRepository repositories.MatchRepository,
	teamRepository repositories.TeamRepository,
) MatchmakingWorker {
	return &matchmakingWorker{
		queueRepository: queueRepository,
		matchRepository: matchRepository,
		teamRepository:  teamRepository,
	}
}

func (w *matchmakingWorker) Start() {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for range ticker.C {
			if err := w.tick(); err != nil {
				fmt.Printf("Error in matchmaking worker: %v\n", err)
			}

			fmt.Printf("matchmaking worker tick at %s\n", time.Now().Format(time.DateTime))
		}
	}()
}

func (w *matchmakingWorker) tick() error {
	queues, err := w.queueRepository.GetAll()
	if err != nil {
		return err
	}

	for _, queue := range queues {
		if len(queue.Teams) < 2 {
			continue
		}

		for len(queue.Teams) >= 2 {
			teamA := queue.Teams[0]
			teamB := queue.Teams[1]

			teamA.QueueId = nil
			teamB.QueueId = nil

			if err = w.teamRepository.Save(&teamA); err != nil {
				return err
			}

			if err = w.teamRepository.Save(&teamB); err != nil {
				return err
			}

			match := &models.Match{
				LeagueId: queue.League.Id,
				GameMode: queue.GameMode,
				Teams:    []models.Team{teamA, teamB},
			}

			if err = w.matchRepository.Create(match); err != nil {
				return err
			}

			fmt.Printf("Match created: %s vs %s\n", teamA.Id, teamB.Id)
			queue.Teams = queue.Teams[2:]
		}
	}

	matches, err := w.matchRepository.GetAll()
	if err != nil {
		return err
	}

	for _, match := range matches {
		if len(match.Teams) == 0 {
			err = w.matchRepository.Delete(match.Id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
