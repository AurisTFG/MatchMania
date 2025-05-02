package repositories

import "MatchManiaAPI/config"

type Repositories struct {
	SessionRepository              SessionRepository
	UserRepository                 UserRepository
	LeagueRepository               LeagueRepository
	TeamRepository                 TeamRepository
	ResultRepository               ResultRepository
	TrackmaniaOAuthStateRepository TrackmaniaOAuthStateRepository
	TrackmaniaTrackRepository      TrackmaniaTrackRepository
}

func SetupRepositories(db *config.DB) *Repositories {
	return &Repositories{
		SessionRepository:              NewSessionRepository(db),
		UserRepository:                 NewUserRepository(db),
		LeagueRepository:               NewLeagueRepository(db),
		TeamRepository:                 NewTeamRepository(db),
		ResultRepository:               NewResultRepository(db),
		TrackmaniaOAuthStateRepository: NewTrackmaniaOAuthStateRepository(db),
		TrackmaniaTrackRepository:      NewTrackmaniaTrackRepository(db),
	}
}
