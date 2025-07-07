#!/bin/sh

echo "⏳ Waiting for Consul..."
until curl -s http://localhost:8500/v1/status/leader | grep -q '"'; do
  sleep 1
done


# LOAD CONFIG RABBITMQ
##########################################################################################################################################################################################################################
source /rabbitmq/init-rabbitmq.sh
source /rabbitmq/init-rabbitmq-exchange.sh
source /rabbitmq/init-rabbitmq-queue.sh
source /rabbitmq/init-rabbitmq-proxy.sh

# LOAD CONFIG DATABASE
##########################################################################################################################################################################################################################
# ELASTICSEARCH
source /database/elasticsearch/init-database-elasticsearch.sh
source /database/elasticsearch/init-database-elasticsearch-proxy.sh

# POSTGRES
source /database/postgres/init-database-postgresql.sh
source /database/postgres/init-database-postgresql-proxy.sh

# REDIS
source /database/redis/init-database-redis.sh
source /database/redis/init-database-redis-proxy.sh

# MONGODB
source /database/mongodb/init-database-mongodb.sh
source /database/mongodb/init-database-mongodb-proxy.sh


# LOAD CONFIG SMTP
##########################################################################################################################################################################################################################
source /smtp/init-smtp.sh
source /smtp/init-smtp-proxy.sh


# LOAD CONFIG TELEMETRY
##########################################################################################################################################################################################################################
source /telemetry/init-telemetry.sh
source /telemetry/init-telemetry-proxy.sh

# LOAD CONFIG PROMETHEUS AND METRICS COLLECTOR
##########################################################################################################################################################################################################################
source /prometheus/init-prometheus-proxy.sh
source /postgres-exporter/register.sh


# LOAD CONFIG TRAEFIK
##########################################################################################################################################################################################################################
source /traefik/init-traefik-proxy.sh

# LOAD CONFIG SERVICES
##########################################################################################################################################################################################################################
source /services/init-api-gateway-service.sh
source /services/init-commerce-service.sh
source /services/init-event-store-service.sh
source /services/init-notification-service.sh
source /services/init-payment-service.sh
source /services/init-product-service.sh
source /services/init-user-service.sh

##########################################################################################################################################################################################################################

echo "INIT CONFIG COMMON"
## LOCAL
consul kv put local/common/SAGA_STATUS/PENDING 'PENDING'
consul kv put local/common/SAGA_STATUS/SUCCESS 'SUCCESS'
consul kv put local/common/SAGA_STATUS/FAILED 'FAILED'

## PRODUCTION
consul kv put production/common/SAGA_STATUS/PENDING 'PENDING'
consul kv put production/common/SAGA_STATUS/SUCCESS 'SUCCESS'
consul kv put production/common/SAGA_STATUS/FAILED 'FAILED'

echo "✅ Done setting key-values."
