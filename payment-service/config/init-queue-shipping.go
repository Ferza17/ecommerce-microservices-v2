package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initQueueShipping(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/SHIPPING/CREATED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get QUEUE/SHIPPING/CREATED host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | QUEUE/SHIPPING/CREATED host is required")
	}
	c.QueueShippingCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/SHIPPING/UPDATED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get QUEUE/SHIPPING/UPDATED host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | QUEUE/SHIPPING/UPDATED host is required")
	}
	c.QueueShippingUpdated = string(pair.Value)
}
