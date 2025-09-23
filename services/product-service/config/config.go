package config

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/metric"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"log"
	"os"
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

	ConfigTelemetry *ConfigTelemetry

	DatabasePostgres      *DatabasePostgres
	DatabaseElasticsearch *DatabaseElasticsearch

	// USER SERVICE
	ConfigServiceUser *ConfigServiceUser

	// PRODUCT SERVICE
	ConfigServiceProduct *ConfigServiceProduct

	BrokerKafka                          *BrokerKafka
	BrokerKafkaTopicProducts             *BrokerKafkaTopicProducts
	BrokerKafkaTopicConnectorSinkProduct *BrokerKafkaTopicConnectorSinkProduct
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

	// Get Consul Key / Value
	c.withServiceProduct(consulClient.KV())
	c.withServiceUser(consulClient.KV())
	c.withDatabaseElasticsearch(consulClient.KV())
	c.withDatabasePostgres(consulClient.KV())
	c.withConfigTelemetry(consulClient.KV())
	c.withBrokerKafka(consulClient.KV())
	c.withBrokerKafkaTopicConnectorSinkProduct(consulClient.KV())
	c.withBrokerKafkaTopicProducts(consulClient.KV())

	if err = c.RegisterConsulService(); err != nil {
		log.Fatalf("SetConfig | could not register service: %v", err)
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
