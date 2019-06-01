package config

import (
	"sync"

	"github.com/caarlos0/env"
)

type Config struct {
	ELASTICSEARCH_URL string `env:"ELASTICSEARCH_URL" envDefault:"http://127.0.0.1:9200"`
	BUCKET_NAME       string `env:"BUCKET_NAME" envDefault:"sephora-k8s-logging"`
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
		env.Parse(instance)
	})
	return instance
}
