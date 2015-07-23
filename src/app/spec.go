package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type spec struct {
	Name    string
	Image   string
	Tag     string
	Command string
	Args    map[string][]string
}

var (
	hasMultipleValues = []string{"attach", "add-host", "cap-add", "cap-drop", "device", "dns", "dns-search", "env", "env-file", "expose", "label", "label-file", "link", "lxc-conf", "publish", "security-opt", "ulimit", "volume", "volumes-from"}
)

func newSpec(name string) *spec {
	return &spec{Name: name}
}

func (s *spec) filePath() string {
	return path.Join(config.SpecPath, s.Name+".yml")
}

func (s *spec) exists() bool {
	_, err := os.Open(s.filePath())
	return !os.IsNotExist(err)
}

func (s *spec) load() error {
	if s.exists() {
		d, err := ioutil.ReadFile(s.filePath())
		if err != nil {
			return err
		}

		o := map[string]interface{}{}

		err = yaml.Unmarshal(d, &o)
		if err != nil {
			return err
		}

		return s.parse(o)
	}

	return errors.New("spec not found")
}

func (s *spec) parse(o map[string]interface{}) error {
	s.Args = map[string][]string{}

	for k, v := range o {
		switch k {
		case "image":
			s.Image = v.(string)
		case "command":
			s.Command = v.(string)
		default:
			if len(k) == 1 {
				return errors.New("do not use single-character argument")
			}

			switch v.(type) {
			case []interface{}:
				for _, i := range v.([]interface{}) {
					s.Args[k] = append(s.Args[k], fmt.Sprint(i))
				}
			default:
				s.Args[k] = append(s.Args[k], fmt.Sprint(v))
			}
		}
	}

	s.Args["detach"] = []string{"true"}

	if s.Image == "" {
		return errors.New("isn't set image in spec-file")
	}

	return nil
}
