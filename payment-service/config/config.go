package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"log"
	"os"
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

	// From Consul
	ServiceName               string
	PaymentOrderCancelledInMs int

	NotificationServiceName string

	JaegerTelemetryHost string
	JaegerTelemetryPort string

	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string

	ExchangeEvent          string
	ExchangeNotification   string
	ExchangePaymentDirect  string
	ExchangePaymentDelayed string

	// EVENT QUEUE
	QueueEventCreated string

	// PAYMENT QUEUE
	QueuePaymentOrderCreated          string
	QueuePaymentOrderDelayedCancelled string

	// NOTIFICATION QUEUE
	QueueNotificationCreated                  string
	QueueNotificationEmailPaymentOrderCreated string

	CommonSagaStatusPending string
	CommonSagaStatusSuccess string
	CommonSagaStatusFailed  string

	PostgresHost         string
	PostgresPort         string
	PostgresUsername     string
	PostgresPassword     string
	PostgresDatabaseName string
	PostgresSSLMode      string

	RpcHost string
	RpcPort string
}

func SetConfig(path string) {
	c = &Config{}
	if c.Env = os.Getenv("ENV"); c.Env != "" {
		log.Println("Load Config from OS ENV")

		c.ConsulHost = os.Getenv("CONSUL_HOST")
		c.ConsulPort = os.Getenv("CONSUL_PORT")
	} else {
		log.Println("Load Config from .env")
		viper.SetConfigType("env")
		viper.AddConfigPath(path)
		viper.SetConfigName(".env")

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Sprintf("config not found: %s", err.Error()))
		}
		if err := viper.Unmarshal(&c); err != nil {
			log.Fatalf("SetConfig | could not parse config: %v", err)
		}
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
		c.initRabbitmqExchange(client.KV())
		c.initRabbitmqQueue(client.KV())
	}()

	// COMMON Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initCommon(client.KV())
	}()

	// Postgres Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initDatabasePostgres(client.KV())
	}()

	// Payment Service Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initPaymentService(client.KV())
	}()

	wg.Wait()
	viper.WatchConfig()
}
