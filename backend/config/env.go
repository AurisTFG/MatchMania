package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	IsDev bool

	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	ClientURL  string `mapstructure:"CLIENT_URL"`

	DBHost         string `mapstructure:"DB_HOST"`
	DBUser         string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBSSLMode      string `mapstructure:"DB_SSLMODE"`

	JWTAccessTokenSecret    string        `mapstructure:"JWT_ACCESS_TOKEN_SECRET"`
	JWTRefreshTokenSecret   string        `mapstructure:"JWT_REFRESH_TOKEN_SECRET"`
	JWTIssuer               string        `mapstructure:"JWT_ISSUER"`
	JWTAudience             string        `mapstructure:"JWT_AUDIENCE"`
	JWTAccessTokenDuration  time.Duration `mapstructure:"JWT_ACCESS_TOKEN_DURATION"`
	JWTRefreshTokenDuration time.Duration `mapstructure:"JWT_REFRESH_TOKEN_DURATION"`

	UserEmail         string `mapstructure:"USER_EMAIL"`
	UserPassword      string `mapstructure:"USER_PASSWORD"`
	ModeratorEmail    string `mapstructure:"MODERATOR_EMAIL"`
	ModeratorPassword string `mapstructure:"MODERATOR_PASSWORD"`
	AdminEmail        string `mapstructure:"ADMIN_EMAIL"`
	AdminPassword     string `mapstructure:"ADMIN_PASSWORD"`
}

func LoadEnv(envName string) (*Env, error) {
	if envName == "" {
		envName = "dev"
	}

	viper.SetConfigFile("./config/.env." + envName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var env Env
	env.IsDev = envName == "dev"

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	if err := viper.Unmarshal(&env); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	if err := env.Validate(); err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}

	return &env, nil
}

func (e *Env) getVarsFromEnv() {
	e.ServerHost = os.Getenv("SERVER_HOST")
	e.ServerPort = os.Getenv("SERVER_PORT")
	e.ClientURL = os.Getenv("CLIENT_URL")
	// etc..
}

func (e *Env) Validate() error {
	if e.ServerPort == "" ||
		e.ClientURL == "" ||
		e.ServerHost == "" {
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
		e.JWTAccessTokenDuration == 0 ||
		e.JWTRefreshTokenDuration == 0 {
		return fmt.Errorf("missing JWT configuration values")
	}

	return nil
}
