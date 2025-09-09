package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type MessageBrokerRabbitMQ struct {
	Username  string
	Password  string
	Host      string
	Port      string
	keyPrefix string
}

func DefaultMessageBrokerRabbitMQ() *MessageBrokerRabbitMQ {
	return &MessageBrokerRabbitMQ{
		Username:  "",
		Password:  "",
		Host:      "",
		Port:      "",
		keyPrefix: "%s/broker/rabbitmq/%s",
	}
}

func (c *MessageBrokerRabbitMQ) WithConsulClient(env string, kv *api.KV) *MessageBrokerRabbitMQ {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "RABBITMQ_USERNAME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_USERNAME host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_USERNAME host is required")
	}
	c.Username = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RABBITMQ_PASSWORD"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_PASSWORD host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_PASSWORD host is required")
	}
	c.Password = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RABBITMQ_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_HOST host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_HOST host is required")
	}
	c.Host = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RABBITMQ_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_PORT host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_PORT host is required")
	}
	c.Port = string(pair.Value)

	return c
}
