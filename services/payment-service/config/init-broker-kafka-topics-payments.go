package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type BrokerKafkaTopicPayments struct {
	EmailOtpCreated          string
	EmailPaymentOrderCreated string
	keyPrefix                string
}

func DefaultKafkaBrokerTopicPayments() *BrokerKafkaTopicPayments {
	return &BrokerKafkaTopicPayments{
		EmailOtpCreated:          "",
		EmailPaymentOrderCreated: "",
		keyPrefix:                "%s/broker/kafka/TOPICS/NOTIFICATION/%s",
	}
}

func (c *Config) withBrokerKafkaTopicPayments(kv *api.KV) *Config {
	c.BrokerKafkaTopicPayments = DefaultKafkaBrokerTopicPayments().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicPayments) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicPayments {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "EMAIL_OTP_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EMAIL_OTP_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EMAIL_OTP_CREATED is required")
	}
	c.EmailOtpCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "EMAIL_PAYMENT_ORDER_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EMAIL_PAYMENT_ORDER_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EMAIL_PAYMENT_ORDER_CREATED is required")
	}
	c.EmailPaymentOrderCreated = string(pair.Value)
	return c
}
