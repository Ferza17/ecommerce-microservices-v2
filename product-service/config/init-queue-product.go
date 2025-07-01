package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initQueueProduct(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/PRODUCT/CREATED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/PRODUCT/CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | /QUEUE/PRODUCT/CREATED is required")
	}
	c.QueueProductCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/PRODUCT/UPDATED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/PRODUCT/UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | /QUEUE/PRODUCT/UPDATED is required")
	}
	c.QueueProductUpdated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/PRODUCT/DELETED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/PRODUCT/DELETED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | /QUEUE/PRODUCT/DELETED is required")
	}
	c.QueueProductDeleted = string(pair.Value)
}
