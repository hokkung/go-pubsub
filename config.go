package go_pubsub

import (
	"github.com/kelseyhightower/envconfig"
)

// Config config
type Config struct {
	ProjectID string `envconfig:"project_id" required:"true"`
}

// NewConfig create an instance
func NewConfig() *Config {
	var c Config
	err := envconfig.Process("pubsub", &c)
	if err != nil {
		panic(err)
	}
	return &c
}
