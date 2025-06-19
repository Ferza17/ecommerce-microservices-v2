package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initRabbitmq(kv *api.KV) {
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
}
