package config

import (
	"fmt"
	"log"
	"os"

	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/metric"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
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

	// Notification SERVICE
	ConfigServiceNotification *ConfigServiceNotification
	ConfigServicePayment      *ConfigServicePayment
	ConfigServiceUser         *ConfigServiceUser
	ConfigServiceShipping     *ConfigServiceShipping

	ConfigTelemetry *ConfigTelemetry

	BrokerKafka                   *BrokerKafka
	BrokerKafkaTopicNotifications *BrokerKafkaTopicNotifications
	BrokerKafkaTopicUsers         *BrokerKafkaTopicUsers

	ConfigSmtp *ConfigSmtp

	DatabaseMongo *DatabaseMongo
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

	if err = c.
		withBrokerKafka(consulClient.KV()).
		withBrokerKafkaTopicNotifications(consulClient.KV()).
		withBrokerKafkaTopicUsers(consulClient.KV()).
		withConfigTelemetry(consulClient.KV()).
		withConfigDatabaseMongo(consulClient.KV()).
		withConfigServiceNotification(consulClient.KV()).
		withConfigServicePayment(consulClient.KV()).
		withConfigServiceShipping(consulClient.KV()).
		withConfigServiceUser(consulClient.KV()).
		withConfigSmtp(consulClient.KV()).
		RegisterConsulService(); err != nil {
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
