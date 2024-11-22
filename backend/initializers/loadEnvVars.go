package initializers

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	ClientURL  string `mapstructure:"CLIENT_URL"`

	DBHost         string `mapstructure:"DB_HOST"`
	DBUser         string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBSSLMode      string `mapstructure:"DB_SSLMODE"`

	JWTSecret                     string `mapstructure:"JWT"`
	JWTIssuer                     string `mapstructure:"JWT_ISSUER"`
	JWTAudience                   string `mapstructure:"JWT_AUDIENCE"`
	JWTTokenExpirationDays        int    `mapstructure:"JWT_TOKEN_EXPIRATION_DAYS"`
	JWTRefreshTokenExpirationDays int    `mapstructure:"JWT_REFRESH_EXPIRATION_DAYS"`
}

func LoadEnvVars() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return nil
}
