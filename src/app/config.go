package main

import (
	"github.com/marvell/envconfig"
)

type Config struct {
	SpecPath string `env:"SPECS_PATH" default:"/etc/duplo" usage:"path to directory of specs"`
	BindAddr string `env:"BIND" default:"0.0.0.0:5732" usage:"binding address"`
	LogLevel string `env:"LOG_LEVEL" default:"INFO" usage:"log level"`
}

var config *Config

func initConfig() {
	config = new(Config)
	envconfig.Parse(config)
}
