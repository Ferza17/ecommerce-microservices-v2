#!/bin/sh

echo "INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE SHIPPING"

## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPING-PROVIDERS "sink-pg-shippings-shipping_providers"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/DLQ/SHIPPING-PROVIDERS "dlq-sink-pg-shippings-shipping_providers"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPINGS "sink-pg-shippings-shippings"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/DLQ/SHIPPINGS "dlq-sink-pg-shippings-shippings"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPING-PROVIDERS "sink-pg-shippings-shipping_providers"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/DLQ/SHIPPING-PROVIDERS "dlq-sink-pg-shippings-shipping_providers"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPINGS "sink-pg-shippings-shippings"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/DLQ/SHIPPINGS "dlq-sink-pg-shippings-shippings"

echo "DONE INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE SHIPPING"
