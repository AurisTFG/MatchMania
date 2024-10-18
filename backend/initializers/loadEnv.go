package initializers

import (
	"log"

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
}

func LoadEnv() {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("Unable to decode config into struct: %s", err)
	}

	log.Println("Loaded environment variables")
}
