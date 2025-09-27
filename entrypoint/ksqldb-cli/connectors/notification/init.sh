#!/bin/sh

echo "REGISTER NOTIFICATION NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/notification/SOURCE_CONNECTOR_MONGO_NOTIFICATIONS_NOTIFICATION_TEMPLATES.json \
     http://kafka-connect-local:8083/connectors

echo "DONE REGISTER NOTIFICATION NAMESPACE CONNECTOR"
