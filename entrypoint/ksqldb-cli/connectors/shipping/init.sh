#!/bin/sh

echo "REGISTER SHIPPING NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/shipping/SINK_PG_CONNECTOR.json \
     http://kafka-connect-local:8083/connectors

#curl -X POST -H "Content-Type: application/json" \
#     --data @/connectors/shipping/sink_pg_shippings_shippings.json \
#     http://kafka-connect-local:8083/connectors

#echo "INIT KSQLDB shipping sink connector"
#ksql http://ksqldb-server-local:8088 < /connectors/shipping/sink_connector.sql


echo "DONE REGISTER SHIPPING NAMESPACE CONNECTOR"
