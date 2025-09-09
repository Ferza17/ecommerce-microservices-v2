package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initShippingService(kv *api.KV) {

	pair, _, err := kv.Get(fmt.Sprintf("%s/services/shipping/SERVICE_NAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.ShippingServiceServiceName = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/shipping/RPC_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_HOST is required")
	}
	c.ShippingServiceRpcHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/shipping/RPC_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_PORT is required")
	}
	c.ShippingServiceRpcPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/shipping/HTTP_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_HOST is required")
	}
	c.ShippingServiceHttpHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/shipping/HTTP_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_PORT is required")
	}
	c.ShippingServiceHttpPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/shipping/METRIC_HTTP_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get METRIC_HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | user/METRIC_HTTP_PORT is required")
	}
	c.ShippingServiceMetricHttpPort = string(pair.Value)
}
