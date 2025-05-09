package workers

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/services"
	"fmt"
	"time"
)

type MatchmakingWorker interface {
	Start()
}

type matchmakingWorker struct {
	queueRepository      repositories.QueueRepository
	matchRepository      repositories.MatchRepository
	teamRepository       repositories.TeamRepository
	trackmaniaApiService services.TrackmaniaApiService
}

func NewMatchmakingWorker(
	queueRepository repositories.QueueRepository,
	matchRepository repositories.MatchRepository,
	teamRepository repositories.TeamRepository,
	trackmaniaApiService services.TrackmaniaApiService,
) MatchmakingWorker {
	return &matchmakingWorker{
		queueRepository:      queueRepository,
		matchRepository:      matchRepository,
		teamRepository:       teamRepository,
		trackmaniaApiService: trackmaniaApiService,
	}
}

func (w *matchmakingWorker) Start() {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for range ticker.C {
			if err := w.tick(); err != nil {
				fmt.Printf("Error in matchmaking worker: %v\n", err)
			}
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

			createResponseDto, err := w.trackmaniaApiService.CreateCompetition(
				getCompetitionLabel(&queue.League, &teamA, &teamB),
				getTrackUids(&queue.League),
			)
			if err != nil {
				return err
			}

			err = w.trackmaniaApiService.AddTeamsToCompetition(&teamA, &teamB, createResponseDto.Competition.Id)
			if err != nil {
				return err
			}

			match := &models.Match{
				GameMode:                queue.GameMode,
				LeagueId:                queue.League.Id,
				League:                  queue.League,
				TrackmaniaCompetitionId: createResponseDto.Competition.Id,
			}

			if err = w.matchRepository.Create(match); err != nil {
				return err
			}

			teamA.QueueId = nil
			teamB.QueueId = nil
			teamA.MatchId = &match.Id
			teamB.MatchId = &match.Id

			if err = w.teamRepository.Save(&teamA); err != nil {
				return err
			}

			if err = w.teamRepository.Save(&teamB); err != nil {
				return err
			}

			fmt.Printf("Match created: %s\n", match.Id)
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

func getCompetitionLabel(league *models.League, teamA *models.Team, teamB *models.Team) string {
	return fmt.Sprintf("%s - %s vs %s", league.Name, teamA.Name, teamB.Name)
}

func getTrackUids(league *models.League) []string {
	var trackUids []string

	for _, track := range league.Tracks {
		trackUids = append(trackUids, track.Uid)
	}

	return trackUids
}
