#!/bin/bash
set -e

#COMPONENT_DIR="/usr/share/kafka/plugins"
#CONNECT_PROPS="/etc/ksqldb-server/connect.properties"
#
## Ensure plugin directory exists
#if [ ! -d "$COMPONENT_DIR" ]; then
#  echo "Creating $COMPONENT_DIR"
#  mkdir -p "$COMPONENT_DIR"
#fi
#
## Install Confluent Hub client if missing
#if ! command -v confluent-hub &> /dev/null; then
#    echo "Installing Confluent Hub client..."
#    curl -sL https://client.hub.confluent.io/confluent-hub-client-latest.tar.gz -o /tmp/confluent-hub.tar.gz
#    mkdir -p /usr/local/confluent-hub
#    tar -xzf /tmp/confluent-hub.tar.gz -C /usr/local/confluent-hub
#    ln -s /usr/local/confluent-hub/bin/confluent-hub /usr/local/bin/confluent-hub
#    rm /tmp/confluent-hub.tar.gz
#fi
#
## Install JDBC connector if missing
#if [ ! -d "$COMPONENT_DIR/confluentinc-kafka-connect-jdbc" ]; then
#  echo "Installing Kafka Connect JDBC..."
#  confluent-hub install --no-prompt confluentinc/kafka-connect-jdbc:latest \
#    --component-dir "$COMPONENT_DIR" \
#    --worker-configs "$CONNECT_PROPS"
#fi
#
## Install Elasticsearch connector if missing
#if [ ! -d "$COMPONENT_DIR/confluentinc-kafka-connect-elasticsearch" ]; then
#  echo "Installing Kafka Connect Elasticsearch..."
#  confluent-hub install --no-prompt confluentinc/kafka-connect-elasticsearch:latest \
#    --component-dir "$COMPONENT_DIR" \
#    --worker-configs "$CONNECT_PROPS"
#fi
#
## Install MongoDB connector if missing
#if [ ! -d "$COMPONENT_DIR/mongodb-kafka-connect-mongodb" ]; then
#  echo "Installing Kafka Connect MongoDB..."
#  confluent-hub install --no-prompt mongodb/kafka-connect-mongodb:latest \
#    --component-dir "$COMPONENT_DIR" \
#    --worker-configs "$CONNECT_PROPS"
#fi
#
## Install Redis connector if missing
#if [ ! -d "$COMPONENT_DIR/jcustenborder-kafka-connect-redis" ]; then
#  echo "Installing Kafka Connect Redis..."
#  confluent-hub install --no-prompt jcustenborder/kafka-connect-redis:latest \
#    --component-dir "$COMPONENT_DIR" \
#    --worker-configs "$CONNECT_PROPS"
#fi
#
#echo "✅ All plugins installed."
#
## Wait a moment for plugins to be fully extracted
#sleep 15
#
#
## List installed plugins for verification
#echo "Installed plugins:"
#ls -a "$COMPONENT_DIR"

# Start KSQL server with Connect enabled


exec ksql-server-start /etc/ksqldb-server/ksql-server.properties