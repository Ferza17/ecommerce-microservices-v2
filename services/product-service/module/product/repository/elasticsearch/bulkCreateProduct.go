package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
)

func (r *productElasticsearchRepository) BulkCreateProduct(ctx context.Context, products []*orm.Product) error {
	bi, err := r.elasticsearchInfrastructure.GetBulkIndexer()
	if err != nil {
		r.logger.Error(fmt.Sprintf("error while get bulk indexer %v", err))
		return err
	}

	for _, product := range products {
		data, err := json.Marshal(product)
		if err != nil {
			r.logger.Error(fmt.Sprintf("error marshaling product: %v", err))
			return err
		}

		if err = bi.Add(ctx,
			esutil.BulkIndexerItem{
				Index:      config.Get().BrokerKafkaTopicConnectorSinkProduct.EsProducts,
				DocumentID: product.ID,
				Action:     "index",
				Body:       bytes.NewReader(data),
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					r.logger.Error(fmt.Sprintf("error while indexing product: %v", err))
				},
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					r.logger.Info(fmt.Sprintf("successfully indexed product: %v", item.DocumentID))
				},
			},
		); err != nil {
			r.logger.Error(fmt.Sprintf("error while indexing product: %v", err))
			return err
		}
	}
	if err = bi.Close(ctx); err != nil {
		r.logger.Error(fmt.Sprintf("error while indexing product: %v", err))
		return err
	}

	return nil
}
