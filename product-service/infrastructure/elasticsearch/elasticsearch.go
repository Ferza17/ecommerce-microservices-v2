package elasticsearch

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
	"time"
)

type (
	IElasticsearchInfrastructure interface {
		GetBulkIndexer() (esutil.BulkIndexer, error)
		GetClient() *elasticsearch.Client
	}

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

	if err != nil {
		logger.Error(fmt.Sprintf("failed to create a bulk indexer: %v", err))
		return nil
	}

	return &elasticsearchInfrastructure{
		client:                  client,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}

func (i *elasticsearchInfrastructure) GetBulkIndexer() (esutil.BulkIndexer, error) {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:        i.client,
		Index:         "products",
		FlushBytes:    5 * 1024 * 1024, // 5MB
		FlushInterval: 30 * time.Second,
		OnError: func(ctx context.Context, err error) {
			i.logger.Error(fmt.Sprintf("ERROR: %s", err))
			return
		},
	})
	if err != nil {
		i.logger.Error(fmt.Sprintf("failed to create a bulk indexer: %v", err))
		return nil, err
	}
	return bi, nil
}

func (i *elasticsearchInfrastructure) GetClient() *elasticsearch.Client {
	return i.client
}
