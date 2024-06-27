package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

type Config struct {
	Server struct {
		Host string
		Port int
	}
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
		Name     string
	}
	Logging struct {
		Level  string
		Format string
	}
	Jwt struct {
		SecretKey         string
		ExpirationMinutes int
	}
	Features struct {
		EnableFeatureX bool
		EnableFeatureY bool
	}
	Roles struct {
		Admin string
		User  string
	}
	Status struct {
		Completed string
		Progress  string
		Planned   string
	}
}

func loadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err.Error())
	}

	return &config
}

var lock = &sync.Mutex{}

var singleInstance *Config

func GetInstance() *Config {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = loadConfig()
		}
	}
	return singleInstance
}
