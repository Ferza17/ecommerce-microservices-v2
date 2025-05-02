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

	PostgresUsername     string `mapstructure:"POSTGRES_USERNAME"`
	PostgresPassword     string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresSSLMode      string `mapstructure:"POSTGRES_SSL_MODE"`
	PostgresHost         string `mapstructure:"POSTGRES_HOST"`
	PostgresPort         string `mapstructure:"POSTGRES_PORT"`
	PostgresDatabaseName string `mapstructure:"POSTGRES_DATABASE_NAME"`

	MongoUsername     string `mapstructure:"MONGO_USERNAME"`
	MongoPassword     string `mapstructure:"MONGO_PASSWORD"`
	MongoHost         string `mapstructure:"MONGO_HOST"`
	MongoPort         string `mapstructure:"MONGO_PORT"`
	MongoDatabaseName string `mapstructure:"MONGO_DATABASE_NAME"`

	RabbitMQUsername string `mapstructure:"RABBITMQ_USERNAME"`
	RabbitMQPassword string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQHost     string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort     string `mapstructure:"RABBITMQ_PORT"`

	RpcHost string `mapstructure:"RPC_HOST"`
	RpcPort string `mapstructure:"RPC_PORT"`
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

	if c.PostgresUsername == "" {
		log.Fatalf("SetConfig | POSTGRES_USERNAME is required")
	}

	if c.PostgresPassword == "" {
		log.Fatalf("SetConfig | POSTGRES_PASSWORD is required")
	}

	if c.PostgresHost == "" {
		log.Fatalf("SetConfig | POSTGRES_HOST is required")
	}

	if c.PostgresPort == "" {
		log.Fatalf("SetConfig | POSTGRES_PORT is required")
	}

	if c.PostgresDatabaseName == "" {
		log.Fatalf("SetConfig | POSTGRES_DATABASE_NAME is required")
	}

	if c.RpcHost == "" {
		log.Fatalf("SetConfig | RPC_HOST is required")
	}

	if c.RpcPort == "" {
		log.Fatalf("SetConfig | RPC_PORT is required")
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

	viper.WatchConfig()
}
