package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initQueueNotification(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/NOTIFICATION/EMAIL/OTP/CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | /QUEUE/NOTIFICATION/EMAIL/OTP/CREATED is required")
	}
	c.QueueNotificationEmailOtpCreated = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get /QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | /QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED is required")
	}
	c.QueueNotificationEmailPaymentOrderCreated = string(pair.Value)
}
