#!/bin/sh

echo "REGISTER PRODUCT NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/product/SINK_CONNECTOR_PG_PRODUCTS_PRODUCTS.json \
     http://kafka-connect-local:8083/connectors

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/product/SINK_CONNECTOR_ES_PRODUCTS_PRODUCTS.json \
     http://kafka-connect-local:8083/connectors

#echo "INIT KSQLDB product sink connector"
#ksql http://ksqldb-server-local:8088 < /connectors/product/sink_connector.sql

echo "DONE REGISTER PRODUCT NAMESPACE CONNECTOR"
