package config

import (
	"MatchManiaAPI/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func ConnectDatabase(env *Env) (*DB, error) {
	db, err := gorm.Open(postgres.Open(env.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database session: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection pool: %w", err)
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{db}, nil
}

func MigrateDatabase(db *DB) error {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	if !db.Migrator().HasTable(&models.User{}) {
		db.Exec(`DROP TYPE IF EXISTS role;`)
		db.Exec(`CREATE TYPE role AS ENUM ('admin', 'moderator', 'user');`)
	}

	err := db.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.Season{},
		&models.Team{},
		&models.Result{},
		&models.TrackmaniaOauthState{},
		&models.TrackmaniaOauthTrack{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
