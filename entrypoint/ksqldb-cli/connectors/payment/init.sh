#!/bin/sh

echo "REGISTER PAYMENT NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/payment/SINK_PG_CONNECTOR_PAYMENT_ITEMS.json \
     http://kafka-connect-local:8083/connectors \

echo "done SINK_PG_CONNECTOR_PAYMENT_ITEMS"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/payment/SINK_PG_CONNECTOR_PAYMENT_PROVIDERS.json \
     http://kafka-connect-local:8083/connectors \

echo "done SINK_PG_CONNECTOR_PAYMENT_PROVIDERS"


curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/payment/SINK_PG_CONNECTOR_PAYMENTS.json \
     http://kafka-connect-local:8083/connectors \

echo "done SINK_PG_CONNECTOR_PAYMENTS"

echo "DONE REGISTER PAYMENT NAMESPACE CONNECTOR"
