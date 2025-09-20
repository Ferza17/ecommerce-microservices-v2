package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/util"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"time"
)

var insertMockCommand = &cobra.Command{
	Use:   "insert-mock",
	Short: "Insert Mock to database",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logger.ProvideLogger()
		kafkaInfra := kafka.ProvideKafkaInfrastructure()

		logger.Info("insert-mock payment to postgresql via sink connector ")

		data, err := os.ReadFile("mock_shipping_providers.json")
		if err != nil {
			logger.Error("error reading file ", zap.Error(err))
			panic(err)
		}

		var providers []orm.Provider
		now, err := util.GetNowWithTimeZone(pkgContext.CtxValueAsiaJakarta)
		if err != nil {
			logger.Error("error getting now with timezone ", zap.Error(err))
			panic(err)
		}
		if err := json.Unmarshal(data, &providers); err != nil {
			logger.Error("error unmarshalling data ", zap.Error(err))
			panic(err)
		}
		for _, provider := range providers {
			provider.CreatedAt = &now
			provider.UpdatedAt = &now

			if err = kafkaInfra.PublishWithJsonSchema(cmd.Context(), config.Get().BrokerKafkaTopicConnectorSinkPgPayment.PaymentProviders, provider.ID, provider); err != nil {
				logger.Error(fmt.Sprintf("error publishing to kafka with id : %s ,name: %s", provider.ID, provider.Name), zap.Error(err))
				continue
			}

			logger.Info("inserted payment providers to Postgresql via sink connector ", zap.Any("product", provider))
			time.Sleep(500 * time.Millisecond)
		}
		if err := kafkaInfra.Close(); err != nil {
			logger.Error("error closing kafka infrastructure ", zap.Error(err))
			panic(err)
		}
	},
}
