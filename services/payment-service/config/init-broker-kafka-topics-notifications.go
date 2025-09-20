package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type BrokerKafkaTopicNotifications struct {
	PaymentOrderCreated          string
	PaymentOrderDelayedCancelled string
	keyPrefix                    string
}

func DefaultKafkaBrokerTopicNotifications() *BrokerKafkaTopicNotifications {
	return &BrokerKafkaTopicNotifications{
		PaymentOrderCreated:          "",
		PaymentOrderDelayedCancelled: "",
		keyPrefix:                    "%s/broker/kafka/TOPICS/PAYMENT/%s",
	}
}

func (c *Config) withBrokerKafkaTopicNotifications(kv *api.KV) *Config {
	c.BrokerKafkaTopicNotifications = DefaultKafkaBrokerTopicNotifications().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicNotifications) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicNotifications {
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
	c.PaymentOrderDelayedCancelled = string(pair.Value)
	return c
}
