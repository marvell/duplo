package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "0.1.0"

func main() {
	initConfig()
	initLogger()

	if len(flag.Args()) > 0 && flag.Args()[0] == "version" {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	log.Infof("Starting server on %s ...", config.BindAddr)
	err := startServer(config.BindAddr)
	if err != nil {
		log.Fatal(err)
	}
}
