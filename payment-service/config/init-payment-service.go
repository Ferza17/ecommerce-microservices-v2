package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

func (c *Config) initPaymentService(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/services/payment/SERVICE_NAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get PAYMENT/SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | PAYMENT/SERVICE_NAME is required")
	}
	c.ServiceName = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/notification/SERVICE_NAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get notification/SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | NOTIFICATION/SERVICE_NAME is required")
	}
	c.NotificationServiceName = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/payment/RPC_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_HOST is required")
	}
	c.RpcHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/payment/RPC_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_PORT is required")
	}
	c.RpcPort = string(pair.Value)

	// Pay
	pair, _, err = kv.Get(fmt.Sprintf("%s/services/payment/PAYMENT_ORDER_CANCELLED_IN_MS", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get  from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul |  is required")
	}
	temp, err := strconv.ParseInt(string(pair.Value), 10, 64)
	if err != nil {
		log.Fatalf("SetConfig | could not parse PAYMENT_ORDER_CANCELLED_IN_MS to int: %v", err)
	}
	c.PaymentOrderCancelledInMs = int(temp) // Explicitly cast int64 to int
}
