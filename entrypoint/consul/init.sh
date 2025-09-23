#!/bin/sh

echo "⏳ Waiting for Consul..."
until curl -s http://consul-local:8500/v1/status/leader | grep -q '"'; do
  sleep 1
done

# LOAD CONFIG KAFKA
##########################################################################################################################################################################################################################
sh ./kafka/kv_register.sh
sh ./kafka/kv_register_topic.sh
sh ./kafka/kv_register_connector.sh

# LOAD CONFIG DATABASE
##########################################################################################################################################################################################################################
# ELASTICSEARCH
sh ./database/elasticsearch/kv_register.sh
sh ./database/elasticsearch/service_register.sh

# POSTGRES
sh ./database/postgres/kv_register.sh
sh ./database/postgres/service_register.sh

# REDIS
sh ./database/redis/kv_register.sh
sh ./database/redis/service_register.sh

# MONGODB
sh ./database/mongodb/kv_register.sh
sh ./database/mongodb/service_register.sh


# LOAD CONFIG SMTP
##########################################################################################################################################################################################################################
sh ./smtp/kv_register.sh
sh ./smtp/service_register.sh

# LOAD CONFIG TELEMETRY
##########################################################################################################################################################################################################################
 sh ./telemetry/kv_register.sh
 sh ./telemetry/service_register.sh

# LOAD CONFIG PROMETHEUS AND METRICS COLLECTOR
##########################################################################################################################################################################################################################
sh ./prometheus/service_register.sh
sh ./postgres-exporter/service_register.sh

# LOAD CONFIG OPEN POLICY AGENT
##########################################################################################################################################################################################################################
 sh ./opa/kv_register.sh
# sh ./opa/register.sh

# LOAD CONFIG TRAEFIK
##########################################################################################################################################################################################################################
sh ./traefik/service_register.sh

# LOAD CONFIG SERVICES
##########################################################################################################################################################################################################################
sh ./services/kv_register_api_gateway.sh
sh ./services/kv_register_commerce.sh
sh ./services/kv_register_event_store.sh
sh ./services/kv_register_notification.sh
sh ./services/kv_register_payment.sh
sh ./services/kv_register_product.sh
sh ./services/kv_register_shipping.sh
sh ./services/kv_register_user.sh

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


# --- Set a flag in Consul KV to indicate initialization is complete ---
# This is crucial for the healthcheck to detect successful setup
echo "Setting Consul KV initialization flag: config/setup/initialized=true"
consul kv put config/setup/initialized true

echo "✅ Done setting key-values."
