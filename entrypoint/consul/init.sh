#!/bin/sh

echo "⏳ Waiting for Consul..."
until curl -s http://localhost:8500/v1/status/leader | grep -q '"'; do
  sleep 1
done

##########################################################################################################################################################################################################################
echo "INIT CONFIG API-GATEWAY-SERVICE"
### Service Info
curl --request PUT --data 'local' http://localhost:8500/v1/kv/services/api-gateway/ENV
curl --request PUT --data 'api-gateway-service' http://localhost:8500/v1/kv/services/api-gateway/SERVICE_NAME

### rabbitMQ
curl --request PUT --data 'rabbitmq' http://localhost:8500/v1/kv/services/api-gateway/RABBITMQ_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/api-gateway/RABBITMQ_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/api-gateway/RABBITMQ_HOST # Change value localhost to container-name if you want to run service via docker-compose
curl --request PUT --data '5672' http://localhost:8500/v1/kv/services/api-gateway/RABBITMQ_PORT

### GRPC services
curl --request PUT --data 'localhost:50051' http://localhost:8500/v1/kv/services/api-gateway/PRODUCT_SERVICE_URL # Change value localhost to container-name if you want to run service via docker-compose
curl --request PUT --data 'localhost:50052' http://localhost:8500/v1/kv/services/api-gateway/USER_SERVICE_URL # Change value localhost to container-name if you want to run service via docker-compose

###HTTP SERVER
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/api-gateway/HTTP_HOST # Change value localhost to container-name if you want to run service via docker-compose
curl --request PUT --data '3000' http://localhost:8500/v1/kv/services/api-gateway/HTTP_PORT

##########################################################################################################################################################################################################################
echo "INIT CONFIG COMMERCE-SERVICE"
curl --request PUT --data 'local' http://localhost:8500/v1/kv/services/commerce/ENV
curl --request PUT --data 'commerce-service' http://localhost:8500/v1/kv/services/commerce/SERVICE_NAME

### MongoDB
curl --request PUT --data 'mongo' http://localhost:8500/v1/kv/services/commerce/MONGO_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/commerce/MONGO_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/commerce/MONGO_HOST
curl --request PUT --data '27017' http://localhost:8500/v1/kv/services/commerce/MONGO_PORT
curl --request PUT --data 'commerce' http://localhost:8500/v1/kv/services/commerce/MONGO_DATABASE_NAME

### rabbitMQ
curl --request PUT --data 'rabbitmq' http://localhost:8500/v1/kv/services/commerce/RABBITMQ_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/commerce/RABBITMQ_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/commerce/RABBITMQ_HOST # Change value localhost to container-name if you want to run service via docker-compose
curl --request PUT --data '5672' http://localhost:8500/v1/kv/services/commerce/RABBITMQ_PORT

### GRPC SERVICES
#PRODUCT
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/commerce/PRODUCT_SERVICE_RPC_HOST # Change value localhost to container-name if you want to run service via docker-compose
curl --request PUT --data '50051' http://localhost:8500/v1/kv/services/commerce/PRODUCT_SERVICE_RPC_PORT # Change value localhost to container-name if you want to run service via docker-compose
#USER
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/commerce/USER_SERVICE_RPC_HOST # Change value localhost to container-name if you want to run service via docker-compose
curl --request PUT --data '50052' http://localhost:8500/v1/kv/services/commerce/USER_SERVICE_RPC_PORT # Change value localhost to container-name if you want to run service via docker-compose

### GRPC SERVER
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/commerce/RPC_HOST
curl --request PUT --data '50054' http://localhost:8500/v1/kv/services/commerce/RPC_PORT

### TELEMETRY
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/commerce/JAEGER_TELEMETRY_HOST
curl --request PUT --data '14268' http://localhost:8500/v1/kv/services/commerce/JAEGER_TELEMETRY_PORT

##########################################################################################################################################################################################################################
echo "INIT EVENT-STORE SERVICE"
curl --request PUT --data 'local' http://localhost:8500/v1/kv/services/event-store/ENV
curl --request PUT --data 'commerce-service' http://localhost:8500/v1/kv/services/event-store/SERVICE_NAME

### MongoDB
curl --request PUT --data 'mongo' http://localhost:8500/v1/kv/services/event-store/MONGO_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/event-store/MONGO_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/event-store/MONGO_HOST
curl --request PUT --data '27017' http://localhost:8500/v1/kv/services/event-store/MONGO_PORT
curl --request PUT --data 'event-store' http://localhost:8500/v1/kv/services/event-store/MONGO_DATABASE_NAME

### rabbitMQ
curl --request PUT --data 'rabbitmq' http://localhost:8500/v1/kv/services/event-store/RABBITMQ_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/event-store/RABBITMQ_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/event-store/RABBITMQ_HOST # Change value localhost to container-name if you want to run service via docker-compose
curl --request PUT --data '5672' http://localhost:8500/v1/kv/services/event-store/RABBITMQ_PORT

### TELEMETRY
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/event-store/JAEGER_TELEMETRY_HOST
curl --request PUT --data '14268' http://localhost:8500/v1/kv/services/event-store/JAEGER_TELEMETRY_PORT

##########################################################################################################################################################################################################################
echo "INIT NOTIFICATION SERVICE"
curl --request PUT --data 'local' http://localhost:8500/v1/kv/services/notification/ENV
curl --request PUT --data 'commerce-service' http://localhost:8500/v1/kv/services/notification/SERVICE_NAME

### MongoDB
curl --request PUT --data 'mongo' http://localhost:8500/v1/kv/services/notification/MONGO_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/notification/MONGO_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/notification/MONGO_HOST
curl --request PUT --data '27017' http://localhost:8500/v1/kv/services/notification/MONGO_PORT
curl --request PUT --data 'notification' http://localhost:8500/v1/kv/services/notification/MONGO_DATABASE_NAME

### rabbitMQ
curl --request PUT --data 'rabbitmq' http://localhost:8500/v1/kv/services/notification/RABBITMQ_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/notification/RABBITMQ_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/notification/RABBITMQ_HOST # Change value localhost to container-name if you want to run service via docker-compose
curl --request PUT --data '5672' http://localhost:8500/v1/kv/services/notification/RABBITMQ_PORT

### SMTP
curl --request PUT --data 'ecommerce@email.com' http://localhost:8500/v1/kv/services/notification/SMTP_SENDER_EMAIL
curl --request PUT --data '' http://localhost:8500/v1/kv/services/notification/SMTP_USERNAME
curl --request PUT --data '' http://localhost:8500/v1/kv/services/notification/SMTP_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/notification/SMTP_HOST
curl --request PUT --data '1025' http://localhost:8500/v1/kv/services/notification/SMTP_PORT

### TELEMETRY
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/notification/JAEGER_TELEMETRY_HOST
curl --request PUT --data '14268' http://localhost:8500/v1/kv/services/notification/JAEGER_TELEMETRY_PORT

##########################################################################################################################################################################################################################
echo "INIT PAYMENT SERVICE"


##########################################################################################################################################################################################################################
echo "INIT PRODUCT SERVICE"
curl --request PUT --data 'local' http://localhost:8500/v1/kv/services/product/ENV
curl --request PUT --data 'product-service' http://localhost:8500/v1/kv/services/product/SERVICE_NAME

### POSTGRES
curl --request PUT --data 'postgres' http://localhost:8500/v1/kv/services/product/POSTGRES_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/product/POSTGRES_PASSWORD
curl --request PUT --data 'disable' http://localhost:8500/v1/kv/services/product/POSTGRES_SSL_MODE
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/product/POSTGRES_HOST
curl --request PUT --data '5432' http://localhost:8500/v1/kv/services/product/POSTGRES_PORT
curl --request PUT --data 'products' http://localhost:8500/v1/kv/services/product/POSTGRES_DATABASE_NAME

### ELASTICSEARCH
curl --request PUT --data '' http://localhost:8500/v1/kv/services/product/ELASTICSEARCH_USERNAME
curl --request PUT --data '' http://localhost:8500/v1/kv/services/product/ELASTICSEARCH_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/product/ELASTICSEARCH_HOST
curl --request PUT --data '9200' http://localhost:8500/v1/kv/services/product/ELASTICSEARCH_PORT

### rabbitMQ
curl --request PUT --data 'rabbitmq' http://localhost:8500/v1/kv/services/product/RABBITMQ_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/product/RABBITMQ_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/product/RABBITMQ_HOST # Change value localhost to container-name if you want to run service via docker-compose
curl --request PUT --data '5672' http://localhost:8500/v1/kv/services/product/RABBITMQ_PORT

### GRPC SERVER
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/product/RPC_HOST
curl --request PUT --data '50051' http://localhost:8500/v1/kv/services/product/RPC_PORT

### TELEMETRY
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/product/JAEGER_TELEMETRY_HOST
curl --request PUT --data '14268' http://localhost:8500/v1/kv/services/product/JAEGER_TELEMETRY_PORT

##########################################################################################################################################################################################################################
echo "INIT USER SERVICE"
curl --request PUT --data 'local' http://localhost:8500/v1/kv/services/user/ENV
curl --request PUT --data 'user-service' http://localhost:8500/v1/kv/services/user/SERVICE_NAME

### POSTGRES
curl --request PUT --data 'postgres' http://localhost:8500/v1/kv/services/user/POSTGRES_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/services/user/POSTGRES_PASSWORD
curl --request PUT --data 'disable' http://localhost:8500/v1/kv/services/user/POSTGRES_SSL_MODE
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/user/POSTGRES_HOST
curl --request PUT --data '5432' http://localhost:8500/v1/kv/services/user/POSTGRES_PORT
curl --request PUT --data 'users' http://localhost:8500/v1/kv/services/user/POSTGRES_DATABASE_NAME

### REDIS
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/user/REDIS_HOST
curl --request PUT --data '6379' http://localhost:8500/v1/kv/services/user/REDIS_PORT
curl --request PUT --data '' http://localhost:8500/v1/kv/services/user/REDIS_PASSWORD
curl --request PUT --data '0' http://localhost:8500/v1/kv/services/user/REDIS_DB

### GRPC SERVER
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/user/RPC_HOST
curl --request PUT --data '50052' http://localhost:8500/v1/kv/services/user/RPC_PORT

### JWT ACCESS TOKEN
curl --request PUT --data 'ecommerce-service-v2' http://localhost:8500/v1/kv/services/user/JWT_ACCESS_TOKEN_SECRET
curl --request PUT --data '1h' http://localhost:8500/v1/kv/services/user/JWT_ACCESS_TOKEN_EXPIRATION_TIME

### JWT REFRESH ACCESS TOKEN
curl --request PUT --data 'v2-service-ecommerce' http://localhost:8500/v1/kv/services/user/JWT_REFRESH_TOKEN_SECRET
curl --request PUT --data '2d' http://localhost:8500/v1/kv/services/user/JWT_REFRESH_TOKEN_EXPIRATION_TIME

# ETC
curl --request PUT --data 'http://localhost:4000?access_token=%s?refresh_token=?' http://localhost:8500/v1/kv/services/user/VERIFICATION_USER_LOGIN_URL

### TELEMETRY
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/services/user/JAEGER_TELEMETRY_HOST
curl --request PUT --data '14268' http://localhost:8500/v1/kv/services/user/JAEGER_TELEMETRY_PORT


echo "✅ Done setting key-values."
