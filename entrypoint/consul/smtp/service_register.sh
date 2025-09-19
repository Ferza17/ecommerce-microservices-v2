#!/bin/sh

echo "Registering Mailhog SMTP to Consul ..."


# Wait for Mailhog to be available (max 10 retries)
echo "Waiting for Mailhog to be available..."
max_retries=10
count=0

until nc -z mailhog-local 1025 >/dev/null 2>&1; do
  count=$((count+1))
  echo "Mailhog not ready yet, attempt $count/$max_retries..."

  if [ $count -ge $max_retries ]; then
    echo "⚠️ Mailhog still not available after $max_retries attempts, continuing anyway..."
    break
  fi
  sleep 2
done
echo "Mailhog check finished"

# Register Mailhog SMTP service
consul services register \
-name=mailhog \
-id=mailhog-main \
-port=1025 \
-address=mailhog-local \
-tag=smtp \
-tag=proxy \
-tag=mailhog \


# Health check for Mailhog
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "mailhog-ui-health",
"Name": "Mailhog UI Health Check",
"HTTP": "http://mailhog-local:8025/",
"Interval": "30s",
"Timeout": "5s",
"ServiceID": "mailhog-main"
}'
