package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type ServiceEventStoreRabbitMQ struct {
	ExchangeEventFanout              string
	QueueEventEventCreated           string
	QueueEventApiGatewayEventCreated string
	keyPrefix                        string
}

func DefaultServiceEventStoreRabbitMQ() *ServiceEventStoreRabbitMQ {
	return &ServiceEventStoreRabbitMQ{
		ExchangeEventFanout:              "",
		QueueEventEventCreated:           "",
		QueueEventApiGatewayEventCreated: "",
		keyPrefix:                        "%s/broker/rabbitmq/%s",
	}
}

func (c *ServiceEventStoreRabbitMQ) WithConsulClient(env string, kv *api.KV) *ServiceEventStoreRabbitMQ {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "EXCHANGE/EVENT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/EVENT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/EVENT is required")
	}
	c.ExchangeEventFanout = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "QUEUE/EVENT/EVENT/EVENT-CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/EVENT/EVENT/EVENT-CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | QUEUE/EVENT/EVENT/EVENT-CREATED is required")
	}
	c.QueueEventEventCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "QUEUE/EVENT/API-GATEWAY/EVENT-CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/EVENT/API-GATEWAY/EVENT-CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | QUEUE/EVENT/API-GATEWAY/EVENT-CREATED is required")
	}
	c.QueueEventApiGatewayEventCreated = string(pair.Value)
	return c
}
