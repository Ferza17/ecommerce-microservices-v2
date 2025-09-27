#!/bin/sh

echo "REGISTER PRODUCT NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/product/SINK_CONNECTOR_PG_PRODUCTS_PRODUCTS.json \
     http://kafka-connect-local:8083/connectors

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/product/SINK_CONNECTOR_ES_PRODUCTS_PRODUCTS.json \
     http://kafka-connect-local:8083/connectors


echo "DONE REGISTER PRODUCT NAMESPACE CONNECTOR"
