package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type ServiceEventStoreRabbitMQ struct {
	ExchangeEventFanout string
	QueueEventCreated   string
}

func DefaultServiceEventStoreRabbitMQ() *ServiceEventStoreRabbitMQ {
	return &ServiceEventStoreRabbitMQ{
		ExchangeEventFanout: "",
		QueueEventCreated:   "",
	}
}

func (c *ServiceEventStoreRabbitMQ) WithConsulClient(env string, kv *api.KV) *ServiceEventStoreRabbitMQ {
	pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/EVENT", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE EVENT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE EVENT is required")
	}
	c.ExchangeEventFanout = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/EVENT/CREATED", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get QUEUE EVENT/CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | QUEUE EVENT CREATED is required")
	}
	c.QueueEventCreated = string(pair.Value)
	return c
}
