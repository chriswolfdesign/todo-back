package config

import (
	"io/ioutil"
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

func ReadConfigFile(fileName string) Config {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Unable to read file:", err)
		os.Exit(1)
	}

	conf := Config{}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal("Unable to unmarshal config:", err)
		os.Exit(1)
	}

	return conf
}
