#!/bin/sh

echo "INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE NOTIFICATION"

## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/NOTIFICATION/NOTIFICATION-TEMPLATES "source-mongo-notification-notification_templates"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/NOTIFICATION/DLQ/NOTIFICATION-TEMPLATES "dlq-source-mongo-notification-notification_templates"


## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/NOTIFICATION/NOTIFICATION-TEMPLATES "source-mongo-notification-notification_templates"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/NOTIFICATION/DLQ/NOTIFICATION-TEMPLATES "dlq-source-mongo-notification-notification_templates"

echo "DONE INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE NOTIFICATION"
