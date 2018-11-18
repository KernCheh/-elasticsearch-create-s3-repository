package config

import (
	"sync"

	"github.com/caarlos0/env"
)

type Config struct {
	ELASTICSEARCH_URL     string `env:"ELASTICSEARCH_URL" envDefault:"http://127.0.0.1:9200"`
	AWS_ACCESS_KEY_ID     string `env:"AWS_ACCESS_KEY_ID,required"`
	AWS_SECRET_ACCESS_KEY string `env:"AWS_SECRET_ACCESS_KEY,required"`
	APP_ENV               string `env:"APP_ENV" envDefault:"development"`
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
