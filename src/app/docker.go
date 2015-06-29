package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func pullImage(s *spec, verbose bool) error {
	log.WithFields(log.Fields{
		"name": s.Image,
		"tag":  s.Tag,
	}).Info("Pulling new image")
	return runCommand("docker pull "+s.Image+":"+s.Tag, verbose)
}

func stopContainer(s *spec) error {
	log.WithField("name", s.Name).Info("Stopping container")
	return runCommand("docker stop "+s.Name, false)
}
func deleteContainer(s *spec) error {
	log.WithField("name", s.Name).Info("Removing container")
	return runCommand("docker rm "+s.Name, false)
}
func runContainer(s *spec, verbose bool) error {
	log.WithField("name", s.Name).Info("Running container")

	cmd := "docker run "
	cmd += fmt.Sprintf("--name=%s ", s.Name)
	for key, values := range s.Args {
		for _, value := range values {
			cmd += fmt.Sprintf("--%s=%s ", key, value)
		}
	}
	cmd += fmt.Sprintf("%s:%s ", s.Image, s.Tag)
	cmd += s.Command

	return runCommand(cmd, verbose)
}
