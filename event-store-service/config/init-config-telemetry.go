package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initTelemetry(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/telemetry/jaeger/JAEGER_TELEMETRY_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get telemetry host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | telemetry host is required")
	}
	c.JaegerTelemetryHost = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf("%s/telemetry/jaeger/JAEGER_TELEMETRY_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get telemetry host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | telemetry host is required")
	}
	c.JaegerTelemetryPort = string(pair.Value)
}
