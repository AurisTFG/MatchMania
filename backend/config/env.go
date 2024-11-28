package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	ClientURL  string `mapstructure:"CLIENT_URL"`

	DBHost         string `mapstructure:"DB_HOST"`
	DBUser         string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBSSLMode      string `mapstructure:"DB_SSLMODE"`

	JWTAccessTokenSecret          string `mapstructure:"JWT_ACCESS_TOKEN_SECRET"`
	JWTRefreshTokenSecret         string `mapstructure:"JWT_REFRESH_TOKEN_SECRET"`
	JWTIssuer                     string `mapstructure:"JWT_ISSUER"`
	JWTAudience                   string `mapstructure:"JWT_AUDIENCE"`
	JWTTokenExpirationDays        int    `mapstructure:"JWT_TOKEN_EXPIRATION_DAYS"`
	JWTRefreshTokenExpirationDays int    `mapstructure:"JWT_REFRESH_EXPIRATION_DAYS"`
}

func LoadEnv(envName string) (*Env, error) {
	if envName == "" {
		envName = "dev"
	}

	viper.AddConfigPath("./config")
	viper.SetConfigName(envName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var env Env
	if err := viper.Unmarshal(&env); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	if err := env.Validate(); err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}

	return &env, nil
}

func (e *Env) Validate() error {
	if e.ServerPort == "" ||
		e.ClientURL == "" {
		return fmt.Errorf("missing server configuration values")
	}

	if e.DBHost == "" ||
		e.DBUser == "" ||
		e.DBUserPassword == "" ||
		e.DBName == "" ||
		e.DBPort == "" ||
		e.DBSSLMode == "" {
		return fmt.Errorf("missing database configuration values")
	}

	if e.JWTAccessTokenSecret == "" ||
		e.JWTRefreshTokenSecret == "" ||
		e.JWTIssuer == "" ||
		e.JWTAudience == "" ||
		e.JWTTokenExpirationDays == 0 ||
		e.JWTRefreshTokenExpirationDays == 0 {
		return fmt.Errorf("missing JWT configuration values")
	}

	return nil
}
