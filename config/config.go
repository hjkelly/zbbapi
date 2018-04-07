package config

import (
	"os"
	"sync"
)

// Config contains all configuration used by the entire app.
type Config struct {
	MongoURL string
}

var config *Config
var once sync.Once

// GetConfig returns the populated, singleton instance of the configuration.
func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			MongoURL: os.Getenv("MONGO_URL"),
		}
	})
	return config
}
