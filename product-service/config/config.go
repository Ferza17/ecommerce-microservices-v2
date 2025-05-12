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

	ElasticsearchUsername string `mapstructure:"ELASTICSEARCH_USERNAME"`
	ElasticsearchPassword string `mapstructure:"ELASTICSEARCH_PASSWORD"`
	ElasticsearchHost     string `mapstructure:"ELASTICSEARCH_HOST"`
	ElasticsearchPort     string `mapstructure:"ELASTICSEARCH_PORT"`

	RabbitMQUsername string `mapstructure:"RABBITMQ_USERNAME"`
	RabbitMQPassword string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQHost     string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort     string `mapstructure:"RABBITMQ_PORT"`

	RpcHost string `mapstructure:"RPC_HOST"`
	RpcPort string `mapstructure:"RPC_PORT"`

	JaegerTelemetryHost string `mapstructure:"JAEGER_TELEMETRY_HOST"`
	JaegerTelemetryPort string `mapstructure:"JAEGER_TELEMETRY_PORT"`
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

	// uncomment this if you have username & password configured on your elasticsearch
	//if c.ElasticsearchUsername == "" {
	//	log.Fatalf("SetConfig | ELASTICSEARCH_USERNAME is required")
	//}
	//if c.ElasticsearchPassword == "" {
	//	log.Fatalf("SetConfig | ELASTICSEARCH_PASSWORD is required")
	//}
	if c.ElasticsearchHost == "" {
		log.Fatalf("SetConfig | ELASTICSEARCH_HOST is required")
	}
	if c.ElasticsearchPort == "" {
		log.Fatalf("SetConfig | ELASTICSEARCH_PORT is required")
	}
	if c.JaegerTelemetryHost == "" {
		log.Fatalf("SetConfig | JAEGER_TELEMETRY_HOST is required")
	}

	if c.JaegerTelemetryPort == "" {
		log.Fatalf("SetConfig | JAEGER_TELEMETRY_PORT is required")
	}
	
	viper.WatchConfig()
}
