package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type DatabaseElasticsearch struct {
	Host     string
	Port     string
	Username string
	Password string

	keyPrefix string
}

func DefaultDatabaseElasticsearch() *DatabaseElasticsearch {
	return &DatabaseElasticsearch{
		Host:      "",
		Port:      "",
		Username:  "",
		Password:  "",
		keyPrefix: "%s/database/elasticsearch/%s",
	}
}

func (c *Config) withDatabaseElasticsearch(kv *api.KV) *Config {
	c.DatabaseElasticsearch = DefaultDatabaseElasticsearch().WithConsulClient(c.Env, kv)
	return c
}

func (c *DatabaseElasticsearch) WithConsulClient(env string, kv *api.KV) *DatabaseElasticsearch {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "ELASTICSEARCH_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ELASTICSEARCH_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ELASTICSEARCH_HOST is required")
	}
	c.Host = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "ELASTICSEARCH_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ELASTICSEARCH_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ELASTICSEARCH_PORT is required")
	}
	c.Port = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "ELASTICSEARCH_USERNAME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ELASTICSEARCH_USERNAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ELASTICSEARCH_USERNAME is required")
	}
	c.Username = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "ELASTICSEARCH_PASSWORD"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ELASTICSEARCH_PASSWORD from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ELASTICSEARCH_PASSWORD is required")
	}
	c.Password = string(pair.Value)

	return c
}
