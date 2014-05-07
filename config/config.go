package config

import (
	// standard library
	"io/ioutil"

	// external packages
	"launchpad.net/goyaml"
)

// The Config represents a configuration for the application.
type Config struct {
	ComPort                string `yaml:"comPort,omitempty"`
	DBpediaEndpoint        string `yaml:"dbpediaEndpoint,omitempty"`
	DBpediaEndpointTimeout int    `yaml:"dbpediaEndpointTimeout,omitempty"`
}

// CONFIG is the global configuration.
var CONFIG *Config = new(Config)

// ReadConfig() reads the config yaml file and set global config variables.
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
