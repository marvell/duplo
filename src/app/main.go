package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

const version = "0.1.0"

var (
	specsPath string
	bindAddr  string
	debug     bool
)

func init() {
	flag.StringVar(&specsPath, "dir", "/etc/duplo", "specs directory")
	flag.StringVar(&bindAddr, "bind", "0.0.0.0:5732", "bind addr")
	flag.BoolVar(&debug, "debug", false, "debug mode")
}

func main() {
	var err error

	flag.Parse()

	if len(flag.Args()) < 1 || len(flag.Args()) > 2 {
		fmt.Printf("usage: %s <spec> <tag>\n", os.Args[0])
		fmt.Printf("usage: %s version\n", os.Args[0])
		os.Exit(1)
	}

	if flag.Args()[0] == "version" {
		fmt.Println(version)
		os.Exit(0)
	}

	if debug == true {
		log.SetLevel(log.DebugLevel)
	}

	if flag.Args()[0] == "server" {
		log.Infoln("Starting server on " + bindAddr + " ...")
		err = startServer(bindAddr)
		if err != nil {
			log.Fatalln(err)
		}

		return
	}

	spec := newSpec(flag.Args()[0])
	err = spec.load()
	if err != nil {
		log.Fatalln(err)
	}

	spec.Tag = flag.Args()[1]

	err = pullImage(spec, false)
	if err != nil {
		log.Fatalln(err)
	}
	stopContainer(spec)
	deleteContainer(spec)
	err = runContainer(spec, false)
	if err != nil {
		log.Fatalln(err)
	}
}
