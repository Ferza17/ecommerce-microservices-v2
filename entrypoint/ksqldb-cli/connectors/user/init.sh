#!/bin/sh

echo "REGISTER USER NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/user/SINK_PG_CONNECTOR.json \
     http://kafka-connect-local:8083/connectors

#curl -X POST -H "Content-Type: application/json" \
#     --data @/connectors/user/sink_pg_users_users.json \
#     http://kafka-connect-local:8083/connectors

#echo "INIT KSQLDB user sink connector"
#ksql http://ksqldb-server-local:8088 < /connectors/user/sink_connector.sql


echo "DONE REGISTER USER NAMESPACE CONNECTOR"
