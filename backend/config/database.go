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

	err := db.AutoMigrate(
		&models.AppSetting{},
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Session{},
		&models.League{},
		&models.Team{},
		&models.Result{},
		&models.TrackmaniaTrack{},
		&models.TrackmaniaOauthState{},
		&models.Match{},
		&models.Queue{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
