package repositories

import "MatchManiaAPI/config"

type Repositories struct {
	SessionRepository SessionRepository
	UserRepository    UserRepository
	SeasonRepository  SeasonRepository
	TeamRepository    TeamRepository
	ResultRepository  ResultRepository
}

func SetupRepositories(db *config.DB) *Repositories {
	return &Repositories{
		SessionRepository: NewSessionRepository(db),
		UserRepository:    NewUserRepository(db),
		SeasonRepository:  NewSeasonRepository(db),
		TeamRepository:    NewTeamRepository(db),
		ResultRepository:  NewResultRepository(db),
	}
}
