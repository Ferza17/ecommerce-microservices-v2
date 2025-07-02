package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

func (c *Config) RegisterConsulService() error {
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%s:%s", c.ConsulHost, c.ConsulPort),
	})
	if err != nil {
		log.Fatalf("SetConfig | could not connect to consul: %v", err)
	}

	port, err := strconv.ParseInt(c.ProductServiceRpcPort, 10, 64)
	if err != nil {
		log.Fatalf("SetConfig | could not parse PORT to int: %v", err)
	}
	if err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Kind:    api.ServiceKindTypical,
		Name:    c.ProductServiceServiceName,
		Address: c.ProductServiceRpcHost,
		Port:    int(port),
		Tags:    []string{"service", "rabbitmq-client"},
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%s", c.ProductServiceRpcHost, c.ProductServiceRpcPort),
			GRPCUseTLS:                     false,
			Interval:                       "10s", // Less frequent checks
			Timeout:                        "5s",  // Reasonable timeout
			DeregisterCriticalServiceAfter: "30s", // Give more time before deregistering
		},
		Connect: &api.AgentServiceConnect{
			Native: true,
		},
	}); err != nil {
		log.Fatalf("Error registering service: %v", err)
		return err
	}
	return nil
}
