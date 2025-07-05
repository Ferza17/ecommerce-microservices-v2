echo "INIT PAYMENT SERVICE"
## LOCAL
consul kv put local/services/payment/ENV 'local'
consul kv put local/services/payment/SERVICE_NAME 'payment-service'
consul kv put local/services/payment/PAYMENT_ORDER_CANCELLED_IN_MS '600000'
consul kv put local/services/payment/RPC_HOST 'localhost'
consul kv put local/services/payment/RPC_PORT '50054'
consul kv put local/services/payment/HTTP_HOST 'localhost'
consul kv put local/services/payment/HTTP_PORT '40054'
consul kv put local/services/payment/METRIC_PORT '30054'
## PRODUCTION
consul kv put production/services/payment/ENV 'production'
consul kv put production/services/payment/SERVICE_NAME 'payment-service'
consul kv put production/services/payment/PAYMENT_ORDER_CANCELLED_IN_MS '600000'
consul kv put production/services/payment/RPC_HOST 'payment-service'
consul kv put production/services/payment/RPC_PORT '50054'
consul kv put production/services/payment/HTTP_HOST 'payment-service'
consul kv put production/services/payment/HTTP_PORT '40054'
consul kv put production/services/payment/METRIC_HTTP_PORT '30054'
echo "DONE INIT PAYMENT SERVICE"
