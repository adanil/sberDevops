package server

import (
	"log"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DATABASEADDR string `env:"DATABASEADDR" envDefault:""`
}

func ReadConfig() (*Config, error){
	config := Config{}
	err := env.Parse(&config)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &config, nil
}