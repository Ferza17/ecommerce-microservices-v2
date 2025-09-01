package config

import (
	"fmt"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/metric"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"log"
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
	EventStoreServiceServiceName    string
	EventStoreServiceRpcHost        string
	EventStoreServiceRpcPort        string
	EventStoreServiceHttpHost       string
	EventStoreServiceHttpPort       string
	EventStoreServiceMetricHttpPort string

	JaegerTelemetryHost string
	JaegerTelemetryPort string

	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string

	ExchangeEvent string

	QueueEventCreated string

	MongoUsername     string
	MongoPassword     string
	MongoHost         string
	MongoPort         string
	MongoDatabaseName string
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

	c.initTelemetry(consulClient.KV())
	c.initServiceEventStore(consulClient.KV())
	c.initRabbitmq(consulClient.KV())
	c.initDatabaseMongodb(consulClient.KV())

	if err = c.RegisterConsulService(); err != nil {
		log.Fatalf("SetConfig | could not register consul service: %v", err)
		return
	}

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
