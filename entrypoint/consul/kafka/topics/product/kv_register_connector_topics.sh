#!/bin/sh

echo "INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE PRODUCT"
## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/ES/PRODUCT/PRODUCTS "sink-es-products-products"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/ES/PRODUCT/DLQ/PRODUCTS "dlq-sink-es-products-products"

## LOCAL
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PRODUCT/PRODUCTS "sink-pg-products-products"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PRODUCT/DLQ/PRODUCTS "dlq-sink-pg-products-products"

echo "DONE INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE PRODUCT"
