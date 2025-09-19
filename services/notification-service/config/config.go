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

	c.ConfigServiceNotification = DefaultConfigServiceNotification().WithConsulClient(c.Env, consulClient.KV())
	c.ConfigServicePayment = DefaultConfigServicePayment().WithConsulClient(c.Env, consulClient.KV())
	c.ConfigServiceShipping = DefaultConfigServiceShipping().WithConsulClient(c.Env, consulClient.KV())
	c.ConfigServiceUser = DefaultConfigServiceUser().WithConsulClient(c.Env, consulClient.KV())
	c.ConfigSmtp = DefaultConfigSmtp().WithConsulClient(c.Env, consulClient.KV())
	c.ConfigTelemetry = DefaultConfigTelemetry().WithConsulClient(c.Env, consulClient.KV())
	c.DatabaseMongo = DefaultDatabaseMongo().WithConsulClient(c.Env, consulClient.KV())
	c.BrokerKafka = DefaultKafkaBroker().WithConsulClient(c.Env, consulClient.KV())
	c.BrokerKafkaTopicNotifications = DefaultKafkaBrokerTopicNotifications().WithConsulClient(c.Env, consulClient.KV())

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
