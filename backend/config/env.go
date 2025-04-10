package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	IsDev  bool
	IsProd bool

	DatabaseURL string `mapstructure:"DATABASE_URL"`
	ServerURL   string `mapstructure:"SERVER_URL"`
	ClientURL   string `mapstructure:"CLIENT_URL"`

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

var (
	invalidString   = "INVALID"
	invalidDuration = time.Duration(0)
)

func LoadEnv(envName string) (*Env, error) {
	var env Env

	var filePostfix string

	env.IsDev = envName == "dev" || envName == "development"
	env.IsProd = envName == "prod" || envName == "production"

	if !env.IsDev && !env.IsProd {
		return nil, fmt.Errorf("invalid environment name: %s", envName)
	}

	if env.IsDev {
		filePostfix = "development"
	} else {
		filePostfix = "production"
	}

	viper.SetConfigFile("./.env." + filePostfix)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	setDefaults() // must set all defaults, otherwise viper will not read from env

	if err := viper.ReadInConfig(); err != nil {
		// Only return an error if it's not a "file not found" error
		var configFileNotFoundErr viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundErr) {
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

func setDefaults() {
	viper.SetDefault("DATABASE_URL", invalidString)
	viper.SetDefault("SERVER_URL", invalidString)
	viper.SetDefault("CLIENT_URL", invalidString)

	viper.SetDefault("JWT_ACCESS_TOKEN_SECRET", invalidString)
	viper.SetDefault("JWT_REFRESH_TOKEN_SECRET", invalidString)
	viper.SetDefault("JWT_ISSUER", invalidString)
	viper.SetDefault("JWT_AUDIENCE", invalidString)
	viper.SetDefault("JWT_ACCESS_TOKEN_DURATION", invalidDuration)
	viper.SetDefault("JWT_REFRESH_TOKEN_DURATION", invalidDuration)

	viper.SetDefault("USER_EMAIL", invalidString)
	viper.SetDefault("USER_PASSWORD", invalidString)
	viper.SetDefault("MODERATOR_EMAIL", invalidString)
	viper.SetDefault("MODERATOR_PASSWORD", invalidString)
	viper.SetDefault("ADMIN_EMAIL", invalidString)
	viper.SetDefault("ADMIN_PASSWORD", invalidString)
}

func (e *Env) Validate() error {
	if e.DatabaseURL == invalidString {
		return errors.New("missing database URL")
	}

	if e.ServerURL == invalidString {
		return errors.New("missing server URL")
	}

	if e.ClientURL == invalidString {
		return errors.New("missing client URL")
	}

	if e.JWTAccessTokenSecret == invalidString ||
		e.JWTRefreshTokenSecret == invalidString ||
		e.JWTIssuer == invalidString ||
		e.JWTAudience == invalidString ||
		e.JWTAccessTokenDuration == invalidDuration ||
		e.JWTRefreshTokenDuration == invalidDuration {
		return errors.New("missing JWT configuration values")
	}

	if e.IsDev {
		if e.UserEmail == invalidString ||
			e.UserPassword == invalidString ||
			e.ModeratorEmail == invalidString ||
			e.ModeratorPassword == invalidString ||
			e.AdminEmail == invalidString ||
			e.AdminPassword == invalidString {
			return errors.New("missing default user credentials")
		}
	}

	return nil
}
