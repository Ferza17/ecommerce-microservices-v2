package config

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/metric"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
	"sync"
)

var c *Config

func Get() *Config {
	return c
}

type Config struct {
	Env        string `mapstructure:"ENV"`
	ConsulHost string `mapstructure:"CONSUL_HOST"`
	ConsulPort string `mapstructure:"CONSUL_PORT"`

	// From CONSUL; Local Service Config
	PaymentOrderCancelledInMs int

	// COMMON Config
	CommonSagaStatusPending string
	CommonSagaStatusSuccess string
	CommonSagaStatusFailed  string

	// JAEGER TELEMETRY Config
	JaegerTelemetryHost string
	JaegerTelemetryPort string

	// RABBITMQ Config
	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string

	// EXCHANGE Config
	ExchangeCommerce       string
	ExchangeEvent          string
	ExchangeNotification   string
	ExchangeProduct        string
	ExchangeUser           string
	ExchangePaymentDelayed string
	ExchangePaymentDirect  string
	ExchangeShipping       string

	// Queue Notification Config
	QueueNotificationEmailOtpCreated          string
	QueueNotificationEmailPaymentOrderCreated string

	// Queue Payment Config
	QueuePaymentOrderCreated          string
	QueuePaymentOrderDelayedCancelled string

	// Queue Shipping Config
	QueueShippingCreated string
	QueueShippingUpdated string

	// POSTGRES CONFIG
	PostgresHost         string
	PostgresPort         string
	PostgresUsername     string
	PostgresPassword     string
	PostgresDatabaseName string
	PostgresSSLMode      string

	// REDIS Config
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	// User Service Config
	UserServiceServiceName    string
	UserServiceRpcHost        string
	UserServiceRpcPort        string
	UserServiceHttpHost       string
	UserServiceHttpPort       string
	UserServiceMetricHttpPort string

	// Product Service Config
	ProductServiceServiceName    string
	ProductServiceRpcHost        string
	ProductServiceRpcPort        string
	ProductServiceHttpHost       string
	ProductServiceHttpPort       string
	ProductServiceMetricHttpPort string

	// Payment Service Config
	PaymentServiceServiceName    string
	PaymentServiceRpcHost        string
	PaymentServiceRpcPort        string
	PaymentServiceHttpHost       string
	PaymentServiceHttpPort       string
	PaymentServiceMetricHttpPort string

	// Shipping Service Config
	ShippingServiceServiceName    string
	ShippingServiceRpcHost        string
	ShippingServiceRpcPort        string
	ShippingServiceHttpHost       string
	ShippingServiceHttpPort       string
	ShippingServiceMetricHttpPort string
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
	if err := viper.Unmarshal(&c); err != nil {
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

	// Local Config
	pair, _, err := client.KV().Get(fmt.Sprintf("%s/services/payment/PAYMENT_ORDER_CANCELLED_IN_MS", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get  from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul |  is required")
	}
	temp, err := strconv.ParseInt(string(pair.Value), 10, 64)
	if err != nil {
		log.Fatalf("SetConfig | could not parse PAYMENT_ORDER_CANCELLED_IN_MS to int: %v", err)
	}
	c.PaymentOrderCancelledInMs = int(temp)
	// End Local Config

	// Get Consul Key / Value
	wg := sync.WaitGroup{}

	// Telemetry Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initTelemetry(client.KV())
	}()

	// RabbitMQ Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initRabbitmq(client.KV())
		c.initExchange(client.KV())
		c.initQueueNotification(client.KV())
		c.initQueuePayment(client.KV())
		c.initQueueShipping(client.KV())
	}()

	// COMMON Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initCommon(client.KV())
	}()

	// Database Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initPostgres(client.KV())
		c.initRedis(client.KV())
	}()

	// Service Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initPaymentService(client.KV())
		c.initUserService(client.KV())
		c.initProductService(client.KV())
		c.initShippingService(client.KV())
	}()

	wg.Wait()

	// Register Prometheus
	prometheus.MustRegister(
		pkgMetric.GrpcRequestsTotal,
		pkgMetric.GrpcRequestDuration,
		pkgMetric.HttpRequestsTotal,
		pkgMetric.HttpRequestDuration,
		pkgMetric.RabbitmqMessagesPublished,
		pkgMetric.RabbitmqMessagesConsumed,
	)

	if err = c.RegisterConsulService(); err != nil {
		log.Fatalf("SetConfig | could not register service: %v", err)
		return
	}

	viper.WatchConfig()
}
