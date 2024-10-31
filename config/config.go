package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	APIPort     string `yaml:"api-port"`
	Mode        string `yaml:"mode"`
	FilepathLog string `yaml:"filepath-log"`
	ConnectDB   string `yaml:"db-connect"`
	NameDB      string `yaml:"db-name"`
	UserDB      string `yaml:"db-user"`
	PasswordDB  string `yaml:"db-password"`
}

var AppConf Config

func NewConfig() (*Config, error) {
	configPath, err := filepath.Abs(filepath.Join("..", "config", "config.yaml"))
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
