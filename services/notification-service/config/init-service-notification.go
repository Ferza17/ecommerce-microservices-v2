package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type ConfigServiceNotification struct {
	ServiceName    string
	RpcHost        string
	RpcPort        string
	HttpHost       string
	HttpPort       string
	MetricHttpPort string

	keyPrefix string
}

func DefaultConfigServiceNotification() *ConfigServiceNotification {
	return &ConfigServiceNotification{
		ServiceName:    "",
		RpcHost:        "",
		RpcPort:        "",
		HttpHost:       "",
		HttpPort:       "",
		MetricHttpPort: "",
		keyPrefix:      "%s/services/notification/%s",
	}
}

func (c *Config) withConfigServiceNotification(kv *api.KV) *Config {
	c.ConfigServiceNotification = DefaultConfigServiceNotification().WithConsulClient(c.Env, kv)
	return c
}

func (c *ConfigServiceNotification) WithConsulClient(env string, kv *api.KV) *ConfigServiceNotification {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "SERVICE_NAME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.ServiceName = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RPC_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_HOST is required")
	}
	c.RpcHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RPC_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_PORT is required")
	}
	c.RpcPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "HTTP_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_HOST is required")
	}
	c.HttpHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "HTTP_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_PORT is required")
	}
	c.HttpPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "METRIC_HTTP_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get METRIC_HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | METRIC_HTTP_PORT is required")
	}
	c.MetricHttpPort = string(pair.Value)
	return c
}
