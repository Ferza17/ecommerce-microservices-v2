package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type BrokerKafkaTopicUsers struct {
	UserUserCreated           string
	ConfirmUserUserCreated    string
	CompensateUserUserCreated string

	UserUserUpdated string

	UserUserLogin           string
	ConfirmUserUserLogin    string
	CompensateUserUserLogin string

	UserUserLogout           string
	ConfirmUserUserLogout    string
	CompensateUserUserLogout string

	keyPrefix string
}

func DefaultKafkaBrokerTopicUsers() *BrokerKafkaTopicUsers {
	return &BrokerKafkaTopicUsers{
		keyPrefix: "%s/broker/kafka/TOPICS/USER/%s",
	}
}

func (c *Config) withBrokerKafkaTopicUsers(kv *api.KV) *Config {
	c.BrokerKafkaTopicUsers = DefaultKafkaBrokerTopicUsers().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicUsers) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicUsers {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "USER_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get USER_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | USER_CREATED is required")
	}
	c.UserUserCreated = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "CONFIRM/USER_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get CONFIRM/USER_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | CONFIRM/USER_CREATED is required")
	}
	c.ConfirmUserUserCreated = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "COMPENSATE/USER_CREATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMPENSATE/USER_CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMPENSATE/USER_CREATED is required")
	}
	c.CompensateUserUserCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "USER_UPDATED"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get USER_UPDATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | USER_UPDATED is required")
	}

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "USER_LOGIN"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get USER_LOGIN from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | USER_LOGIN is required")
	}
	c.UserUserLogin = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "CONFIRM/USER_LOGIN"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get CONFIRM/USER_LOGIN from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | CONFIRM/USER_LOGIN is required")
	}
	c.ConfirmUserUserLogin = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "COMPENSATE/USER_LOGIN"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMPENSATE/USER_LOGIN from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMPENSATE/USER_LOGIN is required")
	}
	c.CompensateUserUserLogin = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "USER_LOGOUT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get USER_LOGOUT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | USER_LOGOUT is required")
	}
	c.UserUserLogout = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "CONFIRM/USER_LOGOUT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get CONFIRM/USER_LOGOUT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | CONFIRM/USER_LOGOUT is required")
	}
	c.ConfirmUserUserLogout = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "COMPENSATE/USER_LOGOUT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get COMPENSATE/USER_LOGOUT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | COMPENSATE/USER_LOGOUT is required")
	}
	c.CompensateUserUserLogout = string(pair.Value)
	return c
}
