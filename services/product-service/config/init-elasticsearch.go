package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initElasticsearch(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/database/elasticsearch/ELASTICSEARCH_USERNAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ELASTICSEARCH_USERNAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ELASTICSEARCH_USERNAME is required")
	}
	c.ElasticsearchUsername = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/elasticsearch/ELASTICSEARCH_PASSWORD", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ELASTICSEARCH_PASSWORD from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ELASTICSEARCH_PASSWORD is required")
	}
	c.ElasticsearchPassword = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/elasticsearch/ELASTICSEARCH_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ELASTICSEARCH_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ELASTICSEARCH_HOST is required")
	}
	c.ElasticsearchHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/elasticsearch/ELASTICSEARCH_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get ELASTICSEARCH_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | ELASTICSEARCH_PORT is required")
	}
	c.ElasticsearchPort = string(pair.Value)
}
