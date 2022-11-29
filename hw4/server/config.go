package main

import (
	"github.com/caarlos0/env/v6"
)

type config struct {
	DatabaseAddr string `env:"DBADDR" default:""`
}

func getConfig() (*config, error) {
	cfg := &config{}
	err := env.Parse(cfg)
	return cfg, err
}
