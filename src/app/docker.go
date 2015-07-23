package main

import (
	"fmt"
)

func pullImage(s *spec, verbose bool) error {
	log.Infof("Pulling new image (%s:%s) ...", s.Image, s.Tag)
	return runCommand("docker pull "+s.Image+":"+s.Tag, verbose)
}

func stopContainer(s *spec) error {
	log.Infof("Stopping container (%s) ...", s.Name)
	return runCommand("docker stop "+s.Name, false)
}
func deleteContainer(s *spec) error {
	log.Infof("Removing container (%s) ...", s.Name)
	return runCommand("docker rm "+s.Name, false)
}
func runContainer(s *spec, verbose bool) error {
	log.Infof("Running container (%s) ...", s.Name)

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
