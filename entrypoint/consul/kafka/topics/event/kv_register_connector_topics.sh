#!/bin/sh

echo "INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE EVENT"

## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/COMMERCE_EVENT_STORE "sink-mongo-events-commerce_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/COMMERCE_EVENT_STORE "dlq-sink-mongo-events-commerce_event_stores"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/NOTIFICATION_EVENT_STORE "sink-mongo-events-notification_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/NOTIFICATION_EVENT_STORE "dlq-sink-mongo-events-notification_event_stores"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/PAYMENT_EVENT_STORE "sink-mongo-events-payment_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/PAYMENT_EVENT_STORE "dlq-sink-mongo-events-payment_event_stores"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/PRODUCT_EVENT_STORE "sink-mongo-events-product_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/PRODUCT_EVENT_STORE "dlq-sink-mongo-events-product_event_stores"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/SHIPPING_EVENT_STORE "sink-mongo-events-shipping_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/SHIPPING_EVENT_STORE "dlq-sink-mongo-events-shipping_event_stores"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/USER_EVENT_STORE "sink-mongo-events-user_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/USER_EVENT_STORE "dlq-sink-mongo-events-user_event_stores"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/COMMERCE_EVENT_STORE "sink-mongo-events-commerce_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/COMMERCE_EVENT_STORE "dlq-sink-mongo-events-commerce_event_stores"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/NOTIFICATION_EVENT_STORE "sink-mongo-events-notification_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/NOTIFICATION_EVENT_STORE "dlq-sink-mongo-events-notification_event_stores"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/PAYMENT_EVENT_STORE "sink-mongo-events-payment_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/PAYMENT_EVENT_STORE "dlq-sink-mongo-events-payment_event_stores"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/PRODUCT_EVENT_STORE "sink-mongo-events-product_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/PRODUCT_EVENT_STORE "dlq-sink-mongo-events-product_event_stores"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/SHIPPING_EVENT_STORE "sink-mongo-events-shipping_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/SHIPPING_EVENT_STORE "dlq-sink-mongo-events-shipping_event_stores"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/USER_EVENT_STORE "sink-mongo-events-user_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/USER_EVENT_STORE "dlq-sink-mongo-events-user_event_stores"

echo "DONE INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE EVENT"
