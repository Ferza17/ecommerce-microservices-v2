package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"log"
)

func (r *productElasticsearchRepository) FindProductById(ctx context.Context, requestId string, id string) (*orm.Product, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "Repository.Elasticsearch.FindProductById")
	defer span.End()

	res, err := r.elasticsearchInfrastructure.GetClient().Get(productIndex, id)
	if err != nil {
		r.logger.Error(fmt.Sprintf("error while get product by id: %v", err))
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		r.logger.Error(fmt.Sprintf("error response: %s", res.String()))
		return nil, fmt.Errorf("error response: %s", res.String())
	}

	// Decode the response body
	var doc struct {
		Source orm.Product `json:"_source"`
	}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	return &doc.Source, nil
}
