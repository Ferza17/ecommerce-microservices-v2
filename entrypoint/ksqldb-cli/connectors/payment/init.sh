#!/bin/sh

echo "REGISTER PAYMENT NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/payment/sink_pg_payments_payment_providers.json \
     http://kafka-connect-local:8083/connectors

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/payment/sink_pg_payments_payments.json \
     http://kafka-connect-local:8083/connectors

echo "DONE REGISTER PAYMENT NAMESPACE CONNECTOR"
