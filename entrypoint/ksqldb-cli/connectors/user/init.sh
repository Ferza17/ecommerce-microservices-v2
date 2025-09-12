#!/bin/sh

echo "REGISTER USER NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/user/sink_pg_users_roles.json \
     http://kafka-connect-local:8083/connectors

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/user/sink_pg_users_users.json \
     http://kafka-connect-local:8083/connectors

echo "DONE REGISTER USER NAMESPACE CONNECTOR"
