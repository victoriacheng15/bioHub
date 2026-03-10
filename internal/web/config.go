package web

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Social struct {
	Platform string `yaml:"Platform"`
	Icon     string `yaml:"Icon"`
	URL      string `yaml:"URL"`
}

type Link struct {
	Name string `yaml:"Name"`
	URL  string `yaml:"URL"`
}

type Params struct {
	Avatar   string   `yaml:"Avatar"`
	Name     string   `yaml:"Name"`
	Headline string   `yaml:"Headline"`
	Socials  []Social `yaml:"Socials"`
	Links    []Link   `yaml:"Links"`
}

type Config struct {
	Params Params `yaml:"Params"`
}

// LoadConfig reads a YAML configuration file and unmarshals it into a Config struct
func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}
