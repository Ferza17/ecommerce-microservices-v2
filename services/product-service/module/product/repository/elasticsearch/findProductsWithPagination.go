package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"time"

	"log"
	"strings"
)

func (r *productElasticsearchRepository) FindProductsWithPagination(ctx context.Context, requestId string, request *productRpc.FindProductsWithPaginationRequest) ([]*orm.Product, int64, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()

	ctxTimeout, span := r.telemetryInfrastructure.StartSpanFromContext(ctxTimeout, "Repository.Elasticsearch.FindProductsWithPagination")
	defer span.End()

	reqBody := map[string]interface{}{}

	if request.Names != nil || request.Ids != nil {
		reqBody["query"] = map[string]interface{}{}

		reqBody["query"].(map[string]interface{})["bool"] = map[string]interface{}{
			"should": []interface{}{},
		}

		if request.Ids != nil {
			reqBody["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"] = append(reqBody["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"].([]interface{}), map[string]interface{}{
				"ids": map[string]interface{}{
					"values": request.Ids,
				},
			})
		}

		if request.Names != nil {
			reqBody["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"] = append(reqBody["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"].([]interface{}), map[string]interface{}{
				"multi_match": map[string]interface{}{
					"query": strings.Join(request.Names, " "),
					"fields": []string{
						"name",
					},
					"type":     "phrase_prefix",
					"operator": "OR",
				},
			})
		}
	}

	offset := int((request.Page - 1) * request.Limit)
	queryJSON, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error encoding query: %v", err)
		return nil, 0, err
	}

	searchResult, err := r.elasticsearchInfrastructure.GetClient().Search(
		r.elasticsearchInfrastructure.GetClient().Search.WithContext(ctxTimeout),
		r.elasticsearchInfrastructure.GetClient().Search.WithIndex(config.Get().BrokerKafkaTopicConnectorSinkProduct.EsProducts),
		r.elasticsearchInfrastructure.GetClient().Search.WithBody(bytes.NewReader(queryJSON)),
		r.elasticsearchInfrastructure.GetClient().Search.WithFrom(offset),
		r.elasticsearchInfrastructure.GetClient().Search.WithSize(int(request.Limit)),
		r.elasticsearchInfrastructure.GetClient().Search.WithTrackTotalHits(true),
	)

	if err != nil {
		r.logger.Error(fmt.Sprintf("requestId: %s , error while get products: %v", requestId, err))
		return nil, 0, err
	}
	defer searchResult.Body.Close()

	// Parse response
	var searchResponse map[string]interface{}
	if err = json.NewDecoder(searchResult.Body).Decode(&searchResponse); err != nil {
		r.logger.Error(fmt.Sprintf("requestId: %s , error while parse response: %v", requestId, err))
		return nil, 0, err
	}

	// Check for Elasticsearch error response
	if errResponse, ok := searchResponse["error"]; ok {
		r.logger.Error(fmt.Sprintf("requestId: %s , elasticsearch error: %v", requestId, errResponse))
		return nil, 0, fmt.Errorf("elasticsearch error: %v", errResponse)
	}

	// Safely extract hits with nil checks
	hitsInterface, ok := searchResponse["hits"]
	if !ok || hitsInterface == nil {
		r.logger.Error(fmt.Sprintf("requestId: %s , missing hits in response", requestId))
		return nil, 0, fmt.Errorf("invalid elasticsearch response: missing hits")
	}

	hitsMap, ok := hitsInterface.(map[string]interface{})
	if !ok {
		r.logger.Error(fmt.Sprintf("requestId: %s , hits is not a map", requestId))
		return nil, 0, fmt.Errorf("invalid elasticsearch response: hits is not a map")
	}

	// Safely extract total hits
	var totalHits int64
	if totalInterface, ok := hitsMap["total"]; ok && totalInterface != nil {
		if totalMap, ok := totalInterface.(map[string]interface{}); ok {
			if valueInterface, ok := totalMap["value"]; ok && valueInterface != nil {
				if value, ok := valueInterface.(float64); ok {
					totalHits = int64(value)
				}
			}
		}
	}

	// Safely extract hits array
	var response []*orm.Product
	hitsArrayInterface, ok := hitsMap["hits"]
	if !ok || hitsArrayInterface == nil {
		return response, totalHits, nil
	}

	hitsArray, ok := hitsArrayInterface.([]interface{})
	if !ok {
		r.logger.Error(fmt.Sprintf("requestId: %s , hits array is not an array", requestId))
		return response, totalHits, nil
	}

	for _, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			r.logger.Error(fmt.Sprintf("requestId: %s , hit is not a map", requestId))
			continue
		}

		sourceInterface, ok := hitMap["_source"]
		if !ok || sourceInterface == nil {
			r.logger.Error(fmt.Sprintf("requestId: %s , missing _source in hit", requestId))
			continue
		}

		productJSON, err := json.Marshal(sourceInterface)
		if err != nil {
			r.logger.Error(fmt.Sprintf("requestId: %s , error while parse response: %v", requestId, err))
			continue
		}
		var product orm.Product
		if err := json.Unmarshal(productJSON, &product); err != nil {
			r.logger.Error(fmt.Sprintf("requestId: %s , error while parse response: %v", requestId, err))
			continue
		}
		response = append(response, &product)
	}

	return response, totalHits, nil
}
