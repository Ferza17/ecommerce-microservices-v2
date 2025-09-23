package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type BrokerKafkaTopicProducts struct {
	ProductCreated string
	ProductUpdated string
	ProductDeleted string
	keyPrefix      string
}

func DefaultBrokerKafkaTopicProducts() *BrokerKafkaTopicProducts {
	return &BrokerKafkaTopicProducts{
		ProductCreated: "",
		ProductUpdated: "",
		ProductDeleted: "",
		keyPrefix:      "%s/broker/kafka/TOPICS/PRODUCT/%s",
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

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PRODUCT_UPDATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PRODUCT_UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PRODUCT_UPDATED is required")
	}
	c.ProductUpdated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PRODUCT_DELETED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PRODUCT_DELETED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PRODUCT_DELETED is required")
	}
	c.ProductDeleted = string(pair.Value)

	return c
}
