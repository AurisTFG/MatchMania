package config

import (
	"MatchManiaAPI/models"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func ConnectDatabase(env *Env) (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		env.DBHost, env.DBUser, env.DBUserPassword, env.DBName, env.DBPort, env.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database session: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection pool: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{db}, nil
}

func MigrateDatabase(db *DB) error {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Exec(`DROP TYPE IF EXISTS role;`)
	db.Exec(`CREATE TYPE role AS ENUM ('admin', 'moderator', 'user');`)

	err := db.AutoMigrate(
		&models.User{},
		&models.Season{},
		&models.Team{},
		&models.Result{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

func SeedDatabase(db *DB) error {
	originalLogger := db.Logger
	db.Logger = originalLogger.LogMode(logger.Silent)

	for _, table := range []string{"results", "teams", "seasons", "users"} {
		if err := db.Exec("DELETE FROM " + table).Error; err != nil {
			return fmt.Errorf("failed to delete rows from table %s: %w", table, err)
		}
	}

	for _, table := range []string{"results", "teams", "seasons"} {
		seqName := fmt.Sprintf("%s_id_seq", table)
		if err := db.Exec("ALTER SEQUENCE " + seqName + " RESTART WITH 1").Error; err != nil {
			return fmt.Errorf("failed to reset id sequence for table %s: %w", table, err)
		}
	}

	seasons := []models.Season{
		{Name: "TO BE DELETED", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 30)},
		{Name: "Fall 2024", StartDate: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Winter 2025", StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Spring 2025", StartDate: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 6, 30, 0, 0, 0, 0, time.UTC)},
		{Name: "Summer 2025", StartDate: time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Fall 2025", StartDate: time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Winter 2026", StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Spring 2026", StartDate: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)},
	}

	teams := []models.Team{
		{Name: "TO BE DELETED", SeasonID: 1, Elo: 1234},
		{Name: "BIG CLAN", SeasonID: 2, Elo: 1123},
		{Name: "Astralis", SeasonID: 2, Elo: 1245},
		{Name: "Natus Vincere", SeasonID: 2, Elo: 182},
		{Name: "G2 Esports", SeasonID: 2, Elo: 945},
		{Name: "Team Liquid", SeasonID: 2, Elo: 885},
		{Name: "FaZe Clan", SeasonID: 2, Elo: 812},
		{Name: "Fnatic", SeasonID: 2, Elo: 789},
	}

	results := []models.Result{
		{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 19, OpponentScore: 9, TeamID: 3, OpponentTeamID: 2, SeasonID: 2},
		{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 15, OpponentScore: 5, TeamID: 4, OpponentTeamID: 3, SeasonID: 2},
		{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 16, OpponentScore: 13, TeamID: 5, OpponentTeamID: 4, SeasonID: 2},
		{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 8, OpponentScore: 6, TeamID: 6, OpponentTeamID: 5, SeasonID: 2},
		{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 11, OpponentScore: 2, TeamID: 7, OpponentTeamID: 6, SeasonID: 2},
		{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 7, OpponentScore: 8, TeamID: 8, OpponentTeamID: 7, SeasonID: 2},
		{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 12, OpponentScore: 15, TeamID: 2, OpponentTeamID: 8, SeasonID: 2},
	}

	users := []models.User{
		{Username: "AdminXD", Email: "adminemail@gmail.com", Password: "AdminPassword", Role: models.AdminRole},
		{Username: "ModeratorXDD", Email: "moderatoremail@gmail.com", Password: "ModeratorPassword", Role: models.ModeratorRole},
		{Username: "UserXDDD", Email: "userremail@gmail.com", Password: "UserPassword", Role: models.UserRole},
	}

	for _, season := range seasons {
		if err := db.Create(&season).Error; err != nil {
			return err
		}
	}

	for _, team := range teams {
		if err := db.Create(&team).Error; err != nil {
			return err
		}
	}

	for _, result := range results {
		if err := db.Create(&result).Error; err != nil {
			return err
		}
	}

	for _, user := range users {
		user.HashPassword()
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	db.Logger = originalLogger

	return nil
}
