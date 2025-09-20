#!/bin/sh

echo "REGISTER USER NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/user/SINK_PG_CONNECTOR_USER_USERS.json \
     http://kafka-connect-local:8083/connectors


echo "DONE REGISTER USER NAMESPACE CONNECTOR"
