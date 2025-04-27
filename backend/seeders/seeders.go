package seeders

import (
	"MatchManiaAPI/config"
)

func SeedDatabase(db *config.DB, env *config.Env) error {
	if err := SeedPermissions(db, env); err != nil {
		return err
	}

	if err := SeedRoles(db, env); err != nil {
		return err
	}

	if env.IsProd {
		return nil
	}

	if err := SeedUsers(db, env); err != nil {
		return err
	}

	if err := SeedSeasons(db, env); err != nil {
		return err
	}

	if err := SeedTeams(db, env); err != nil {
		return err
	}

	if err := SeedResults(db, env); err != nil {
		return err
	}

	return nil
}
