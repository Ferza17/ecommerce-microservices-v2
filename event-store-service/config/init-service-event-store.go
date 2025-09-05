package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type ServiceEventStore struct {
	ServiceName    string
	RpcHost        string
	RpcPort        string
	HttpHost       string
	HttpPort       string
	MetricHttpPort string
}

func DefaultServiceEventStore() *ServiceEventStore {
	return &ServiceEventStore{
		ServiceName:    "event-store-service",
		RpcHost:        "",
		RpcPort:        "",
		HttpHost:       "",
		HttpPort:       "",
		MetricHttpPort: ""}
}

func (c *ServiceEventStore) WithConsulClient(env string, kv *api.KV) *ServiceEventStore {

	pair, _, err := kv.Get(fmt.Sprintf("%s/services/event-store/SERVICE_NAME", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.ServiceName = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/RPC_HOST", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_HOST is required")
	}
	c.RpcHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/RPC_PORT", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_PORT is required")
	}
	c.RpcPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/HTTP_HOST", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_HOST is required")
	}
	c.HttpHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/HTTP_PORT", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_PORT is required")
	}
	c.HttpPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/event-store/METRIC_HTTP_PORT", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get METRIC_HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | METRIC_HTTP_PORT is required")
	}
	c.MetricHttpPort = string(pair.Value)

	return c
}
