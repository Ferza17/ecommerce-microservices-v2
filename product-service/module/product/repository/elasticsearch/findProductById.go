package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"log"
	"time"
)

func (r *productElasticsearchRepository) FindProductById(ctx context.Context, requestId string, id string) (*orm.Product, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()

	ctxTimeout, span := r.telemetryInfrastructure.Tracer(ctxTimeout, "Repository.Elasticsearch.FindProductById")
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

	var doc map[string]any
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		r.logger.Error(fmt.Sprintf("error while decode response: %v", err))
		return nil, err
	}

	if _, ok := doc["_source"]; !ok {
		r.logger.Error("product not found")
		return nil, fmt.Errorf("product not found")
	}

	var product orm.Product
	source, err := json.MarshalIndent(doc["_source"], "", "  ")
	if err != nil {
		r.logger.Error(fmt.Sprintf("error marshaling product: %v", err))
		return nil, err
	}

	if err = json.Unmarshal(source, &product); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return nil, err
	}

	return &product, nil
}
