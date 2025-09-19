package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type BrokerKafka struct {
	Broker1        string
	SchemaRegistry string
	keyPrefix      string
}

func DefaultKafkaBroker() *BrokerKafka {
	return &BrokerKafka{
		Broker1:        "",
		SchemaRegistry: "",
		keyPrefix:      "%s/broker/kafka/%s",
	}
}

func (c *Config) withBrokerKafka(kv *api.KV) *Config {
	c.BrokerKafka = DefaultKafkaBroker().WithConsulClient(c.Env, kv)
	return c
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

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "SCHEMA_REGISTRY"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SCHEMA_REGISTRY from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SCHEMA_REGISTRY is required")
	}
	c.SchemaRegistry = string(pair.Value)

	return c
}
