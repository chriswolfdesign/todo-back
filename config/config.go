package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func ReadConfigFile(fileName string) (*Config, error) {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Unable to read file:", err)
		return nil, err
	}

	conf := Config{}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal("Unable to unmarshal config:", err)
		return nil, err
	}

	return &conf, nil
}
