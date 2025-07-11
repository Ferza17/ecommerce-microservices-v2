package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
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

	ExchangeEvent string

	QueueEventCreated string

	MongoUsername     string
	MongoPassword     string
	MongoHost         string
	MongoPort         string
	MongoDatabaseName string
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
		c.initRabbitmq(kv)

		// EXCHANGE
		pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/EVENT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get EXCHANGE/EVENT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | EXCHANGE/EVENT is required")
		}
		c.ExchangeEvent = string(pair.Value)

		// QUEUE
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/EVENT/CREATED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get QUEUE/EVENT/CREATED host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | QUEUE/EVENT/CREATED host is required")
		}
		c.QueueEventCreated = string(pair.Value)
	}()

	// MongoDB Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_USERNAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_USERNAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_USERNAME host is required")
		}
		c.MongoUsername = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_PASSWORD host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_PASSWORD host is required")
		}
		c.MongoPassword = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_HOST host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_HOST host is required")
		}
		c.MongoHost = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_PORT host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_PORT host is required")
		}
		c.MongoPort = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_DATABASE_NAME/EVENT_STORE host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_DATABASE_NAME/EVENT_STORE host is required")
		}
		c.MongoDatabaseName = string(pair.Value)
	}()

	// Event Store Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/services/event-store/SERVICE_NAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SERVICE_NAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SERVICE_NAME host is required")
		}
		c.ServiceName = string(pair.Value)
	}()

	wg.Wait()

	if err = consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name: c.ServiceName,
		Tags: []string{"v1"},
	}); err != nil {
		log.Fatalf("Error registering service: %v", err)
	}
	viper.WatchConfig()
}
