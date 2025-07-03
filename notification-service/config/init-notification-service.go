package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initNotificationService(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/services/notification/SERVICE_NAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.NotificationServiceServiceName = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/notification/RPC_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_HOST is required")
	}
	c.NotificationServiceRpcHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/notification/RPC_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_PORT is required")
	}
	c.NotificationServiceRpcPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/notification/HTTP_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_HOST is required")
	}
	c.NotificationServiceHttpHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/notification/METRIC_HTTP_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get METRIC_HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | METRIC_HTTP_PORT is required")
	}
	c.NotificationServiceMetricHttpPort = string(pair.Value)
}
