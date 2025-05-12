package elasticsearch

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type (
	IElasticsearchInfrastructure interface{}

	elasticsearchInfrastructure struct {
		logger                  pkg.IZapLogger
		client                  *elasticsearch.Client
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}
)

func NewElasticsearchInfrastructure(
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) IElasticsearchInfrastructure {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%s", config.Get().ElasticsearchHost, config.Get().ElasticsearchPort),
		},
		Username: config.Get().ElasticsearchUsername,
		Password: config.Get().ElasticsearchPassword,
	})
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to elasticsearch: %v", err))
		return nil
	}

	if _, err = client.Ping(); err != nil {
		logger.Error(fmt.Sprintf("failed to connect to elasticsearch: %v", err))
		return nil
	}

	return &elasticsearchInfrastructure{
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
