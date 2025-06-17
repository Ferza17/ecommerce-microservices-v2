package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initRabbitmqExchange(kv *api.KV) {
	// EXCHANGE
	pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/EVENT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/EVENT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/EVENT is required")
	}
	c.ExchangeEvent = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/NOTIFICATION", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/NOTIFICATION from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/NOTIFICATION is required")
	}
	c.ExchangeNotification = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/PAYMENT/DIRECT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/PAYMENT/DIRECT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/PAYMENT/DIRECT is required")
	}
	c.ExchangePaymentDirect = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/EXCHANGE/PAYMENT/DELAYED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EXCHANGE/PAYMENT/DELAYED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EXCHANGE/PAYMENT/DELAYED is required")
	}
	c.ExchangePaymentDelayed = string(pair.Value)
}
