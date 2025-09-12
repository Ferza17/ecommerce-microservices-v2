#!/bin/sh

echo "REGISTER SHIPPING NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/shipping/sink_pg_shippings_shipping_providers.json \
     http://kafka-connect-local:8083/connectors

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/shipping/sink_pg_shippings_shippings.json \
     http://kafka-connect-local:8083/connectors

echo "DONE REGISTER SHIPPING NAMESPACE CONNECTOR"
