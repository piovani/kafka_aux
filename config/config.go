package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// API Rest
	ApiRestPort int `envconfig:"API_REST_PORT"`

	// REDIS
	RedisHost     string `envconfig:"REDIS_HOST"`
	RedisPassword string `envconfig:"REDIS_PASSWORD"`

	// KAFKA
	KafkaHost  string `envconfig:"KAFKA_HOST"`
	KafkaTopic string `envconfig:"KAFKA_TOPIC"`
}

var Env Config

func StarConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}
