#!/bin/sh

echo "REGISTER PAYMENT NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/payment/SINK_PG_CONNECTOR.json \
     http://kafka-connect-local:8083/connectors

#echo "INIT KSQLDB payment sink connector"
#ksql http://ksqldb-server-local:8088 < /connectors/payment/sink_connector.sql

echo "DONE REGISTER PAYMENT NAMESPACE CONNECTOR"
