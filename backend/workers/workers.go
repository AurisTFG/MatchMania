package workers

import (
	"MatchManiaAPI/repositories"
)

type Workers struct {
	MatchmakingWorker MatchmakingWorker
}

func NewWorkers(
	repos *repositories.Repositories,
) *Workers {
	return &Workers{
		MatchmakingWorker: NewMatchmakingWorker(repos.QueueRepository, repos.MatchRepository, repos.TeamRepository),
	}
}

func (w *Workers) Start() {
	w.MatchmakingWorker.Start()
}
