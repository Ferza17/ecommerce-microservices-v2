#!/bin/sh

echo "REGISTER EVENT NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/event/SINK_MONGO_EVENTS_COMMERCE_EVENT_STORES.json \
     http://kafka-connect-local:8083/connectors

echo "done SINK_MONGO_EVENTS_COMMERCE_EVENT_STORES"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/event/SINK_MONGO_EVENTS_NOTIFICATION_EVENT_STORES.json \
     http://kafka-connect-local:8083/connectors

echo "done SINK_MONGO_EVENTS_NOTIFICATION_EVENT_STORES"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/event/SINK_MONGO_EVENTS_PAYMENT_EVENT_STORES.json \
     http://kafka-connect-local:8083/connectors

echo "done SINK_MONGO_EVENTS_PAYMENT_EVENT_STORES"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/event/SINK_MONGO_EVENTS_PRODUCT_EVENT_STORES.json \
     http://kafka-connect-local:8083/connectors

echo "done SINK_MONGO_EVENTS_PRODUCT_EVENT_STORES"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/event/SINK_MONGO_EVENTS_SHIPPING_EVENT_STORES.json \
     http://kafka-connect-local:8083/connectors

echo "done SINK_MONGO_EVENTS_SHIPPING_EVENT_STORES"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/event/SINK_MONGO_EVENTS_USER_EVENT_STORES.json \
     http://kafka-connect-local:8083/connectors

echo "done SINK_MONGO_EVENTS_USER_EVENT_STORES"

echo "DONE REGISTER EVENT NAMESPACE CONNECTOR"
