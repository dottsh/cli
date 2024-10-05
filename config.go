package main

import (
	"fmt"
	"os"

	"github.com/mateothegreat/go-util/files"
	"github.com/mateothegreat/go-util/validation"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name     string
	Packages []Group `yaml:"packages"`
}

type Group struct {
	Group string `yaml:"group"`
	Items []Item `yaml:"items"`
	Repo  string `yaml:"repo" required:"false"`
}
type ItemType string

const (
	File ItemType = "file"
	Brew ItemType = "brew"
)

type Item struct {
	Name string   `yaml:"name"`
	Type ItemType `yaml:"type" required:"false"`
	Dest string   `yaml:"dest" required:"false"`
}

func GetConfig() (*Config, error) {
	config := &Config{}

	configPath := files.WalkFile("test.yaml", 10)

	f, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, config)
	if err != nil {
		return nil, err
	}

	// Validate the final config struct.
	emptyFields, err := validation.ValidateStructFields(config, "")
	if err != nil {
		return nil, err
	}
	if len(emptyFields) > 0 {
		return nil, fmt.Errorf("empty fields: %v", emptyFields)
	}

	return config, nil
}
