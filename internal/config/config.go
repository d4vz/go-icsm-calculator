package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	cfg *Config
)

type Config struct {
	EnableBackpressure      bool
	BackpressureMaxAttempts int
}

func Load() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(dir)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found, using defaults")
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	viper.SetDefault("ENABLE_BACKPRESSURE", false)
	viper.SetDefault("BACKPRESSURE_MAX_ATTEMPTS", 5)

	cfg = &Config{
		EnableBackpressure:      viper.GetBool("ENABLE_BACKPRESSURE"),
		BackpressureMaxAttempts: viper.GetInt("BACKPRESSURE_MAX_ATTEMPTS"),
	}
}

func Get() *Config {
	return cfg
}
