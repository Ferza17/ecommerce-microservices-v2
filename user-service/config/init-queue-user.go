package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initQueueUser(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/USER/CREATED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/USER/CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | /QUEUE/USER/CREATED is required")
	}
	c.QueueUserCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/USER/UPDATED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/USER/UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | /QUEUE/USER/CREATED UPDATED required")
	}
	c.QueueUserUpdated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/USER/LOGIN", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/USER/LOGIN from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | /QUEUE/USER/CREATED LOGIN required")
	}
	c.QueueUserLogin = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/USER/LOGOUT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/USER/LOGOUT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | /QUEUE/USER/CREATED LOGOUT required")
	}
	c.QueueUserLogout = string(pair.Value)
}
