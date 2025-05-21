package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
	"sync"

	"github.com/spf13/viper"
)

var c *Config

func Get() *Config {
	return c
}

type Config struct {
	// From ENV
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

	ProductServiceURL  string
	ProductServiceName string

	UserServiceName string
	UserServiceURL  string

	HttpHost string
	HttpPort string
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
	wg := &sync.WaitGroup{}

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

	// Product Service
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/services/product/RPC_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get PRODUCT RPC_HOST host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | PRODUCT RPC_HOST host is required")
		}
		productServiceRpcHost := string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/services/product/RPC_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get PRODUCT RPC_PORT host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | PRODUCT RPC_PORT host is required")
		}
		productServiceRpcPort := string(pair.Value)
		c.ProductServiceURL = fmt.Sprintf("%s:%s", productServiceRpcHost, productServiceRpcPort)

		pair, _, err = kv.Get(fmt.Sprintf("%s/services/product/SERVICE_NAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get PRODUCT SERVICE_NAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | PRODUCT SERVICE_NAME host is required")
		}
		c.ProductServiceName = string(pair.Value)
	}()

	// User Service
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/services/user/RPC_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get USER RPC_HOST host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | USER RPC_HOST host is required")
		}
		userServiceRpcHost := string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/RPC_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get USER RPC_PORT host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | USER RPC_PORT host is required")
		}
		userServiceRpcPort := string(pair.Value)
		c.UserServiceURL = fmt.Sprintf("%s:%s", userServiceRpcHost, userServiceRpcPort)

		pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/SERVICE_NAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get USER SERVICE_NAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | USER SERVICE_NAME host is required")
		}
		c.UserServiceName = string(pair.Value)
	}()

	// API GATEWAY
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/services/api-gateway/SERVICE_NAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get API-GATEWAY SERVICE_NAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | USER SERVICE_NAME host is required")
		}
		c.ServiceName = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/services/api-gateway/HTTP_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get API-GATEWAY HTTP_HOST host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | USER HTTP_HOST host is required")
		}
		c.HttpHost = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/services/api-gateway/HTTP_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get API-GATEWAY HTTP_PORT host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | USER HTTP_PORT host is required")
		}
		c.HttpPort = string(pair.Value)
	}()

	wg.Wait()

	httpPortInt, err := strconv.ParseInt(c.HttpPort, 10, 64)
	if err != nil {
		log.Fatalf("SetConfig | could not parse HTTP_PORT to int: %v", err)
	}
	if err = consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    c.ServiceName,
		Address: c.HttpHost,
		Port:    int(httpPortInt),
		Tags:    []string{"v1"},
	}); err != nil {
		log.Fatalf("Error registering service: %v", err)
	}
	viper.WatchConfig()
}
