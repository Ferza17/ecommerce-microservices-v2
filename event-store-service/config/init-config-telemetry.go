package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type ConfigTelemetry struct {
	JaegerTelemetryHost string
	JaegerTelemetryPort string
}

func DefaultConfigTelemetry() *ConfigTelemetry {
	return &ConfigTelemetry{
		JaegerTelemetryHost: "",
		JaegerTelemetryPort: "",
	}
}

func (c *ConfigTelemetry) WithConsulClient(env string, kv *api.KV) *ConfigTelemetry {
	pair, _, err := kv.Get(fmt.Sprintf("%s/telemetry/jaeger/JAEGER_TELEMETRY_HOST", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get telemetry host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | telemetry host is required")
	}
	c.JaegerTelemetryHost = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf("%s/telemetry/jaeger/JAEGER_TELEMETRY_PORT", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get telemetry host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | telemetry host is required")
	}
	c.JaegerTelemetryPort = string(pair.Value)
	return c
}
