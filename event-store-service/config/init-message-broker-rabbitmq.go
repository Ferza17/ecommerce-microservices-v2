package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type MessageBrokerRabbitMQ struct {
	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string
}

func DefaultMessageBrokerRabbitMQ() *MessageBrokerRabbitMQ {
	return &MessageBrokerRabbitMQ{
		RabbitMQUsername: "",
		RabbitMQPassword: "",
		RabbitMQHost:     "",
		RabbitMQPort:     "",
	}
}

func (c *MessageBrokerRabbitMQ) WithConsulClient(env string, kv *api.KV) *MessageBrokerRabbitMQ {
	pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_USERNAME", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_USERNAME host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_USERNAME host is required")
	}
	c.RabbitMQUsername = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_PASSWORD", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_PASSWORD host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_PASSWORD host is required")
	}
	c.RabbitMQPassword = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_HOST", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_HOST host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_HOST host is required")
	}
	c.RabbitMQHost = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/RABBITMQ_PORT", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RABBITMQ_PORT host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RABBITMQ_PORT host is required")
	}
	c.RabbitMQPort = string(pair.Value)

	return c
}
