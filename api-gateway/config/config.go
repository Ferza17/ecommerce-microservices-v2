package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var c *Config

func Get() *Config {
	return c
}

type Config struct {
	Env         string `mapstructure:"ENV"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
	AppURL      string `mapstructure:"APP_URL"`

	MongoUsername     string `mapstructure:"MONGO_USERNAME"`
	MongoPassword     string `mapstructure:"MONGO_PASSWORD"`
	MongoHost         string `mapstructure:"MONGO_HOST"`
	MongoPort         string `mapstructure:"MONGO_PORT"`
	MongoDatabaseName string `mapstructure:"MONGO_DATABASE_NAME"`

	RabbitMQUsername string `mapstructure:"RABBITMQ_USERNAME"`
	RabbitMQPassword string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQHost     string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort     string `mapstructure:"RABBITMQ_PORT"`

	ProductServiceURL string `mapstructure:"PRODUCT_SERVICE_URL"`
	UserServiceURL    string `mapstructure:"USER_SERVICE_URL"`

	HttpHost string `mapstructure:"HTTP_HOST"`
	HttpPort string `mapstructure:"HTTP_PORT"`
}

func SetConfig(path string) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("config not found: %s", err.Error()))
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("SetConfig | could not parse config: %v", err)
	}

	if c.Env == "" {
		c.Env = "local"
	}

	if c.ServiceName == "" {
		log.Fatalf("SetConfig | SERVICE_NAME is required")
	}

	if c.HttpHost == "" {
		log.Fatalf("SetConfig | HTTP_HOST is required")
	}

	if c.HttpPort == "" {
		log.Fatalf("SetConfig | HTTP_PORT is required")
	}

	if c.MongoUsername == "" {
		log.Fatalf("SetConfig | MONGO_USERNAME is required")
	}

	if c.MongoPassword == "" {
		log.Fatalf("SetConfig | MONGO_PASSWORD is required")
	}

	if c.MongoHost == "" {
		log.Fatalf("SetConfig | MONGO_HOST is required")
	}

	if c.MongoPort == "" {
		log.Fatalf("SetConfig | MONGO_PORT is required")
	}

	if c.MongoDatabaseName == "" {
		log.Fatalf("SetConfig | MONGO_DATABASE_NAME is required")
	}

	if c.RabbitMQUsername == "" {
		log.Fatalf("SetConfig | RABBITMQ_USERNAME is required")
	}
	if c.RabbitMQPassword == "" {
		log.Fatalf("SetConfig | RABBITMQ_PASSWORD is required")
	}

	if c.RabbitMQHost == "" {
		log.Fatalf("SetConfig | RABBITMQ_HOST is required")
	}

	if c.RabbitMQPort == "" {
		log.Fatalf("SetConfig | RABBITMQ_PORT is required")
	}

	if c.ProductServiceURL == "" {
		log.Fatalf("SetConfig | PRODUCT_SERVICE_URL is required")
	}

	if c.UserServiceURL == "" {
		log.Fatalf("SetConfig | USER_SERVICE_URL is required")
	}

	viper.WatchConfig()
}
