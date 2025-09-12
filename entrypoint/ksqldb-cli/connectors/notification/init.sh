#!/bin/sh

echo "REGISTER NOTIFICATION NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/notification/source_connector_notifications_template.json \
     http://kafka-connect-local:8083/connectors

echo "DONE REGISTER NOTIFICATION NAMESPACE CONNECTOR"
