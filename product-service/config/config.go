package config

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	"github.com/hashicorp/consul/api"
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

	JaegerTelemetryHost string
	JaegerTelemetryPort string

	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string

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

	QueueEventCreated string

	CommonSagaStatusPending string
	CommonSagaStatusSuccess string
	CommonSagaStatusFailed  string

	PostgresHost         string
	PostgresPort         string
	PostgresUsername     string
	PostgresPassword     string
	PostgresDatabaseName string
	PostgresSSLMode      string

	ElasticsearchHost     string
	ElasticsearchPort     string
	ElasticsearchUsername string
	ElasticsearchPassword string

	// USER SERVICE
	UserServiceServiceName string
	UserServiceRpcHost     string
	UserServiceRpcPort     string
	UserServiceHttpHost    string
	UserServiceHttpPort    string

	// USER SERVICE
	ProductServiceServiceName string
	ProductServiceRpcHost     string
	ProductServiceRpcPort     string
	ProductServiceHttpHost    string
	ProductServiceHttpPort    string
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
	kv := consulClient.KV()
	c.initTelemetry(kv)
	c.initCommon(kv)
	c.initRabbitmq(kv)
	c.initPostgres(kv)
	c.initUserService(kv)
	c.initProductService(kv)
	c.initElasticsearch(kv)
	c.initExchange(kv)
	c.initQueueProduct(kv)

	if err = c.RegisterConsulService(); err != nil {
		log.Fatalf("SetConfig | could not register service: %v", err)
		return
	}

	viper.WatchConfig()
}
