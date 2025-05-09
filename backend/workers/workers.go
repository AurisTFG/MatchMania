package workers

import (
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/services"
)

type Workers struct {
	MatchmakingWorker MatchmakingWorker
}

func NewWorkers(
	repos *repositories.Repositories,
	services *services.Services,
) *Workers {
	return &Workers{
		MatchmakingWorker: NewMatchmakingWorker(
			repos.QueueRepository,
			repos.MatchRepository,
			repos.TeamRepository,
			services.TrackmaniaApiService,
		),
	}
}

func (w *Workers) Start() {
	w.MatchmakingWorker.Start()
}
