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
	// From ENV
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

	ExchangeEvent        string
	ExchangeNotification string

	QueueNotificationEmailOtpCreated string
	QueueEventCreated                string

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
	wg := &sync.WaitGroup{}

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

		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/NOTIFICATION", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get EXCHANGE/NOTIFICATION from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | EXCHANGE/NOTIFICATION is required")
		}
		c.ExchangeNotification = string(pair.Value)

		// QUEUE
		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get QUEUE/NOTIFICATION/EMAIL/OTP/CREATED host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | QUEUE/NOTIFICATION/EMAIL/OTP/CREATED host is required")
		}
		c.QueueNotificationEmailOtpCreated = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/EVENT/CREATED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get QUEUE/EVENT/CREATED host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | QUEUE/EVENT/CREATED host is required")
		}
		c.QueueEventCreated = string(pair.Value)

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

	// SMTP Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/smtp/SMTP_SENDER_EMAIL", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SMTP_SENDER_EMAIL host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SMTP_SENDER_EMAIL host is required")
		}
		c.SmtpSenderEmail = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/smtp/SMTP_USERNAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SMTP_USERNAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SMTP_USERNAME host is required")
		}
		c.SmtpUsername = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/smtp/SMTP_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SMTP_PASSWORD host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SMTP_PASSWORD host is required")
		}
		c.SmtpPassword = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/smtp/SMTP_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SMTP_HOST host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SMTP_HOST host is required")
		}
		c.SmtpHost = string(pair.Value)

		pair, _, err = kv.Get(fmt.Sprintf("%s/smtp/SMTP_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SMTP_PORT host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SMTP_PORT host is required")
		}
		c.SmtpPort = string(pair.Value)
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

		pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_DATABASE_NAME/NOTIFICATION", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_DATABASE_NAME/NOTIFICATION host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_DATABASE_NAME/NOTIFICATION host is required")
		}
		c.MongoDatabaseName = string(pair.Value)
	}()

	// Notification Service Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := kv.Get(fmt.Sprintf("%s/services/notification/SERVICE_NAME", c.Env), nil)
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
