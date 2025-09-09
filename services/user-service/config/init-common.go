package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initCommon(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/common/SAGA_STATUS/PENDING", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SAGA_STATUS/PENDING from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SAGA_STATUS/PENDING host is required")
	}
	c.CommonSagaStatusPending = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/common/SAGA_STATUS/SUCCESS", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SAGA_STATUS/SUCCESS from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SAGA_STATUS/SUCCESS host is required")
	}
	c.CommonSagaStatusSuccess = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/common/SAGA_STATUS/FAILED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SAGA_STATUS/FAILED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SAGA_STATUS/FAILED host is required")
	}
	c.CommonSagaStatusFailed = string(pair.Value)
}
