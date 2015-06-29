package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func runCommand(command string, verbose bool) error {
	command = strings.TrimSpace(command)

	if command == "" {
		return errors.New("empty command")
	}

	log.WithField("cmd", command).Debug("Running command...")

	command_parts := strings.Split(command, " ")
	command = command_parts[0]
	args := []string{}
	if len(command_parts) > 1 {
		args = command_parts[1:]
	}

	cmd := exec.Command(command, args...)

	errPipe := bytes.NewBuffer(make([]byte, 1024))
	cmd.Stderr = errPipe

	if verbose == true {
		cmd.Stdout = os.Stdout
	}

	cmd.Run()

	err, _ := ioutil.ReadAll(errPipe)
	errMsg := strings.Split(strings.Trim(string(err), "\x00"), "\n")[0]
	if len(errMsg) > 0 {
		return parseError(errMsg)
	}

	return nil
}

func parseError(err string) error {
	r, _ := regexp.Compile("msg=\"([^\"]+)\"")
	msg := r.FindStringSubmatch(err)
	if len(msg) > 0 {
		err = msg[1]
	}

	return errors.New(err)
}
