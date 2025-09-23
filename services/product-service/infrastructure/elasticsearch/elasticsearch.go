package elasticsearch

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
	"time"
)

type (
	IElasticsearchInfrastructure interface {
		GetBulkIndexer() (esutil.BulkIndexer, error)
		GetClient() *elasticsearch.Client
	}

	elasticsearchInfrastructure struct {
		logger                  logger.IZapLogger
		client                  *elasticsearch.Client
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}
)

var Set = wire.NewSet(NewElasticsearchInfrastructure)

func NewElasticsearchInfrastructure(
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) IElasticsearchInfrastructure {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%s", config.Get().DatabaseElasticsearch.Host, config.Get().DatabaseElasticsearch.Port),
		},
		Username: config.Get().DatabaseElasticsearch.Username,
		Password: config.Get().DatabaseElasticsearch.Password,
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
