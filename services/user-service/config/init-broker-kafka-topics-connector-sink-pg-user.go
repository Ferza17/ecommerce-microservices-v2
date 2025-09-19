package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type BrokerKafkaTopicConnectorSinkPgUser struct {
	Users     string
	Roles     string
	keyPrefix string
}

func DefaultBrokerKafkaTopicsConnectorSinkPgUser() *BrokerKafkaTopicConnectorSinkPgUser {
	return &BrokerKafkaTopicConnectorSinkPgUser{
		Users:     "",
		Roles:     "",
		keyPrefix: "%s/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/%s",
	}
}

func (c *Config) withBrokerKafkaTopicConnectorSinkPgUser(kv *api.KV) *Config {
	c.BrokerKafkaTopicConnectorSinkPgUser = DefaultBrokerKafkaTopicsConnectorSinkPgUser().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicConnectorSinkPgUser) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicConnectorSinkPgUser {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "USERS"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get USERS from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | USERS is required")
	}
	c.Users = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "ROLES"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ROLES from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ROLES is required")
	}
	c.Roles = string(pair.Value)

	return c
}
