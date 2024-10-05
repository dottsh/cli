package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mateothegreat/go-util/files"
	"github.com/mateothegreat/go-util/validation"
)

type Config struct {
	Name     string
	Packages []Group `yaml:"packages"`
}

type Group struct {
	Group string `yaml:"group"`
	Items []Item `yaml:"items"`
}

type Item struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func GetConfig() (*Config, error) {
	config := &Config{}

	configPath := files.WalkFile("test.yaml", 10)

	f, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(f, config)
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
