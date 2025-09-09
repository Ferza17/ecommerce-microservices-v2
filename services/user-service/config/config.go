package config

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/metric"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"github.com/xhit/go-str2duration/v2"
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

	// Queue User
	QueueUserCreated string
	QueueUserUpdated string
	QueueUserLogin   string
	QueueUserLogout  string

	QueueEventCreated string

	// Queue Notification
	QueueNotificationEmailOtpCreated          string
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

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	JwtAccessTokenSecret          string
	JwtAccessTokenExpirationTime  time.Duration
	JwtRefreshTokenSecret         string
	JwtRefreshTokenExpirationTime time.Duration
	VerificationUserLoginUrl      string
	OtpExpirationTime             time.Duration

	// USER SERVICE
	UserServiceServiceName    string
	UserServiceRpcHost        string
	UserServiceRpcPort        string
	UserServiceHttpHost       string
	UserServiceHttpPort       string
	UserServiceMetricHttpPort string

	// EVENT STORE SERVICE
	EventStoreServiceRabbitMQ *ServiceEventStoreRabbitMQ
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

	// Get Consul Key / ValueconsulClient.KV()
	c.initTelemetry(consulClient.KV())
	c.initCommon(consulClient.KV())
	c.initRabbitmq(consulClient.KV())
	c.initPostgres(consulClient.KV())
	c.initRedis(consulClient.KV())
	c.initUserService(consulClient.KV())
	c.initExchange(consulClient.KV())
	c.initQueueProduct(consulClient.KV())
	c.initQueueUser(consulClient.KV())
	c.initQueueNotification(consulClient.KV())

	c.withServiceEventStoreRabbitMQ(consulClient.KV())

	// User Service Config
	pair, _, err := consulClient.KV().Get(fmt.Sprintf("%s/services/notification/SERVICE_NAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get notification/SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.NotificationServiceName = string(pair.Value)

	// Access Token Config
	pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/services/user/JWT_ACCESS_TOKEN_SECRET", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get JWT_ACCESS_TOKEN_SECRET from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | JWT_ACCESS_TOKEN_SECRET is required")
	}
	c.JwtAccessTokenSecret = string(pair.Value)
	pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/services/user/JWT_ACCESS_TOKEN_EXPIRATION_TIME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get JWT_ACCESS_TOKEN_EXPIRATION_TIME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | JWT_ACCESS_TOKEN_EXPIRATION_TIME is required")
	}
	c.JwtAccessTokenExpirationTime, err = str2duration.ParseDuration(string(pair.Value))
	if err != nil {
		log.Fatalf("SetConfig | JWT_ACCESS_TOKEN_EXPIRATION_TIME is invalid")
	}

	// Refresh Token Token Config
	pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/services/user/JWT_REFRESH_TOKEN_SECRET", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get JWT_REFRESH_TOKEN_SECRET from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | JWT_REFRESH_TOKEN_SECRET is required")
	}
	c.JwtRefreshTokenSecret = string(pair.Value)
	pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/services/user/JWT_REFRESH_TOKEN_EXPIRATION_TIME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get JWT_REFRESH_TOKEN_EXPIRATION_TIME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | JWT_REFRESH_TOKEN_EXPIRATION_TIME is required")
	}
	c.JwtRefreshTokenExpirationTime, err = str2duration.ParseDuration(string(pair.Value))
	if err != nil {
		log.Fatalf("SetConfig | JWT_REFRESH_TOKEN_EXPIRATION_TIME is invalid")
	}

	// Verification User Login Url Config
	pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/services/user/VERIFICATION_USER_LOGIN_URL", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get VERIFICATION_USER_LOGIN_URL from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | VERIFICATION_USER_LOGIN_URL is required")
	}
	c.VerificationUserLoginUrl = string(pair.Value)

	// OTP Expiration Time
	pair, _, err = consulClient.KV().Get(fmt.Sprintf("%s/services/user/OTP_EXPIRATION_TIME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get  from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul |  is required")
	}
	c.OtpExpirationTime, err = str2duration.ParseDuration(string(pair.Value))
	if err != nil {
		log.Fatalf("SetConfig |  is invalid")
	}

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

func (c *Config) withServiceEventStoreRabbitMQ(kv *api.KV) *Config {
	c.EventStoreServiceRabbitMQ = DefaultServiceEventStoreRabbitMQ().WithConsulClient(c.Env, kv)
	return c
}
