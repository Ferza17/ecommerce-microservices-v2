#!/bin/bash

set -e

echo "=== Kafka Topic Creator ==="
echo "Waiting for Kafka to be ready..."

# Wait for Kafka to be ready
MAX_WAIT=60
WAIT_TIME=0
while ! kafka-topics --bootstrap-server kafka-local:9092 --list > /dev/null 2>&1; do
    if [ $WAIT_TIME -ge $MAX_WAIT ]; then
        echo "❌ Timeout: Kafka is not ready after ${MAX_WAIT} seconds"
        exit 1
    fi
    echo "⏳ Kafka is not ready yet. Waiting... (${WAIT_TIME}s/${MAX_WAIT}s)"
    sleep 5
    WAIT_TIME=$((WAIT_TIME + 5))
done

echo "✅ Kafka is ready!"
echo ""

# Function to extract topics from YAML using basic text processing
parse_yaml_topics() {
    local yaml_file="$1"

    # Extract topic configurations using grep and sed
    grep -A 3 "^  - name:" "$yaml_file" | while IFS= read -r line; do
        if [[ $line =~ ^[[:space:]]*-[[:space:]]*name:[[:space:]]*(.+) ]]; then
            # Extract topic name
            topic_name=$(echo "$line" | sed 's/.*name:[[:space:]]*//' | sed 's/[[:space:]]*$//')

            # Read next lines for partitions and replication_factor
            read -r partitions_line
            read -r replication_line

            partitions=$(echo "$partitions_line" | sed 's/.*partitions:[[:space:]]*//' | sed 's/[[:space:]]*$//')
            replication_factor=$(echo "$replication_line" | sed 's/.*replication_factor:[[:space:]]*//' | sed 's/[[:space:]]*$//')

            echo "$topic_name:$partitions:$replication_factor"
        fi
    done
}

# Parse topics from YAML file
echo "📋 Parsing topics from /topics.yaml..."
TOPICS=$(parse_yaml_topics "/topics.yaml")

if [ -z "$TOPICS" ]; then
    echo "❌ No topics found in YAML file"
    exit 1
fi

echo "📝 Found $(echo "$TOPICS" | wc -l) topics to create"
echo ""

# Create topics
CREATED_COUNT=0
FAILED_COUNT=0

while IFS= read -r topic_config; do
    if [ -z "$topic_config" ]; then
        continue
    fi

    IFS=':' read -r name partitions replication_factor <<< "$topic_config"

    echo "🔄 Creating topic: $name"
    echo "   ├── Partitions: $partitions"
    echo "   └── Replication Factor: $replication_factor"

    if kafka-topics --create \
        --bootstrap-server kafka-local:9092 \
        --topic "$name" \
        --partitions "$partitions" \
        --replication-factor "$replication_factor" \
        --if-not-exists > /dev/null 2>&1; then
        echo "   ✅ Success"
        CREATED_COUNT=$((CREATED_COUNT + 1))
    else
        echo "   ❌ Failed"
        FAILED_COUNT=$((FAILED_COUNT + 1))
    fi
    echo ""
done <<< "$TOPICS"

# Summary
echo "=== Summary ==="
echo "✅ Successfully created: $CREATED_COUNT topics"
echo "❌ Failed to create: $FAILED_COUNT topics"
echo ""

# List all topics to verify
echo "📋 Current topics in Kafka:"
kafka-topics --bootstrap-server kafka-local:9092 --list | sort

echo ""
echo "🎉 Topic creation process completed!"
echo "🌐 You can view topics in Kafka UI at: http://localhost:8080"