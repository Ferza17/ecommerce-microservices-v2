#!/bin/sh

echo "INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE EVENT"
# LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/EVENT_ENVELOPES "sink-mongo-events-event_envelopes"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/EVENT_ENVELOPES "dlq-sink-mongo-events-event_envelopes"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/EVENT/EVENT_ENVELOPES "source-mongo-events-event_envelopes"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/EVENT/DLQ/EVENT_ENVELOPES "dlq-source-mongo-events-event_envelopes"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/EVENT_ENVELOPES "sink-mongo-events-event_envelopes"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/DLQ/EVENT_ENVELOPES "dlq-sink-mongo-events-event_envelopes"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/EVENT/EVENT_ENVELOPES "source-mongo-events-event_envelopes"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/EVENT/DLQ/EVENT_ENVELOPES "dlq-source-mongo-events-event_envelopes"

echo "DONE INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE EVENT"
