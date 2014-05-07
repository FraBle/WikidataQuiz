package config

import (
	"io/ioutil"
	"launchpad.net/goyaml"
)

type Config struct {
	ComPort string `yaml:"comPort,omitempty"`
}

var CONFIG *Config = new(Config)

func ReadConfig() error {
	file, err := ioutil.ReadFile("../src/github.com/FraBle/WikidataQuiz/config/config.yaml")
	if err != nil {
		return err
	}
	if err = goyaml.Unmarshal(file, CONFIG); err != nil {
		return err
	}
	return nil
}
