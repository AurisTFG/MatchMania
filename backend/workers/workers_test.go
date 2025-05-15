// nolint
package workers_test

import (
	"testing"

	"MatchManiaAPI/repositories"
	"MatchManiaAPI/services"
	"MatchManiaAPI/workers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock MatchmakingWorker ---

type MockMatchmakingWorker struct {
	mock.Mock
}

func (m *MockMatchmakingWorker) Start() {
	m.Called()
}

// --- Replace NewMatchmakingWorker temporarily ---

var originalNewMatchmakingWorker = workers.NewMatchmakingWorker

func resetNewMatchmakingWorker() {
	workers.NewMatchmakingWorker = originalNewMatchmakingWorker
}

// --- Tests ---

func TestNewWorkers(t *testing.T) {
	defer resetNewMatchmakingWorker()

	mockWorker := &MockMatchmakingWorker{}
	workers.NewMatchmakingWorker = func(
		queueRepo repositories.QueueRepository,
		matchRepo repositories.MatchRepository,
		teamRepo repositories.TeamRepository,
		tmApiService services.TrackmaniaApiService,
	) workers.MatchmakingWorker {
		return mockWorker
	}

	repos := &repositories.Repositories{}
	services := &services.Services{}

	w := workers.NewWorkers(repos, services)

	assert.NotNil(t, w)
	assert.Equal(t, mockWorker, w.MatchmakingWorker)
}

func TestWorkers_Start(t *testing.T) {
	defer resetNewMatchmakingWorker()

	mockWorker := &MockMatchmakingWorker{}
	mockWorker.On("Start").Return()

	workers.NewMatchmakingWorker = func(
		queueRepo repositories.QueueRepository,
		matchRepo repositories.MatchRepository,
		teamRepo repositories.TeamRepository,
		tmApiService services.TrackmaniaApiService,
	) workers.MatchmakingWorker {
		return mockWorker
	}

	repos := &repositories.Repositories{}
	services := &services.Services{}

	w := workers.NewWorkers(repos, services)
	w.Start()

	mockWorker.AssertCalled(t, "Start")
}
