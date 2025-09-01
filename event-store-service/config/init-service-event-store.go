package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initServiceEventStore(kv *api.KV) {

	pair, _, err := kv.Get(fmt.Sprintf("%s/services/event-store/SERVICE_NAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.EventStoreServiceServiceName = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/RPC_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_HOST is required")
	}
	c.EventStoreServiceRpcHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/RPC_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_PORT is required")
	}
	c.EventStoreServiceRpcPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/HTTP_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_HOST is required")
	}
	c.EventStoreServiceHttpHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/HTTP_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_PORT is required")
	}
	c.EventStoreServiceHttpPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/METRIC_HTTP_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get METRIC_HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | METRIC_HTTP_PORT is required")
	}
	c.EventStoreServiceMetricHttpPort = string(pair.Value)
}
