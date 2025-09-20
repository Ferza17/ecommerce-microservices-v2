package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/util"
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
		logger.Info("insert-mock products to postgresql via sink connector ")
		kafkaInfra := kafka.ProvideKafkaInfrastructure()

		data, err := os.ReadFile("mock_products.json")
		if err != nil {
			logger.Error("error reading file ", zap.Error(err))
			panic(err)
		}

		var products []orm.Product
		now, err := util.GetNowWithTimeZone(pkgContext.CtxValueAsiaJakarta)
		if err != nil {
			logger.Error("error getting now with timezone ", zap.Error(err))
			panic(err)
		}
		if err := json.Unmarshal(data, &products); err != nil {
			logger.Error("error unmarshalling data ", zap.Error(err))
			panic(err)
		}

		for _, product := range products {
			product.CreatedAt = &now
			product.UpdatedAt = &now

			logger.Info("insert product to Postgresql via sink connector ", zap.Any("product", product))
			if err = kafkaInfra.PublishWithJsonSchema(cmd.Context(), config.Get().BrokerKafkaTopicConnectorSinkProduct.PgProducts, product.ID, product); err != nil {
				logger.Error(fmt.Sprintf("error publishing to kafka with id : %s ,name: %s", product.ID, product.Name), zap.Error(err))
				continue
			}

			logger.Info("insert product to Elasticsearch via sink connector ", zap.Any("product", product))
			if err = kafkaInfra.PublishWithJsonSchema(cmd.Context(), config.Get().BrokerKafkaTopicConnectorSinkProduct.EsProducts, product.ID, product); err != nil {
				logger.Error(fmt.Sprintf("error publishing to kafka with id : %s ,name: %s", product.ID, product.Name), zap.Error(err))
				continue
			}

			time.Sleep(500 * time.Millisecond)
		}

		if err := kafkaInfra.Close(); err != nil {
			logger.Error("error closing kafka infrastructure ", zap.Error(err))
			panic(err)
		}
	},
}
