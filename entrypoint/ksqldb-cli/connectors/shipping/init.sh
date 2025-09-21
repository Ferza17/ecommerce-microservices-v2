#!/bin/sh

echo "REGISTER SHIPPING NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/shipping/SINK_PG_CONNECTOR_SHIPPING_SHIPPING_PROVIDERS.json \
     http://kafka-connect-local:8083/connectors \

echo "done SINK_PG_CONNECTOR_SHIPPING_SHIPPING_PROVIDERS"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/shipping/SINK_PG_CONNECTOR_SHIPPING_SHIPPINGS.json \
     http://kafka-connect-local:8083/connectors \

echo "done SINK_PG_CONNECTOR_SHIPPING_SHIPPINGS"

echo "DONE REGISTER SHIPPING NAMESPACE CONNECTOR"
