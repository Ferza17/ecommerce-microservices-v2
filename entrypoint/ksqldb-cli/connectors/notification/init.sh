#!/bin/sh

echo "REGISTER NOTIFICATION NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/notification/SOURCE_MONGO_CONNECTOR.json \
     http://kafka-connect-local:8083/connectors

#echo "INIT KSQLDB notification source connector"
#ksql http://ksqldb-server-local:8088 < /connectors/notification/source_connector.sql


echo "DONE REGISTER NOTIFICATION NAMESPACE CONNECTOR"
