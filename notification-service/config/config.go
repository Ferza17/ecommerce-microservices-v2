package config

import (
	"fmt"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/metric"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
	"sync"

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
	NotificationServiceServiceName    string
	NotificationServiceRpcHost        string
	NotificationServiceRpcPort        string
	NotificationServiceHttpHost       string
	NotificationServiceHttpPort       string
	NotificationServiceMetricHttpPort string

	JaegerTelemetryHost string
	JaegerTelemetryPort string

	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string

	QueueEventCreated string

	// EXCHANGE
	ExchangeCommerce       string
	ExchangeEvent          string
	ExchangeNotification   string
	ExchangeProduct        string
	ExchangeUser           string
	ExchangePaymentDelayed string
	ExchangePaymentDirect  string

	// Queue Notification
	QueueNotificationEmailOtpCreated          string
	QueueNotificationEmailPaymentOrderCreated string

	CommonSagaStatusPending string
	CommonSagaStatusSuccess string
	CommonSagaStatusFailed  string

	SmtpSenderEmail string
	SmtpHost        string
	SmtpPort        string
	SmtpUsername    string
	SmtpPassword    string

	MongoUsername     string
	MongoPassword     string
	MongoHost         string
	MongoPort         string
	MongoDatabaseName string
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
	wg := new(sync.WaitGroup)

	// Telemetry Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initTelemetry(consulClient.KV())
	}()

	// RabbitMQ Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initRabbitmq(consulClient.KV())
		c.initExchange(consulClient.KV())
		c.initQueueNotification(consulClient.KV())
	}()

	// COMMON Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initCommon(consulClient.KV())
	}()

	// SMTP Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initSmtp(consulClient.KV())
	}()

	// MongoDB Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initMongoDB(consulClient.KV())
	}()

	// Notification Service Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initNotificationService(consulClient.KV())
	}()

	wg.Wait()

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
