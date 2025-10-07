package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type BrokerKafkaTopicProducts struct {
	ProductCreated           string
	ConfirmProductCreated    string
	CompensateProductCreated string

	ProductUpdated           string
	ConfirmProductUpdated    string
	CompensateProductUpdated string

	ProductDeleted           string
	ConfirmProductDeleted    string
	CompensateProductDeleted string

	keyPrefix string
}

func DefaultBrokerKafkaTopicProducts() *BrokerKafkaTopicProducts {
	return &BrokerKafkaTopicProducts{
		keyPrefix: "%s/broker/kafka/TOPICS/PRODUCT/%s",
	}
}

func (c *Config) withBrokerKafkaTopicProducts(kv *api.KV) *Config {
	c.BrokerKafkaTopicProducts = DefaultBrokerKafkaTopicProducts().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicProducts) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicProducts {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "PRODUCT_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PRODUCT_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PRODUCT_CREATED is required")
	}
	c.ProductCreated = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "CONFIRM/PRODUCT_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get CONFIRM/PRODUCT_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | CONFIRM/PRODUCT_CREATED is required")
	}
	c.ConfirmProductCreated = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "COMPENSATE/PRODUCT_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMPENSATE/PRODUCT_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMPENSATE/PRODUCT_CREATED is required")
	}
	c.CompensateProductCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PRODUCT_UPDATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PRODUCT_UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PRODUCT_UPDATED is required")
	}
	c.ProductUpdated = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "CONFIRM/PRODUCT_UPDATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get CONFIRM/PRODUCT_UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | CONFIRM/PRODUCT_UPDATED is required")
	}
	c.ConfirmProductUpdated = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "COMPENSATE/PRODUCT_UPDATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMPENSATE/PRODUCT_UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMPENSATE/PRODUCT_UPDATED is required")
	}
	c.CompensateProductUpdated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PRODUCT_DELETED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PRODUCT_DELETED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PRODUCT_DELETED is required")
	}
	c.ProductDeleted = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "CONFIRM/PRODUCT_DELETED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get CONFIRM/PRODUCT_DELETED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | CONFIRM/PRODUCT_DELETED is required")
	}
	c.ConfirmProductDeleted = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "COMPENSATE/PRODUCT_DELETED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMPENSATE/PRODUCT_DELETED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMPENSATE/PRODUCT_DELETED is required")
	}
	c.CompensateProductDeleted = string(pair.Value)

	return c
}
