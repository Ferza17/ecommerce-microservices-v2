#!/bin/sh

echo "REGISTER EVENT NAMESPACE CONNECTOR"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/event/SINK_MONGO_EVENTS_EVENT_ENVELOPES.json \
     http://kafka-connect-local:8083/connectors

echo "done SINK_MONGO_EVENTS_EVENT_ENVELOPES"

curl -X POST -H "Content-Type: application/json" \
     --data @/connectors/event/SOURCE_MONGO_EVENTS_EVENT_ENVELOPES.json \
     http://kafka-connect-local:8083/connectors

echo "done SOURCE_MONGO_EVENTS_EVENT_ENVELOPES"

echo "DONE REGISTER EVENT NAMESPACE CONNECTOR"
