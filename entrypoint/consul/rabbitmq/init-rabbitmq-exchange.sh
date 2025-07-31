#!/bin/sh

echo "INIT CONFIG RABBITMQ EXCHANGE"

# EXCHANGE LOCAL
consul kv put local/broker/rabbitmq/EXCHANGE/PRODUCT "product.direct.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/COMMERCE "commerce.direct.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/EVENT "event.direct.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/NOTIFICATION "notification.direct.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/USER "user.direct.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/PAYMENT/DIRECT "payment.direct.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/PAYMENT/DELAYED "payment.delayed.exchange"
consul kv put local/broker/rabbitmq/EXCHANGE/SHIPPING "shipping.direct.exchange"

# EXCHANGE PRODUCTION
consul kv put production/broker/rabbitmq/EXCHANGE/PRODUCT "product.direct.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/COMMERCE "commerce.direct.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/EVENT "event.direct.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/NOTIFICATION "notification.direct.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/USER "user.direct.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/PAYMENT/DIRECT "payment.direct.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/PAYMENT/DELAYED "payment.delayed.exchange"
consul kv put production/broker/rabbitmq/EXCHANGE/SHIPPING "shipping.direct.exchange"

echo "DONE INIT CONFIG RABBITMQ EXCHANGE"
