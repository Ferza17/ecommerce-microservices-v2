package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type BrokerKafkaTopicConnectorSinkMongoEvent struct {
	EventEnvelopes    string
	DlqEventEnvelopes string

	SourceConnectorEventEnvelopes string

	keyPrefix string
}

func DefaultBrokerKafkaTopicConnectorSinkMongoEvent() *BrokerKafkaTopicConnectorSinkMongoEvent {
	return &BrokerKafkaTopicConnectorSinkMongoEvent{
		keyPrefix: "%s/broker/kafka/TOPICS/CONNECTOR/%s/MONGO/EVENT/%s",
	}
}

func (c *Config) withBrokerKafkaTopicConnectorSinkMongoEvent(kv *api.KV) *Config {
	c.BrokerKafkaTopicConnectorSinkMongoEvent = DefaultBrokerKafkaTopicConnectorSinkMongoEvent().WithConsulClient(c.Env, kv)
	return c
}

func (c *BrokerKafkaTopicConnectorSinkMongoEvent) WithConsulClient(env string, kv *api.KV) *BrokerKafkaTopicConnectorSinkMongoEvent {
	t := fmt.Sprintf(c.keyPrefix, env, "SINK", "EVENT_ENVELOPES")
	pair, _, err := kv.Get(t, nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EVENT_ENVELOPES from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | EVENT_ENVELOPES is required")
	}
	c.EventEnvelopes = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "SINK", "DLQ/EVENT_ENVELOPES"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get DLQ/EVENT_ENVELOPES from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | DLQ/EVENT_ENVELOPES is required")
	}
	c.DlqEventEnvelopes = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "SOURCE", "EVENT_ENVELOPES"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get EVENT_ENVELOPES from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SOURCE EVENT_ENVELOPES is required")
	}
	c.SourceConnectorEventEnvelopes = string(pair.Value)
	return c
}
