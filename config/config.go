package config

import (
	"log"
	"os"

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

func Init(filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(data, &AppConf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
