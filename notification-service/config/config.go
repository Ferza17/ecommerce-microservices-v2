package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
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

	QueueNotificationEmailOtpCreated          string
	QueueNotificationEmailPaymentOrderCreated string
	QueueEventCreated                         string

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

	RpcHost string
	RpcPort string
}

func SetConfig(path string) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	switch os.Getenv("ENV") {
	case enum.CONFIG_ENV_LOCAL:
		viper.SetConfigName(".env.local")
	case enum.CONFIG_ENV_PROD:
		viper.SetConfigName(".env.production")
	default:
		log.Fatal("SetConfig | env is required")
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

		// EXCHANGE
		pair, _, err := consulClient.KV().Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/EVENT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get EXCHANGE/EVENT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | EXCHANGE/EVENT is required")
		}
		c.ExchangeEvent = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/NOTIFICATION", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get EXCHANGE/NOTIFICATION from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | EXCHANGE/NOTIFICATION is required")
		}
		c.ExchangeNotification = string(pair.Value)

		// QUEUE
		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get QUEUE/NOTIFICATION/EMAIL/OTP/CREATED host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | QUEUE/NOTIFICATION/EMAIL/OTP/CREATED host is required")
		}
		c.QueueNotificationEmailOtpCreated = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED host is required")
		}
		c.QueueNotificationEmailPaymentOrderCreated = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/EVENT/CREATED", c.Env), nil)
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
		c.initCommon(consulClient.KV())
	}()

	// SMTP Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		pair, _, err := consulClient.KV().Get(fmt.Sprintf("%s/smtp/SMTP_SENDER_EMAIL", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SMTP_SENDER_EMAIL host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SMTP_SENDER_EMAIL host is required")
		}
		c.SmtpSenderEmail = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/smtp/SMTP_USERNAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SMTP_USERNAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SMTP_USERNAME host is required")
		}
		c.SmtpUsername = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/smtp/SMTP_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SMTP_PASSWORD host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SMTP_PASSWORD host is required")
		}
		c.SmtpPassword = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/smtp/SMTP_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SMTP_HOST host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SMTP_HOST host is required")
		}
		c.SmtpHost = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/smtp/SMTP_PORT", c.Env), nil)
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
		pair, _, err := consulClient.KV().Get(fmt.Sprintf("%s/database/mongodb/MONGO_USERNAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_USERNAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_USERNAME host is required")
		}
		c.MongoUsername = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/database/mongodb/MONGO_PASSWORD", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_PASSWORD host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_PASSWORD host is required")
		}
		c.MongoPassword = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/database/mongodb/MONGO_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_HOST host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_HOST host is required")
		}
		c.MongoHost = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/database/mongodb/MONGO_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get MONGO_PORT host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | MONGO_PORT host is required")
		}
		c.MongoPort = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/database/mongodb/MONGO_DATABASE_NAME/NOTIFICATION", c.Env), nil)
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
		pair, _, err := consulClient.KV().Get(fmt.Sprintf("%s/services/notification/SERVICE_NAME", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get SERVICE_NAME host from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | SERVICE_NAME host is required")
		}
		c.ServiceName = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/services/notification/RPC_HOST", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RPC_HOST is required")
		}
		c.RpcHost = string(pair.Value)

		pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/services/notification/RPC_PORT", c.Env), nil)
		if err != nil {
			log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
		}
		if pair == nil {
			log.Fatal("SetConfig | Consul | RPC_PORT is required")
		}
		c.RpcPort = string(pair.Value)

	}()

	wg.Wait()

	viper.WatchConfig()
}
