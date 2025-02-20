package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type AppConfig struct {
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

var Config *AppConfig

func InitConfig(path string) {
	var (
		data []byte
		err  error
	)
	data, err = os.ReadFile(path) //ReadFile从ioutil移动到了os包

	if err != nil {
		log.Fatalf("read config file from %s error: %s", path, err)
	}

	Config = &AppConfig{}

	if err = yaml.Unmarshal(data, Config); err != nil {
		log.Fatalf("parse config file %s error: %s", path, err)
	}

}
