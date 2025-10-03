package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type BrokerKafkaTopicConnectorSinkMongoEvent struct {
	Commerce    string
	DlqCommerce string

	Notification    string
	DlqNotification string

	Payment    string
	DlqPayment string

	Product    string
	DlqProduct string

	User    string
	DlqUser string

	keyPrefix string
}

func DefaultBrokerKafkaTopicConnectorSinkMongoEvent() *BrokerKafkaTopicConnectorSinkMongoEvent {
	return &BrokerKafkaTopicConnectorSinkMongoEvent{
		keyPrefix: "%s/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/%s",
	}
}

func (c *Config) withBrokerKafkaTopicConnectorSinkMongoEvent(kv *api.KV) *Config {
	c.BrokerKafkaTopicConnectorSinkMongoEvent = DefaultBrokerKafkaTopicConnectorSinkMongoEvent().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicConnectorSinkMongoEvent) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicConnectorSinkMongoEvent {
	t := fmt.Sprintf(c.keyPrefix, env, "COMMERCE_EVENT_STORE")
	pair, _, err := kv.Get(t, nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMMERCE_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMMERCE_EVENT_STORE is required")
	}
	c.Commerce = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "DLQ/COMMERCE_EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get DLQ/COMMERCE_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | DLQ/COMMERCE_EVENT_STORE is required")
	}
	c.DlqCommerce = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "NOTIFICATION_EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get NOTIFICATION_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | NOTIFICATION_EVENT_STORE is required")
	}
	c.Notification = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "DLQ/NOTIFICATION_EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get DLQ/NOTIFICATION_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | DLQ/NOTIFICATION_EVENT_STORE is required")
	}
	c.DlqNotification = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PAYMENT_EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PAYMENT_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PAYMENT_EVENT_STORE is required")
	}
	c.Payment = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "DLQ/PAYMENT_EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get DLQ/PAYMENT_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | DLQ/PAYMENT_EVENT_STORE is required")
	}
	c.DlqPayment = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PRODUCT_EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PRODUCT_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PRODUCT_EVENT_STORE is required")
	}
	c.Product = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "DLQ/PRODUCT_EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get DLQ/PRODUCT_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | DLQ/PRODUCT_EVENT_STORE is required")
	}
	c.DlqProduct = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "USER_EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get USER_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | USER_EVENT_STORE is required")
	}
	c.User = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "DLQ/USER_EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get DLQ/USER_EVENT_STORE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | DLQ/USER_EVENT_STORE is required")
	}
	c.DlqUser = string(pair.Value)
	return c
}
