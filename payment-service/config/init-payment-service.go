package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initPaymentService(kv *api.KV) {

	pair, _, err := kv.Get(fmt.Sprintf("%s/services/payment/SERVICE_NAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.PaymentServiceServiceName = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/payment/RPC_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_HOST is required")
	}
	c.PaymentServiceRpcHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/payment/RPC_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_PORT is required")
	}
	c.PaymentServiceRpcPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/payment/HTTP_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_HOST is required")
	}
	c.PaymentServiceHttpHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/payment/HTTP_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_PORT is required")
	}
	c.PaymentServiceHttpPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/payment/METRIC_HTTP_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get METRIC_HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | METRIC_HTTP_PORT is required")
	}
	c.PaymentServiceMetricHttpPort = string(pair.Value)
}
