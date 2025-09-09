package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initQueuePayment(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get QUEUE/PAYMENT/ORDER/CREATED host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | QUEUE/PAYMENT/ORDER/CREATED host is required")
	}
	c.QueuePaymentOrderCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get QUEUE/PAYMENT/ORDER/CANCELLED host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | QUEUE/PAYMENT/ORDER/CANCELLED host is required")
	}
	c.QueuePaymentOrderDelayedCancelled = string(pair.Value)
}
