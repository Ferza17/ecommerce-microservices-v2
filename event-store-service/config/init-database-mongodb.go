package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type DatabaseMongodb struct {
	MongoUsername     string
	MongoPassword     string
	MongoHost         string
	MongoPort         string
	MongoDatabaseName string
}

func DefaultDatabaseMongodb() *DatabaseMongodb {
	return &DatabaseMongodb{
		MongoUsername:     "",
		MongoPassword:     "",
		MongoHost:         "",
		MongoPort:         "",
		MongoDatabaseName: "",
	}
}

func (c *DatabaseMongodb) WithConsulClient(env string, kv *api.KV) *DatabaseMongodb {

	pair, _, err := kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_USERNAME", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_USERNAME host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_USERNAME host is required")
	}
	c.MongoUsername = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_PASSWORD", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_PASSWORD host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_PASSWORD host is required")
	}
	c.MongoPassword = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_HOST", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_HOST host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_HOST host is required")
	}
	c.MongoHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_PORT", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_PORT host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_PORT host is required")
	}
	c.MongoPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_DATABASE_NAME/EVENT_STORE host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_DATABASE_NAME/EVENT_STORE host is required")
	}
	c.MongoDatabaseName = string(pair.Value)

	return c
}
