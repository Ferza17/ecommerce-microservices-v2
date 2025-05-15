package config

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"github.com/xhit/go-str2duration/v2"
)

var c *Config

func Get() *Config {
	return c
}

type Config struct {
	Env        string `mapstructure:"ENV"`
	ConsulHost string `mapstructure:"CONSUL_HOST"`
	ConsulPort string `mapstructure:"CONSUL_PORT"`

	// From Consul
	ServiceName string

	JaegerTelemetryHost string
	JaegerTelemetryPort string

	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string

	PostgresHost         string
	PostgresPort         string
	PostgresUsername     string
	PostgresPassword     string
	PostgresDatabaseName string
	PostgresSSLMode      string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	JwtAccessTokenSecret          string
	JwtAccessTokenExpirationTime  time.Duration
	JwtRefreshTokenSecret         string
	JwtRefreshTokenExpirationTime time.Duration
	VerificationUserLoginUrl      string

	RpcHost string
	RpcPort string
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
		log.Fatal("SetConfig | env is required")
	}
	if c.ConsulHost == "" {
		log.Fatal("SetConfig | consul host is required")
	}
	if c.ConsulPort == "" {
		log.Fatal("SetConfig | consul port is required")
	}

	consulClient, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%s:%s", c.ConsulHost, c.ConsulPort),
	})
	if err != nil {
		log.Fatalf("SetConfig | could not connect to consul: %v", err)
	}

	// Get Consul Key / Value
	kv := consulClient.KV()
	wg := sync.WaitGroup{}

	// Telemetry Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/telemetry/jaeger/JAEGER_TELEMETRY_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get telemetry host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | telemetry host is required")
		}
		c.JaegerTelemetryHost = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/telemetry/jaeger/JAEGER_TELEMETRY_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get telemetry host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | telemetry host is required")
		}
		c.JaegerTelemetryPort = string(pair.Value)
	}()

	// RabbitMQ Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_USERNAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RABBITMQ_USERNAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RABBITMQ_USERNAME host is required")
		}
		c.RabbitMQUsername = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RABBITMQ_PASSWORD host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RABBITMQ_PASSWORD host is required")
		}
		c.RabbitMQPassword = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RABBITMQ_HOST host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RABBITMQ_HOST host is required")
		}
		c.RabbitMQHost = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RABBITMQ_PORT host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RABBITMQ_PORT host is required")
		}
		c.RabbitMQPort = string(pair.Value)
	}()

	// Postgres Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_USERNAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_USERNAME from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_USERNAME is required")
		}
		c.PostgresUsername = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_PASSWORD from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_PASSWORD is required")
		}
		c.PostgresPassword = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_SSL_MODE", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_SSL_MODE from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_SSL_MODE is required")
		}
		c.PostgresSSLMode = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_HOST from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_HOST is required")
		}
		c.PostgresHost = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_PORT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_PORT is required")
		}
		c.PostgresPort = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_DATABASE_NAME/USERS", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_DATABASE_NAME/USERS from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_DATABASE_NAME/PRODUCTS is required")
		}
		c.PostgresDatabaseName = string(pair.Value)
	}()

	// Redis Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/database/redis/REDIS_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get REDIS_HOST from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | REDIS_HOST is required")
		}
		c.RedisHost = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/database/redis/REDIS_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get REDIS_PORT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | REDIS_PORT is required")
		}
		c.RedisPort = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/database/redis/REDIS_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get REDIS_PASSWORD from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | REDIS_PASSWORD is required")
		}
		c.RedisPassword = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/database/redis/REDIS_DB", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get REDIS_DB from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | REDIS_DB is required")
		}
		redisDB, err := strconv.ParseInt(string(pair.Value), 10, 64)
		if err != nil {
			log.Fatalf("SetConfig | could not parse REDIS_DB to int: %v", err)
		}
		c.RedisDB = int(redisDB)
	}()

	// User Service Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/services/user/SERVICE_NAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
		}
		c.ServiceName = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/RPC_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RPC_HOST is required")
		}
		c.RpcHost = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/RPC_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RPC_PORT is required")
		}
		c.RpcPort = string(pair.Value)

		// Access Token Config
		pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/JWT_ACCESS_TOKEN_SECRET", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get JWT_ACCESS_TOKEN_SECRET from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | JWT_ACCESS_TOKEN_SECRET is required")
		}
		c.JwtAccessTokenSecret = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/JWT_ACCESS_TOKEN_EXPIRATION_TIME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get JWT_ACCESS_TOKEN_EXPIRATION_TIME from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | JWT_ACCESS_TOKEN_EXPIRATION_TIME is required")
		}
		expirationTime := string(pair.Value)
		c.JwtRefreshTokenExpirationTime, err = str2duration.ParseDuration(expirationTime)
		if err != nil {
			log.Fatalf("SetConfig | JWT_REFRESH_TOKEN_EXPIRATION_TIME is invalid")
		}

		// Refresh Token Token Config
		pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/JWT_REFRESH_TOKEN_SECRET", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get JWT_REFRESH_TOKEN_SECRET from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | JWT_REFRESH_TOKEN_SECRET is required")
		}
		c.JwtRefreshTokenSecret = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/JWT_REFRESH_TOKEN_EXPIRATION_TIME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get JWT_REFRESH_TOKEN_EXPIRATION_TIME from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | JWT_REFRESH_TOKEN_EXPIRATION_TIME is required")
		}
		expirationTime = string(pair.Value)
		c.JwtRefreshTokenExpirationTime, err = str2duration.ParseDuration(expirationTime)
		if err != nil {
			log.Fatalf("SetConfig | JWT_REFRESH_TOKEN_EXPIRATION_TIME is invalid")
		}

		// Verification User Login Url Config
		pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/VERIFICATION_USER_LOGIN_URL", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get VERIFICATION_USER_LOGIN_URL from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | VERIFICATION_USER_LOGIN_URL is required")
		}
		c.VerificationUserLoginUrl = string(pair.Value)
	}()

	wg.Wait()

	port, err := strconv.ParseInt(c.RpcPort, 10, 64)
	if err != nil {
		log.Fatalf("SetConfig | could not parse PORT to int: %v", err)
	}
	if err = consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    c.ServiceName,
		Address: c.RpcHost,
		Port:    int(port),
		Tags:    []string{"v1"},
	}); err != nil {
		log.Fatalf("Error registering service: %v", err)
	}

	viper.WatchConfig()
}
