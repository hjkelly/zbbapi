package config

import (
	"os"
	"sync"
)

type Config struct {
	MONGO_URL string
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			MONGO_URL: os.Getenv("MONGO_URL"),
		}
	})
	return config
}
