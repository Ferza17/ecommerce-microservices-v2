package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/consul/api"
)

func (c *Config) RegisterConsulService() error {
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%s:%s", c.ConsulHost, c.ConsulPort),
	})
	if err != nil {
		log.Fatalf("SetConfig | could not connect to consul: %v", err)
	}

	port, err := strconv.ParseInt(c.ConfigServiceUser.RpcPort, 10, 64)
	if err != nil {
		log.Fatalf("SetConfig | could not parse PORT to int: %v", err)
	}
	if err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Kind:    api.ServiceKindTypical,
		Name:    c.ConfigServiceUser.ServiceName,
		Address: c.ConfigServiceUser.RpcHost,
		Port:    int(port),
		Tags:    []string{"service", "kafka-subscriber"},
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%s", c.ConfigServiceUser.RpcHost, c.ConfigServiceUser.RpcPort),
			GRPCUseTLS:                     false,
			Interval:                       "30s", // Less frequent checks
			Timeout:                        "5s",  // Reasonable timeout
			DeregisterCriticalServiceAfter: "40s", // Give more time before deregistering
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
