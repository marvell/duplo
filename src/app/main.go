package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

const VERSION = "0.1.0"

func main() {
	initConfig()

	if len(flag.Args()) > 0 && flag.Args()[0] == "version" {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if config.LogLevel == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	}

	log.Infoln("Starting server on " + config.BindAddr + " ...")
	err := startServer(config.BindAddr)
	if err != nil {
		log.Fatalln(err)
	}
}
