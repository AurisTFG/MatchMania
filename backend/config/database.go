package config

import (
	"MatchManiaAPI/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		&models.Session{},
		&models.Season{},
		&models.Team{},
		&models.Result{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

func SeedDatabase(db *DB, env *Env) error {
	if !env.IsDev {
		return nil
	}

	if err := db.Migrator().DropTable(
		&models.User{},
		&models.Session{},
		&models.Season{},
		&models.Team{},
		&models.Result{},
	); err != nil {
		return err
	}

	if err := MigrateDatabase(db); err != nil {
		return err
	}

	users := []models.User{
		{UUID: uuid.New(), Username: "AdminXD", Email: env.AdminEmail, Password: env.AdminPassword, Role: models.AdminRole},
		{UUID: uuid.New(), Username: "ModeratorXDD", Email: env.ModeratorEmail, Password: env.ModeratorPassword, Role: models.ModeratorRole},
		{UUID: uuid.New(), Username: "UserXDDD", Email: env.UserEmail, Password: env.UserPassword, Role: models.UserRole},
	}

	seasons := []models.Season{
		{UserUUID: users[0].UUID, Name: "TO BE DELETED", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 30)},
		{UserUUID: users[0].UUID, Name: "Fall 2024", StartDate: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{UserUUID: users[0].UUID, Name: "Winter 2025", StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)},
		{UserUUID: users[0].UUID, Name: "Spring 2025", StartDate: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 6, 30, 0, 0, 0, 0, time.UTC)},
		{UserUUID: users[0].UUID, Name: "Summer 2025", StartDate: time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)},
		{UserUUID: users[0].UUID, Name: "Fall 2025", StartDate: time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)},
		{UserUUID: users[0].UUID, Name: "Winter 2026", StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC)},
		{UserUUID: users[0].UUID, Name: "Spring 2026", StartDate: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)},
	}

	teams := []models.Team{
		{UserUUID: users[0].UUID, Name: "TO BE DELETED", SeasonID: 1, Elo: 1234},
		{UserUUID: users[0].UUID, Name: "BIG CLAN", SeasonID: 2, Elo: 1123},
		{UserUUID: users[0].UUID, Name: "Astralis", SeasonID: 2, Elo: 1245},
		{UserUUID: users[0].UUID, Name: "Natus Vincere", SeasonID: 2, Elo: 182},
		{UserUUID: users[0].UUID, Name: "G2 Esports", SeasonID: 2, Elo: 945},
		{UserUUID: users[0].UUID, Name: "Team Liquid", SeasonID: 2, Elo: 885},
		{UserUUID: users[0].UUID, Name: "FaZe Clan", SeasonID: 2, Elo: 812},
		{UserUUID: users[0].UUID, Name: "Fnatic", SeasonID: 2, Elo: 789},
	}

	results := []models.Result{
		{UserUUID: users[0].UUID, MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: "19", OpponentScore: "9", TeamID: 3, OpponentTeamID: 2, SeasonID: 2},
		{UserUUID: users[0].UUID, MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: "15", OpponentScore: "5", TeamID: 4, OpponentTeamID: 3, SeasonID: 2},
		{UserUUID: users[0].UUID, MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: "16", OpponentScore: "13", TeamID: 5, OpponentTeamID: 4, SeasonID: 2},
		{UserUUID: users[0].UUID, MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: "8", OpponentScore: "6", TeamID: 6, OpponentTeamID: 5, SeasonID: 2},
		{UserUUID: users[0].UUID, MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: "11", OpponentScore: "2", TeamID: 7, OpponentTeamID: 6, SeasonID: 2},
		{UserUUID: users[0].UUID, MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: "7", OpponentScore: "8", TeamID: 8, OpponentTeamID: 7, SeasonID: 2},
		{UserUUID: users[0].UUID, MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: "12", OpponentScore: "15", TeamID: 2, OpponentTeamID: 8, SeasonID: 2},
	}

	for _, user := range users {
		if err := user.HashPassword(); err != nil {
			return err
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}
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

	return nil
}
