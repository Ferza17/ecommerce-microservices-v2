package config

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/metric"
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
	NotificationServiceName string

	// EXCHANGE
	ExchangeCommerce       string
	ExchangeEvent          string
	ExchangeNotification   string
	ExchangeProduct        string
	ExchangeUser           string
	ExchangePaymentDelayed string
	ExchangePaymentDirect  string

	// Queue Product
	QueueProductCreated string
	QueueProductUpdated string
	QueueProductDeleted string

	// Queue User
	QueueUserCreated string
	QueueUserUpdated string
	QueueUserLogin   string
	QueueUserLogout  string

	QueueEventCreated string

	// Queue Notification
	QueueNotificationEmailOtpCreated          string
	QueueNotificationEmailPaymentOrderCreated string

	BrokerKafka                         *BrokerKafka
	BrokerKafkaTopic                    *BrokerKafkaTopic
	BrokerKafkaTopicConnectorSinkPgUser *BrokerKafkaTopicConnectorSinkPgUser

	ConfigTelemetry *ConfigTelemetry

	BrokerRabbitMQ            *BrokerRabbitMQ
	EventStoreServiceRabbitMQ *ServiceEventStoreRabbitMQ

	DatabasePostgres *DatabasePostgres
	DatabaseRedis    *DatabaseRedis

	//  SERVICE
	ConfigServiceUser *ConfigServiceUser
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

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("config not found: %s", err.Error()))
	}
	if err = viper.Unmarshal(&c); err != nil {
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
	c.initExchange(consulClient.KV())
	c.initQueueProduct(consulClient.KV())
	c.initQueueUser(consulClient.KV())
	c.initQueueNotification(consulClient.KV())

	c.
		withServiceUser(consulClient.KV()).
		withConfigTelemetry(consulClient.KV()).
		withServiceEventStoreRabbitMQ(consulClient.KV()).
		withDatabasePostgres(consulClient.KV()).
		withDatabaseRedis(consulClient.KV()).
		withBrokerRabbitMQ(consulClient.KV()).
		withBrokerKafka(consulClient.KV()).
		withBrokerKafkaTopic(consulClient.KV()).
		withBrokerKafkaTopicConnectorSinkPgUser(consulClient.KV())

	// User Service Config
	pair, _, err := consulClient.KV().Get(fmt.Sprintf("%s/services/notification/SERVICE_NAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get notification/SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.NotificationServiceName = string(pair.Value)

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

func (c *Config) withServiceUser(kv *api.KV) *Config {
	c.ConfigServiceUser = DefaultConfigServiceUser().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withServiceEventStoreRabbitMQ(kv *api.KV) *Config {
	c.EventStoreServiceRabbitMQ = DefaultServiceEventStoreRabbitMQ().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withDatabasePostgres(kv *api.KV) *Config {
	c.DatabasePostgres = DefaultDatabasePostgres().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withDatabaseRedis(kv *api.KV) *Config {
	c.DatabaseRedis = DefaultDatabaseRedis().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withBrokerRabbitMQ(kv *api.KV) *Config {
	c.BrokerRabbitMQ = DefaultRabbitMQBroker().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withBrokerKafka(kv *api.KV) *Config {
	c.BrokerKafka = DefaultKafkaBroker().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withBrokerKafkaTopic(kv *api.KV) *Config {
	c.BrokerKafkaTopic = DefaultKafkaBrokerTopic().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withBrokerKafkaTopicConnectorSinkPgUser(kv *api.KV) *Config {
	c.BrokerKafkaTopicConnectorSinkPgUser = DefaultBrokerKafkaTopicsConnectorSinkPgUser().WithConsulClient(c.Env, kv)
	return c
}

func (c *Config) withConfigTelemetry(kv *api.KV) *Config {
	c.ConfigTelemetry = DefaultConfigTelemetry().WithConsulClient(c.Env, kv)
	return c
}
