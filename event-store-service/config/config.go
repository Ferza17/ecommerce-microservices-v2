package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/enum"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/metric"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
)

var c *Config = DefaultConfig()

func Get() *Config {
	return c
}

type Config struct {
	Env        string `mapstructure:"ENV"`
	ConsulHost string `mapstructure:"CONSUL_HOST"`
	ConsulPort string `mapstructure:"CONSUL_PORT"`

	// Service Config
	EventStoreService         *ServiceEventStore
	EventStoreServiceRabbitMQ *ServiceEventStoreRabbitMQ
	// Database MongoDB
	DatabaseMongodb *DatabaseMongodb
	// Telemetry
	ConfigTelemetry *ConfigTelemetry
	// MESSAGE BROKER RABBITMQ
	MessageBrokerRabbitMQ *MessageBrokerRabbitMQ
}

func DefaultConfig() *Config {
	return &Config{
		Env:                       "",
		ConsulHost:                "",
		ConsulPort:                "",
		EventStoreService:         DefaultServiceEventStore(),
		EventStoreServiceRabbitMQ: DefaultServiceEventStoreRabbitMQ(),
		DatabaseMongodb:           DefaultDatabaseMongodb(),
		ConfigTelemetry:           DefaultConfigTelemetry(),
		MessageBrokerRabbitMQ:     DefaultMessageBrokerRabbitMQ(),
	}
}

func SetConfig(path string) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	switch os.Getenv("ENV") {
	case enum.CONFIG_ENV_PROD:
		viper.SetConfigName(".env.production")
	default:
		viper.SetConfigName(".env.local")
	}

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

	c.
		withServiceEventStore(consulClient.KV()).
		withServiceEventStoreRabbitMQ(consulClient.KV()).
		withDatabaseMongoDB(consulClient.KV()).
		withConfigTelemetry(consulClient.KV()).
		withMessageBrokerRabbitMQ(consulClient.KV()).
		withRegisterConsulService()

	// Register Prometheus
	prometheus.MustRegister(
		pkgMetric.GrpcRequestsTotal,
		pkgMetric.GrpcRequestDuration,
		pkgMetric.HttpRequestsTotal,
		pkgMetric.HttpRequestDuration,
		pkgMetric.RabbitmqMessagesPublished,
		pkgMetric.RabbitmqMessagesConsumed,
	)

	viper.WatchConfig()
}

func (c *Config) withServiceEventStore(kv *api.KV) *Config {
	c.EventStoreService = DefaultServiceEventStore().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withServiceEventStoreRabbitMQ(kv *api.KV) *Config {
	c.EventStoreServiceRabbitMQ = DefaultServiceEventStoreRabbitMQ().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withDatabaseMongoDB(kv *api.KV) *Config {
	c.DatabaseMongodb = DefaultDatabaseMongodb().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withConfigTelemetry(kv *api.KV) *Config {
	c.ConfigTelemetry = DefaultConfigTelemetry().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withMessageBrokerRabbitMQ(kv *api.KV) *Config {
	c.MessageBrokerRabbitMQ = DefaultMessageBrokerRabbitMQ().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withRegisterConsulService() *Config {
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%s:%s", c.ConsulHost, c.ConsulPort),
	})
	if err != nil {
		log.Fatalf("SetConfig | could not connect to consul: %v", err)
	}

	port, err := strconv.ParseInt(c.EventStoreService.RpcPort, 10, 64)
	if err != nil {
		log.Fatalf("SetConfig | could not parse PORT to int: %v", err)
	}
	if err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Kind:    api.ServiceKindTypical,
		Name:    c.EventStoreService.ServiceName,
		Address: c.EventStoreService.RpcHost,
		Port:    int(port),
		Tags:    []string{"service", "rabbitmq-client"},
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%s", c.EventStoreService.RpcHost, c.EventStoreService.RpcPort),
			GRPCUseTLS:                     false,
			Interval:                       "30s", // Less frequent checks
			Timeout:                        "5s",  // Reasonable timeout
			DeregisterCriticalServiceAfter: "40s", // Give more time before deregistering
		},
		Connect: &api.AgentServiceConnect{
			Native: true,
		},
	}); err != nil {
		panic(fmt.Sprintf("Error registering service: %v", err))
	}
	return nil
}
