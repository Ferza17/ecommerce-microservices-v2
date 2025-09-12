#!/bin/sh

echo "REGISTER PRODUCT NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/product/sink_es_products_products.json \
     http://kafka-connect-local:8083/connectors

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/product/sink_pg_products_products.json \
     http://kafka-connect-local:8083/connectors

echo "DONE REGISTER PRODUCT NAMESPACE CONNECTOR"
