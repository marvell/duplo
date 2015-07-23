package main

import (
	"github.com/marvell/senger"
)

var log *senger.Logger

func initLogger() {
	log = senger.NewDefaultLogger()

	lvl := senger.ParseLevel(config.LogLevel)
	if lvl != nil {
		log.SetLevel(lvl)
	}
}
