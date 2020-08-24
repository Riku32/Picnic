package config

import (
	"fmt"
	"os"

	"github.com/Riku32/Picnic/logger"
	"gopkg.in/yaml.v2"
)

// Config : config object
type Config struct {
	Token string `yaml:"token"`
}

// Load : load config file
func Load() Config {
	configf, err := os.Open("./config.yaml")
	if err != nil {
		logger.Panic(fmt.Sprintf("could not find config"))
	}
	defer configf.Close()

	var config Config

	err = yaml.NewDecoder(configf).Decode(&config)
	if err != nil {
		logger.Panic(fmt.Sprintf("invalid config file"))
	}

	return config
}
