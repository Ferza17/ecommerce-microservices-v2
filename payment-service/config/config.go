package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"log"
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

	QueueEventCreated                 string
	QueuePaymentOrderCreated          string
	QueuePaymentOrderDelayedCancelled string

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
	kv := consulClient.KV()
	wg := sync.WaitGroup{}

	// Telemetry Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initTelemetry(kv)
	}()

	// RabbitMQ Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initRabbitmq(kv)
		c.initRabbitmqExchange(kv)
		c.initRabbitmqQueue(kv)
	}()

	// COMMON Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initCommon(kv)
	}()

	// Postgres Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initDatabasePostgres(kv)
	}()

	// Payment Service Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.initPaymentService(kv)
	}()

	wg.Wait()

	port, err := strconv.ParseInt(c.RpcPort, 10, 64)
	if err != nil {
		log.Fatalf("SetConfig | could not parse PORT to int: %v", err)
	}
	if err = consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    c.ServiceName,
		Address: c.RpcHost,
		Port:    int(port),
		Tags:    []string{"v1"},
	}); err != nil {
		log.Fatalf("Error registering service: %v", err)
	}

	viper.WatchConfig()
}
