#!/bin/sh

echo "INIT CONFIG KAFKA NAMESPACE PRODUCT TOPICS"

## LOCAL
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_CREATED "snapshot-products-product_created"
consul kv put local/broker/kafka/TOPICS/PRODUCT/CONFIRM/PRODUCT_CREATED "confirm-snapshot-products-product_created"
consul kv put local/broker/kafka/TOPICS/PRODUCT/COMPENSATE/PRODUCT_CREATED "compensate-snapshot-products-product_created"

consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_UPDATED "snapshot-products-product_updated"
consul kv put local/broker/kafka/TOPICS/PRODUCT/CONFIRM/PRODUCT_UPDATED "confirm-snapshot-products-product_updated"
consul kv put local/broker/kafka/TOPICS/PRODUCT/COMPENSATE/PRODUCT_UPDATED "compensate-snapshot-products-product_updated"

consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_DELETED "snapshot-products-product_deleted"
consul kv put local/broker/kafka/TOPICS/PRODUCT/CONFIRM/PRODUCT_DELETED "confirm-snapshot-products-product_deleted"
consul kv put local/broker/kafka/TOPICS/PRODUCT/COMPENSATE/PRODUCT_DELETED "compensate-snapshot-products-product_deleted"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_CREATED "snapshot-products-product_created"
consul kv put production/broker/kafka/TOPICS/PRODUCT/CONFIRM/PRODUCT_CREATED "confirm-snapshot-products-product_created"
consul kv put production/broker/kafka/TOPICS/PRODUCT/COMPENSATE/PRODUCT_CREATED "compensate-snapshot-products-product_created"

consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_UPDATED "snapshot-products-product_updated"
consul kv put production/broker/kafka/TOPICS/PRODUCT/CONFIRM/PRODUCT_UPDATED "confirm-snapshot-products-product_updated"
consul kv put production/broker/kafka/TOPICS/PRODUCT/COMPENSATE/PRODUCT_UPDATED "compensate-snapshot-products-product_updated"

consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_DELETED "snapshot-products-product_deleted"
consul kv put production/broker/kafka/TOPICS/PRODUCT/CONFIRM/PRODUCT_DELETED "confirm-snapshot-products-product_deleted"
consul kv put production/broker/kafka/TOPICS/PRODUCT/COMPENSATE/PRODUCT_DELETED "compensate-snapshot-products-product_deleted"

echo "DONE INIT CONFIG KAFKA NAMESPACE PRODUCT TOPICS"
