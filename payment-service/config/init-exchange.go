package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initExchange(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/COMMERCE", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/COMMERCE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/COMMERCE host is required")
	}
	c.ExchangeCommerce = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/EVENT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/EVENT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/EVENT host is required")
	}
	c.ExchangeEvent = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/NOTIFICATION", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/NOTIFICATION from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/NOTIFICATION host is required")
	}
	c.ExchangeNotification = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/PRODUCT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/PRODUCT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/PRODUCT host is required")
	}
	c.ExchangeProduct = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/USER", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/USER from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/USER host is required")
	}
	c.ExchangeUser = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/PAYMENT/DELAYED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/PAYMENT/DELAYED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/PAYMENT/DELAYED host is required")
	}
	c.ExchangePaymentDelayed = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/PAYMENT/DIRECT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/PAYMENT/DIRECT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/PAYMENT/DIRECT host is required")
	}
	c.ExchangePaymentDirect = string(pair.Value)
}
