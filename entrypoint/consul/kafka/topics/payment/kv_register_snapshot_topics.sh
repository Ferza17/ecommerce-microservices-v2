#!/bin/sh

echo "INIT CONFIG KAFKA NAMESPACE PAYMENT TOPICS"

## LOCAL
consul kv put local/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED "snapshot-payments-payment_order_created"
consul kv put local/broker/kafka/TOPICS/PAYMENT/CONFIRM/PAYMENT_ORDER_CREATED "confirm-snapshot-payments-payment_order_created"
consul kv put local/broker/kafka/TOPICS/PAYMENT/COMPENSATE/PAYMENT_ORDER_CREATED "compensate-snapshot-payments-payment_order_created"

consul kv put local/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED_DELAYED "snapshot-payments-payment_order_cancelled_delayed"
consul kv put local/broker/kafka/TOPICS/PAYMENT/CONFIRM/PAYMENT_ORDER_CREATED_DELAYED "confirm-snapshot-payments-payment_order_cancelled_delayed"
consul kv put local/broker/kafka/TOPICS/PAYMENT/COMPENSATE/PAYMENT_ORDER_CREATED_DELAYED "compensate-snapshot-payments-payment_order_cancelled_delayed"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED "snapshot-payments-payment_order_created"
consul kv put production/broker/kafka/TOPICS/PAYMENT/CONFIRM/PAYMENT_ORDER_CREATED "confirm-snapshot-payments-payment_order_created"
consul kv put production/broker/kafka/TOPICS/PAYMENT/COMPENSATE/PAYMENT_ORDER_CREATED "compensate-snapshot-payments-payment_order_created"

consul kv put production/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED_DELAYED "snapshot-payments-payment_order_cancelled_delayed"
consul kv put production/broker/kafka/TOPICS/PAYMENT/CONFIRM/PAYMENT_ORDER_CREATED_DELAYED "confirm-snapshot-payments-payment_order_cancelled_delayed"
consul kv put production/broker/kafka/TOPICS/PAYMENT/COMPENSATE/PAYMENT_ORDER_CREATED_DELAYED "compensate-snapshot-payments-payment_order_cancelled_delayed"

echo "DONE INIT CONFIG KAFKA NAMESPACE PAYMENT TOPICS"
