package initializers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() error {
	if Cfg.DBHost == "" || Cfg.DBUser == "" || Cfg.DBUserPassword == "" || Cfg.DBName == "" || Cfg.DBPort == "" || Cfg.DBSSLMode == "" {
		return fmt.Errorf("missing database configuration values")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		Cfg.DBHost, Cfg.DBUser, Cfg.DBUserPassword, Cfg.DBName, Cfg.DBPort, Cfg.DBSSLMode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to initialize database session: %w", err)
	}

	return nil
}
