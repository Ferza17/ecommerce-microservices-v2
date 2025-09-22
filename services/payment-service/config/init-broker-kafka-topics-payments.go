package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type BrokerKafkaTopicPayments struct {
	PaymentOrderCreated        string
	PaymentOrderCreatedDelayed string
	keyPrefix                  string
}

func DefaultKafkaBrokerTopicPayments() *BrokerKafkaTopicPayments {
	return &BrokerKafkaTopicPayments{
		PaymentOrderCreated:        "",
		PaymentOrderCreatedDelayed: "",
		keyPrefix:                  "%s/broker/kafka/TOPICS/PAYMENT/%s",
	}
}

func (c *Config) withBrokerKafkaTopicPayments(kv *api.KV) *Config {
	c.BrokerKafkaTopicPayments = DefaultKafkaBrokerTopicPayments().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicPayments) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicPayments {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "PAYMENT_ORDER_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PAYMENT_ORDER_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PAYMENT_ORDER_CREATED is required")
	}
	c.PaymentOrderCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PAYMENT_ORDER_CREATED_DELAYED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PAYMENT_ORDER_CREATED_DELAYED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PAYMENT_ORDER_CREATED_DELAYED is required")
	}
	c.PaymentOrderCreatedDelayed = string(pair.Value)
	return c
}
