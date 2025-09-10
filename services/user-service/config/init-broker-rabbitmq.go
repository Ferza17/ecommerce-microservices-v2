package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type BrokerRabbitMQ struct {
	Username  string
	Password  string
	Host      string
	Port      string
	keyPrefix string
}

func DefaultRabbitMQBroker() *BrokerRabbitMQ {
	return &BrokerRabbitMQ{
		Username:  "",
		Password:  "",
		Host:      "",
		Port:      "",
		keyPrefix: "%s/broker/rabbitmq/%s",
	}
}

func (c *BrokerRabbitMQ) WithConsulClient(env string, kv *api.KV) *BrokerRabbitMQ {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "RABBITMQ_USERNAME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_USERNAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_USERNAME is required")
	}
	c.Username = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RABBITMQ_PASSWORD"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_PASSWORD from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_PASSWORD is required")
	}
	c.Password = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RABBITMQ_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_HOST is required")
	}
	c.Host = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RABBITMQ_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_PORT is required")
	}
	c.Port = string(pair.Value)
	return c
}
