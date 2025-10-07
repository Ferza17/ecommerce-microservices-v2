package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type BrokerKafkaTopicShippings struct {
	ShippingCreated           string
	ConfirmShippingCreated    string
	CompensateShippingCreated string

	ShippingUpdated           string
	ConfirmShippingUpdated    string
	CompensateShippingUpdated string

	keyPrefix string
}

func DefaultKafkaBrokerTopicShippings() *BrokerKafkaTopicShippings {
	return &BrokerKafkaTopicShippings{
		keyPrefix: "%s/broker/kafka/TOPICS/SHIPPING/%s",
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
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "CONFIRM/SHIPPING_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get CONFIRM/SHIPPING_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | CONFIRM/SHIPPING_CREATED is required")
	}
	c.ConfirmShippingCreated = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "COMPENSATE/SHIPPING_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMPENSATE/SHIPPING_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMPENSATE/SHIPPING_CREATED is required")
	}
	c.CompensateShippingCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "SHIPPING_UPDATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SHIPPING_UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SHIPPING_UPDATED is required")
	}
	c.ShippingUpdated = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "CONFIRM/SHIPPING_UPDATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get CONFIRM/SHIPPING_UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | CONFIRM/SHIPPING_UPDATED is required")
	}
	c.ConfirmShippingUpdated = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "COMPENSATE/SHIPPING_UPDATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMPENSATE/SHIPPING_UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMPENSATE/SHIPPING_UPDATED is required")
	}
	c.CompensateShippingUpdated = string(pair.Value)
	return c
}
