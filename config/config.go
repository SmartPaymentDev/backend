package config

import (
	"github.com/caarlos0/env"
	"log"
	"sync"
)

type Config struct {
	Port         int    `env:"HTTP_PORT"`
	Host         string `env:"DB_HOST"`
	PostgresPort int    `env:"DB_PORT"`
	Username     string `env:"DB_USER"`
	Password     string `env:"DB_PASSWORD"`
	Dbname       string `env:"DB_DATABASE"`
	SECREET_KEY  string `env:"SECREET_KEY"`
}

var (
	configuration Config
	mutex         sync.Once
)

func GetConfig() Config {
	mutex.Do(func() {
		configuration = newConfig()
	})

	return configuration
}

func newConfig() Config {
	var cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err.Error())
	}

	return cfg
}
