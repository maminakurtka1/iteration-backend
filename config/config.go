package config

import (
	"github.com/apex/log"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port int `json:"port"`
}

func LoadConfig() *Configuration {
	configuration := &Configuration{}
	err := gonfig.GetConf("config.json", configuration)
	if err != nil {
		log.WithError(err).Fatal("error loading configuration")
	}
	return configuration
}
