package repositories

import "MatchManiaAPI/config"

type Repositories struct {
	SessionRepository              SessionRepository
	UserRepository                 UserRepository
	RoleRepository                 RoleRepository
	PlayerRepository               PlayerRepository
	LeagueRepository               LeagueRepository
	TeamRepository                 TeamRepository
	ResultRepository               ResultRepository
	TrackmaniaOAuthStateRepository TrackmaniaOAuthStateRepository
	TrackmaniaTrackRepository      TrackmaniaTrackRepository
	QueueRepository                QueueRepository
	MatchRepository                MatchRepository
}

func NewRepositories(
	db *config.DB,
) *Repositories {
	return &Repositories{
		SessionRepository:              NewSessionRepository(db),
		UserRepository:                 NewUserRepository(db),
		RoleRepository:                 NewRoleRepository(db),
		PlayerRepository:               NewPlayerRepository(db),
		LeagueRepository:               NewLeagueRepository(db),
		TeamRepository:                 NewTeamRepository(db),
		ResultRepository:               NewResultRepository(db),
		TrackmaniaOAuthStateRepository: NewTrackmaniaOAuthStateRepository(db),
		TrackmaniaTrackRepository:      NewTrackmaniaTrackRepository(db),
		QueueRepository:                NewQueueRepository(db),
		MatchRepository:                NewMatchRepository(db),
	}
}
