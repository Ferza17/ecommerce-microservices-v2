package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type BrokerKafkaTopicConnectorSinkMongoEvent struct {
	Commerce     string
	Notification string
	Payment      string
	Product      string
	User         string
	keyPrefix    string
}

func DefaultBrokerKafkaTopicConnectorSinkMongoEvent() *BrokerKafkaTopicConnectorSinkMongoEvent {
	return &BrokerKafkaTopicConnectorSinkMongoEvent{
		Commerce:     "",
		Notification: "",
		Payment:      "",
		Product:      "",
		User:         "",
		keyPrefix:    "%s/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/%s",
	}
}

func (c *Config) withBrokerKafkaTopicConnectorSinkMongoEvent(kv *api.KV) *Config {
	c.BrokerKafkaTopicConnectorSinkMongoEvent = DefaultBrokerKafkaTopicConnectorSinkMongoEvent().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicConnectorSinkMongoEvent) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicConnectorSinkMongoEvent {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "COMMERCE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMMERCE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMMERCE is required")
	}
	c.Commerce = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "NOTIFICATION"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get NOTIFICATION from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | NOTIFICATION is required")
	}
	c.Notification = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PAYMENT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PAYMENT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PAYMENT is required")
	}
	c.Payment = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PRODUCT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PRODUCT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PRODUCT is required")
	}
	c.Product = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "USER"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get USER from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | USER is required")
	}
	c.User = string(pair.Value)
	return c
}
