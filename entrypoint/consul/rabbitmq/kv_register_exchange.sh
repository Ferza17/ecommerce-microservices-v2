#!/bin/sh

echo "INIT CONFIG RABBITMQ EXCHANGE"

# EXCHANGE LOCAL
consul kv put local/broker/rabbitmq/EXCHANGE/PRODUCT "direct.product.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/COMMERCE "direct.commerce.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/EVENT "fanout.event.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/NOTIFICATION "direct.notification.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/USER "direct.user.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/PAYMENT/DIRECT "direct.payment.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/PAYMENT/DELAYED "delayed.payment.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/SHIPPING "direct.shipping.exchange"

# EXCHANGE PRODUCTION
consul kv put production/broker/rabbitmq/EXCHANGE/PRODUCT "direct.product.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/COMMERCE "direct.commerce.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/EVENT "fanout.event.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/NOTIFICATION "direct.notification.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/USER "direct.user.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/PAYMENT/DIRECT "direct.payment.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/PAYMENT/DELAYED "delayed.payment.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/SHIPPING "direct.shipping.exchange"

echo "DONE INIT CONFIG RABBITMQ EXCHANGE"
