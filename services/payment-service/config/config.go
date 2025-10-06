package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/metric"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
)

var c *Config

func Get() *Config {
	return c
}

type Config struct {
	Env        string `mapstructure:"ENV"`
	ConsulHost string `mapstructure:"CONSUL_HOST"`
	ConsulPort string `mapstructure:"CONSUL_PORT"`

	ConfigTelemetry *ConfigTelemetry

	// POSTGRES CONFIG
	DatabasePostgres *DatabasePostgres
	// REDIS Config
	DatabaseRedis *DatabaseRedis
	DatabaseMongo *DatabaseMongo

	// User Service Config
	ConfigServiceUser *ConfigServiceUser

	// Product Service Config
	ConfigServiceProduct *ConfigServiceProduct

	// Payment Service Config
	ConfigServicePayment *ConfigServicePayment

	// Shipping Service Config
	ConfigServiceShipping *ConfigServiceShipping

	BrokerKafka                             *BrokerKafka
	BrokerKafkaTopicConnectorSinkPgPayment  *BrokerKafkaTopicConnectorSinkPgPayment
	BrokerKafkaTopicPayments                *BrokerKafkaTopicPayments
	BrokerKafkaTopicNotifications           *BrokerKafkaTopicNotifications
	BrokerKafkaTopicShippings               *BrokerKafkaTopicShippings
	BrokerKafkaTopicConnectorSinkMongoEvent *BrokerKafkaTopicConnectorSinkMongoEvent
}

func SetConfig(path string) {
	c = &Config{}
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	switch os.Getenv("ENV") {
	case enum.CONFIG_ENV_PROD:
		viper.SetConfigName(".env.production")
	default:
		viper.SetConfigName(".env.local")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("config not found: %s", err.Error()))
	}
	if err = viper.Unmarshal(&c); err != nil {
		log.Fatalf("SetConfig | could not parse config: %v", err)
	}

	if c.Env == "" {
		log.Fatal("SetConfig | ENV is required")
	}
	if c.ConsulHost == "" {
		log.Fatal("SetConfig | CONSUL_HOST host is required")
	}
	if c.ConsulPort == "" {
		log.Fatal("SetConfig | CONSUL_PORT port is required")
	}

	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%s:%s", c.ConsulHost, c.ConsulPort),
	})
	if err != nil {
		log.Fatalf("SetConfig | could not connect to consul: %v", err)
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

	if err = c.
		withBrokerKafka(client.KV()).
		withBrokerKafkaTopicConnectorSinkPgPayment(client.KV()).
		withBrokerKafkaTopicNotifications(client.KV()).
		withBrokerKafkaTopicShippings(client.KV()).
		withBrokerKafkaTopicPayments(client.KV()).
		withConfigTelemetry(client.KV()).
		withConfigServicePayment(client.KV()).
		withServiceShipping(client.KV()).
		withServiceUser(client.KV()).
		withServiceProduct(client.KV()).
		withDatabasePostgres(client.KV()).
		withDatabaseRedis(client.KV()).
		withConfigDatabaseMongo(client.KV()).
		withBrokerKafkaTopicConnectorSinkMongoEvent(client.KV()).
		RegisterConsulService(); err != nil {
		log.Fatalf("SetConfig | could not register service: %v", err)
		return
	}

	viper.WatchConfig()
}
