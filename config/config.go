package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	APIPort string `yaml:"api-port"`
}

var AppConf Config

func Init() {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(data, &AppConf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
