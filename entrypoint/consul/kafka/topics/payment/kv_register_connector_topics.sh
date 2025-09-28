#!/bin/sh

echo "INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE PAYMENT"

## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENT-PROVIDERS "sink-pg-payments-payment_providers"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/DLQ/PAYMENT-PROVIDERS "dlq-sink-pg-payments-payment_providers"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENT-ITEMS "sink-pg-payments-payment_items"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/DLQ/PAYMENT-ITEMS "dlq-sink-pg-payments-payment_items"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENTS "sink-pg-payments-payments"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/DLQ/PAYMENTS "dlq-sink-pg-payments-payments"


## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENT-PROVIDERS "sink-pg-payments-payment_providers"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/DLQ/PAYMENT-PROVIDERS "dlq-sink-pg-payments-payment_providers"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENT-ITEMS "sink-pg-payments-payment_items"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/DLQ/PAYMENT-ITEMS "dlq-sink-pg-payments-payment_items"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENTS "sink-pg-payments-payments"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/DLQ/PAYMENTS "dlq-sink-pg-payments-payments"
echo "DONE INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE PAYMENT"
