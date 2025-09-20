package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type BrokerKafkaTopicConnectorSinkPgPayment struct {
	Payments         string
	PaymentItems     string
	PaymentProviders string
	keyPrefix        string
}

func DefaultBrokerKafkaTopicsConnectorSinkPgPayment() *BrokerKafkaTopicConnectorSinkPgPayment {
	return &BrokerKafkaTopicConnectorSinkPgPayment{
		Payments:         "",
		PaymentItems:     "",
		PaymentProviders: "",
		keyPrefix:        "%s/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/%s",
	}
}

func (c *Config) withBrokerKafkaTopicConnectorSinkPgPayment(kv *api.KV) *Config {
	c.BrokerKafkaTopicConnectorSinkPgPayment = DefaultBrokerKafkaTopicsConnectorSinkPgPayment().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicConnectorSinkPgPayment) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicConnectorSinkPgPayment {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "PAYMENTS"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PAYMENTS from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PAYMENTS is required")
	}
	c.Payments = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PAYMENT-ITEMS"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PAYMENT-ITEMS from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PAYMENT-ITEMS is required")
	}
	c.PaymentItems = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PAYMENT-PROVIDERS"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PAYMENT-PROVIDERS from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PAYMENT-PROVIDERS is required")
	}
	c.PaymentProviders = string(pair.Value)

	return c
}
