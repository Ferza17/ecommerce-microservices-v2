package config

import (
	"fmt"
	"github.com/xhit/go-str2duration/v2"
	"log"
	"time"

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

	RabbitMQUsername string `mapstructure:"RABBITMQ_USERNAME"`
	RabbitMQPassword string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQHost     string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort     string `mapstructure:"RABBITMQ_PORT"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDB       int    `mapstructure:"REDIS_DB"`

	JwtAccessTokenSecret               string `mapstructure:"JWT_REFRESH_TOKEN_SECRET"`
	JwtAccessTokenExpirationTimeString string `mapstructure:"JWT_ACCESS_TOKEN_EXPIRATION_TIME"`
	JwtAccessTokenExpirationTime       time.Duration

	JwtRefreshTokenSecret               string `mapstructure:"JWT_REFRESH_TOKEN_SECRET"`
	JwtRefreshTokenExpirationTimeString string `mapstructure:"JWT_REFRESH_TOKEN_EXPIRATION_TIME"`
	JwtRefreshTokenExpirationTime       time.Duration

	RpcHost string `mapstructure:"RPC_HOST"`
	RpcPort string `mapstructure:"RPC_PORT"`

	VerificationUserLoginUrl string `mapstructure:"VERIFICATION_USER_LOGIN_URL"`

	JaegerTelemetryHost string `mapstructure:"JAEGER_TELEMETRY_HOST"`
	JaegerTelemetryPort string `mapstructure:"JAEGER_TELEMETRY_PORT"`
}

func SetConfig(path string) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
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

	if c.RedisHost == "" {
		log.Fatalf("SetConfig | REDIS_HOST is required")
	}

	if c.RedisPort == "" {
		log.Fatalf("SetConfig | REDIS_PORT is required")
	}

	if c.JwtAccessTokenSecret == "" {
		log.Fatalf("SetConfig | JWT_REFRESH_TOKEN_SECRET is required")
	}

	if c.JwtAccessTokenExpirationTimeString == "" {
		log.Fatalf("SetConfig | JWT_ACCESS_TOKEN_EXPIRATION_TIME is required")
	}
	c.JwtAccessTokenExpirationTime, err = str2duration.ParseDuration(c.JwtAccessTokenExpirationTimeString)
	if err != nil {
		log.Fatalf("SetConfig | JWT_ACCESS_TOKEN_EXPIRATION_TIME is invalid")
	}

	if c.JwtRefreshTokenSecret == "" {
		log.Fatalf("SetConfig | JWT_REFRESH_TOKEN_SECRET is required")
	}
	if c.JwtRefreshTokenExpirationTimeString == "" {
		log.Fatalf("SetConfig | JWT_REFRESH_TOKEN_EXPIRATION_TIME is required")
	}
	c.JwtRefreshTokenExpirationTime, err = str2duration.ParseDuration(c.JwtRefreshTokenExpirationTimeString)
	if err != nil {
		log.Fatalf("SetConfig | JWT_REFRESH_TOKEN_EXPIRATION_TIME is invalid")
	}

	if c.VerificationUserLoginUrl == "" {
		log.Fatalf("SetConfig | VERIFICATION_USER_LOGIN_URL is required")
	}

	//if c.RedisPassword == "" {
	//	log.Fatalf("SetConfig | REDIS_PASSWORD is required")
	//}

	if c.JaegerTelemetryHost == "" {
		log.Fatalf("SetConfig | JAEGER_TELEMETRY_HOST is required")
	}

	if c.JaegerTelemetryPort == "" {
		log.Fatalf("SetConfig | JAEGER_TELEMETRY_PORT is required")
	}

	viper.WatchConfig()
}
