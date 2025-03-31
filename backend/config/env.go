package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	IsDev  bool
	IsProd bool

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

var invalidString = "INVALID"
var invalidDuration = time.Duration(0)

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

	// must set all defaults, otherwise viper will not read from env
	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if env.IsDev {
			return nil, fmt.Errorf("unable to read config file: %w", err)
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
	viper.SetDefault("SERVER_HOST", invalidString)
	viper.SetDefault("SERVER_PORT", invalidString)
	viper.SetDefault("CLIENT_URL", invalidString)

	viper.SetDefault("DB_HOST", invalidString)
	viper.SetDefault("DB_USER", invalidString)
	viper.SetDefault("DB_PASSWORD", invalidString)
	viper.SetDefault("DB_NAME", invalidString)
	viper.SetDefault("DB_PORT", invalidString)
	viper.SetDefault("DB_SSLMODE", invalidString)

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
	if e.ServerHost == invalidString ||
		e.ServerPort == invalidString ||
		e.ClientURL == invalidString {
		return fmt.Errorf("missing server configuration values")
	}

	if e.DBHost == invalidString ||
		e.DBUser == invalidString ||
		e.DBUserPassword == invalidString ||
		e.DBName == invalidString ||
		e.DBPort == invalidString ||
		e.DBSSLMode == invalidString {
		return fmt.Errorf("missing database configuration values")
	}

	if e.JWTAccessTokenSecret == invalidString ||
		e.JWTRefreshTokenSecret == invalidString ||
		e.JWTIssuer == invalidString ||
		e.JWTAudience == invalidString ||
		e.JWTAccessTokenDuration == invalidDuration ||
		e.JWTRefreshTokenDuration == invalidDuration {
		return fmt.Errorf("missing JWT configuration values")
	}

	return nil
}
