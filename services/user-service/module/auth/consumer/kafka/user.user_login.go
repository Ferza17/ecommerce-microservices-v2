package kafka

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"go.uber.org/zap"
)

func (c *authConsumer) UserLogin(ctx context.Context) error {

	cs, err := c.kafkaInfrastructure.Consume(config.Get().BrokerKafkaTopic.UserUserLogin)
	if err != nil {
		c.logger.Error("UserLogin", zap.Error(err))
		return err
	}

	for _, consumer := range cs {
		for {
			select {
			case message := <-consumer.Messages():
				if message == nil || message.Value == nil {
					continue
				}

			case err = <-consumer.Errors():
				c.logger.Error("UserLogin", zap.Error(err))
				return err
			}
		}
	}
	return nil
}
