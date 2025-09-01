package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initDatabaseMongodb(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_USERNAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_USERNAME host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_USERNAME host is required")
	}
	c.MongoUsername = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_PASSWORD", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_PASSWORD host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_PASSWORD host is required")
	}
	c.MongoPassword = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_HOST host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_HOST host is required")
	}
	c.MongoHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_PORT host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_PORT host is required")
	}
	c.MongoPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_DATABASE_NAME/EVENT_STORE host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_DATABASE_NAME/EVENT_STORE host is required")
	}
	c.MongoDatabaseName = string(pair.Value)
}
