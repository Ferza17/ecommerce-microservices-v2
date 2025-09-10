package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type BrokerKafka struct {
	Broker1   string
	keyPrefix string
}

func DefaultKafkaBroker() *BrokerKafka {
	return &BrokerKafka{
		Broker1:   "",
		keyPrefix: "%s/broker/kafka/%s",
	}
}

func (c *BrokerKafka) WithConsulClient(env string, kv *api.KV) *BrokerKafka {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "BROKER_1"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get BROKER_1 from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | BROKER_1 is required")
	}
	c.Broker1 = string(pair.Value)

	return c
}
