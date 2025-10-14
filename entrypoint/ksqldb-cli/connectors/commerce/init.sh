#!/bin/sh

echo "REGISTER COMMERCE NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/commerce/SINK_MONGO_COMMERCE_CARTS.json \
     http://kafka-connect-local:8083/connectors

echo "done SINK_MONGO_COMMERCE_CARTS"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/commerce/SINK_MONGO_COMMERCE_WISHLISTS.json \
     http://kafka-connect-local:8083/connectors

echo "done SINK_MONGO_COMMERCE_WISHLISTS"
