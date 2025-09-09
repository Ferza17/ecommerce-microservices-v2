package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

//TODO: HTTP HOST & PORT
func (c *Config) initUserService(kv *api.KV) {

	pair, _, err := kv.Get(fmt.Sprintf("%s/services/user/SERVICE_NAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.UserServiceServiceName = string(pair.Value)


	pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/RPC_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_HOST is required")
	}
	c.UserServiceRpcHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/services/user/RPC_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_PORT is required")
	}
	c.UserServiceRpcPort = string(pair.Value)
}
