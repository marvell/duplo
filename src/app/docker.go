package main

import (
	"fmt"
)

func pullImage(s *spec, verbose bool) error {
	log.Infof("Pulling new image (%s:%s) ...", s.Image, s.Tag)

	cmd := fmt.Sprintf("docker pull %s", s.ImageWithTag())
	return runCommand(cmd, verbose)
}

func stopContainer(s *spec) error {
	log.Infof("Stopping container (%s) ...", s.Name)

	cmd := fmt.Sprintf("docker stop %s", s.NameWithPrefix())
	return runCommand(cmd, false)
}
func deleteContainer(s *spec) error {
	log.Infof("Removing container (%s) ...", s.Name)

	cmd := fmt.Sprintf("docker rm %s", s.NameWithPrefix())
	return runCommand(cmd, false)
}
func runContainer(s *spec, verbose bool) error {
	log.Infof("Running container (%s) ...", s.Name)

	cmd := "docker run "
	cmd += fmt.Sprintf("--name=%s ", s.NameWithPrefix())
	for key, values := range s.Args {
		for _, value := range values {
			cmd += fmt.Sprintf("--%s=%s ", key, value)
		}
	}
	cmd += fmt.Sprintf("%s:%s ", s.Image, s.Tag)
	cmd += s.Command

	return runCommand(cmd, verbose)
}
