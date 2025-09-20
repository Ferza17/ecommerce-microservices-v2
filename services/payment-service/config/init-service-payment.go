package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/consul/api"
)

type ConfigServicePayment struct {
	ServiceName string

	RpcHost                   string
	RpcPort                   string
	HttpHost                  string
	HttpPort                  string
	MetricHttpPort            string
	PaymentOrderCancelledInMs int

	keyPrefix string
}

func DefaultConfigServicePayment() *ConfigServicePayment {
	return &ConfigServicePayment{
		ServiceName:               "",
		RpcHost:                   "",
		RpcPort:                   "",
		HttpHost:                  "",
		HttpPort:                  "",
		MetricHttpPort:            "",
		PaymentOrderCancelledInMs: 0,
		keyPrefix:                 "%s/services/payment/%s",
	}
}

func (c *Config) withConfigServicePayment(kv *api.KV) *Config {
	c.ConfigServicePayment = DefaultConfigServicePayment().WithConsulClient(c.Env, kv)
	return c
}

func (c *ConfigServicePayment) WithConsulClient(env string, kv *api.KV) *ConfigServicePayment {

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

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "PAYMENT_ORDER_CANCELLED_IN_MS"), nil)
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
	c.PaymentOrderCancelledInMs = int(temp)

	return c
}
