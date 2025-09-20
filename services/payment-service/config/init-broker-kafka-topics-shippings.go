package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type BrokerKafkaTopicShippings struct {
	ShippingCreated string
	ShippingUpdated string
	keyPrefix       string
}

func DefaultKafkaBrokerTopicShippings() *BrokerKafkaTopicShippings {
	return &BrokerKafkaTopicShippings{
		ShippingCreated: "",
		ShippingUpdated: "",
		keyPrefix:       "%s/broker/kafka/TOPICS/SHIPPING/%s",
	}
}

func (c *Config) withBrokerKafkaTopicShippings(kv *api.KV) *Config {
	c.BrokerKafkaTopicShippings = DefaultKafkaBrokerTopicShippings().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicShippings) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicShippings {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "SHIPPING_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SHIPPING_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SHIPPING_CREATED is required")
	}
	c.ShippingCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "SHIPPING_UPDATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SHIPPING_UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SHIPPING_UPDATED is required")
	}
	c.ShippingUpdated = string(pair.Value)
	return c
}
