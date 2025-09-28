#!/bin/sh

echo "INIT CONFIG KAFKA NAMESPACE SHIPPING TOPICS"

## LOCAL
consul kv put local/broker/kafka/TOPICS/SHIPPING/SHIPPING_CREATED "snapshot-shippings-shipping_created"
consul kv put local/broker/kafka/TOPICS/SHIPPING/CONFIRM/SHIPPING_CREATED "confirm-snapshot-shippings-shipping_created"
consul kv put local/broker/kafka/TOPICS/SHIPPING/COMPENSATE/SHIPPING_CREATED "compensate-snapshot-shippings-shipping_created"

consul kv put local/broker/kafka/TOPICS/SHIPPING/SHIPPING_UPDATED "snapshot-shippings-shipping_updated"
consul kv put local/broker/kafka/TOPICS/SHIPPING/CONFIRM/SHIPPING_UPDATED "confirm-snapshot-shippings-shipping_updated"
consul kv put local/broker/kafka/TOPICS/SHIPPING/COMPENSATE/SHIPPING_UPDATED "compensate-snapshot-shippings-shipping_updated"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/SHIPPING/SHIPPING_CREATED "snapshot-shippings-shipping_created"
consul kv put production/broker/kafka/TOPICS/SHIPPING/CONFIRM/SHIPPING_CREATED "confirm-snapshot-shippings-shipping_created"
consul kv put production/broker/kafka/TOPICS/SHIPPING/COMPENSATE/SHIPPING_CREATED "compensate-snapshot-shippings-shipping_created"

consul kv put production/broker/kafka/TOPICS/SHIPPING/SHIPPING_UPDATED "snapshot-shippings-shipping_updated"
consul kv put production/broker/kafka/TOPICS/SHIPPING/CONFIRM/SHIPPING_UPDATED "confirm-snapshot-shippings-shipping_updated"
consul kv put production/broker/kafka/TOPICS/SHIPPING/COMPENSATE/SHIPPING_UPDATED "compensate-snapshot-shippings-shipping_updated"

echo "DONE INIT CONFIG KAFKA NAMESPACE SHIPPING TOPICS"
