package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
	"sync"

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

	// From Consul
	ServiceName string

	JaegerTelemetryHost string
	JaegerTelemetryPort string

	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string

	ExchangeEvent   string
	ExchangeProduct string

	QueueEventCreated   string
	QueueProductCreated string
	QueueProductDeleted string
	QueueProductUpdated string

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

	RpcHost string
	RpcPort string
}

func SetConfig(path string) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")

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
	wg := sync.WaitGroup{}

	// Telemetry Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/telemetry/jaeger/JAEGER_TELEMETRY_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get telemetry host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | telemetry host is required")
		}
		c.JaegerTelemetryHost = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/telemetry/jaeger/JAEGER_TELEMETRY_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get telemetry host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | telemetry host is required")
		}
		c.JaegerTelemetryPort = string(pair.Value)
	}()

	// RabbitMQ Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_USERNAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RABBITMQ_USERNAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RABBITMQ_USERNAME host is required")
		}
		c.RabbitMQUsername = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RABBITMQ_PASSWORD host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RABBITMQ_PASSWORD host is required")
		}
		c.RabbitMQPassword = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RABBITMQ_HOST host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RABBITMQ_HOST host is required")
		}
		c.RabbitMQHost = string(pair.Value)
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RABBITMQ_PORT host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RABBITMQ_PORT host is required")
		}
		c.RabbitMQPort = string(pair.Value)

		// EXCHANGE
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/EVENT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get EXCHANGE/EVENT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | EXCHANGE/EVENT is required")
		}
		c.ExchangeEvent = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/PRODUCT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get EXCHANGE/PRODUCT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | EXCHANGE/PRODUCT is required")
		}
		c.ExchangeProduct = string(pair.Value)

		// QUEUE
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/EVENT/CREATED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get QUEUE/EVENT/CREATED host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | QUEUE/EVENT/CREATED host is required")
		}
		c.QueueEventCreated = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/PRODUCT/CREATED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get QUEUE/PRODUCT/CREATED host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | QUEUE/PRODUCT/CREATED host is required")
		}
		c.QueueProductCreated = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/PRODUCT/DELETED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get QUEUE/PRODUCT/DELETED host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | QUEUE/PRODUCT/DELETED host is required")
		}
		c.QueueProductDeleted = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/PRODUCT/UPDATED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get QUEUE/PRODUCT/UPDATED host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | QUEUE/PRODUCT/UPDATED host is required")
		}
		c.QueueProductUpdated = string(pair.Value)

	}()

	// COMMON Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/common/SAGA_STATUS/PENDING", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SAGA_STATUS/PENDING from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SAGA_STATUS/PENDING host is required")
		}
		c.CommonSagaStatusPending = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/common/SAGA_STATUS/SUCCESS", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SAGA_STATUS/SUCCESS from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SAGA_STATUS/SUCCESS host is required")
		}
		c.CommonSagaStatusSuccess = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/common/SAGA_STATUS/FAILED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SAGA_STATUS/FAILED from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SAGA_STATUS/FAILED host is required")
		}
		c.CommonSagaStatusFailed = string(pair.Value)
	}()

	// Postgres Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_USERNAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_USERNAME from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_USERNAME is required")
		}
		c.PostgresUsername = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_PASSWORD from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_PASSWORD is required")
		}
		c.PostgresPassword = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_SSL_MODE", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_SSL_MODE from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_SSL_MODE is required")
		}
		c.PostgresSSLMode = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_HOST from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_HOST is required")
		}
		c.PostgresHost = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_PORT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_PORT is required")
		}
		c.PostgresPort = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_DATABASE_NAME/PRODUCTS", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get POSTGRES_DATABASE_NAME/PRODUCTS from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | POSTGRES_DATABASE_NAME/PRODUCTS is required")
		}
		c.PostgresDatabaseName = string(pair.Value)
	}()

	// Elasticsearch Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/database/elasticsearch/ELASTICSEARCH_USERNAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get ELASTICSEARCH_USERNAME from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | ELASTICSEARCH_USERNAME is required")
		}
		c.ElasticsearchUsername = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/elasticsearch/ELASTICSEARCH_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get ELASTICSEARCH_PASSWORD from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | ELASTICSEARCH_PASSWORD is required")
		}
		c.ElasticsearchPassword = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/elasticsearch/ELASTICSEARCH_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get ELASTICSEARCH_HOST from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | ELASTICSEARCH_HOST is required")
		}
		c.ElasticsearchHost = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/elasticsearch/ELASTICSEARCH_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get ELASTICSEARCH_PORT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | ELASTICSEARCH_PORT is required")
		}
		c.ElasticsearchPort = string(pair.Value)
	}()

	// Product Service Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/services/product/SERVICE_NAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
		}
		c.ServiceName = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/services/product/RPC_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RPC_HOST is required")
		}
		c.RpcHost = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/services/product/RPC_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RPC_PORT is required")
		}
		c.RpcPort = string(pair.Value)
	}()

	wg.Wait()

	if err = consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name: c.ServiceName,
		Tags: []string{"v1"},
	}); err != nil {
		log.Fatalf("Error registering service: %v", err)
	}

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
