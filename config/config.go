package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppPort    string `yaml:"app-port"`
	AppMode    string `yaml:"app-mode"`
	AppFilelog string `yaml:"app-filelog"`
	DbConnect  string `yaml:"db-connect"`
	DbName     string `yaml:"db-name"`
	DbUser     string `yaml:"db-user"`
	DbPassword string `yaml:"db-password"`
}

var AppConf Config

func NewConfig() (*Config, error) {
	configPath, err := filepath.Abs(filepath.Join("config", "config.yaml"))
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
