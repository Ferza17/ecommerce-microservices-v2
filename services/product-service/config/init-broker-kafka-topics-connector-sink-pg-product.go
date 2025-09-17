package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type BrokerKafkaTopicConnectorSinkProduct struct {
	PgProducts string
	EsProducts string
	keyPrefix  string
}

func DefaultBrokerKafkaTopicConnectorSinkProduct() *BrokerKafkaTopicConnectorSinkProduct {
	return &BrokerKafkaTopicConnectorSinkProduct{
		PgProducts: "",
		EsProducts: "",
		keyPrefix:  "%s/broker/kafka/TOPICS/CONNECTOR/SINK/%s",
	}
}

func (c *BrokerKafkaTopicConnectorSinkProduct) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicConnectorSinkProduct {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "PG/PRODUCT/PRODUCTS"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PG/PRODUCT/PRODUCTS from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PG/PRODUCT/PRODUCTS is required")
	}
	c.PgProducts = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "ES/PRODUCT/PRODUCTS"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ES/PRODUCT/PRODUCTS from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ES/PRODUCT/PRODUCTS is required")
	}
	c.EsProducts = string(pair.Value)

	return c
}
