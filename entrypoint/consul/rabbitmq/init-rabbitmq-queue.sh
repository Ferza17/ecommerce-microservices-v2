#!/bin/sh

echo "INIT CONFIG RABBITMQ QUEUE"

# QUEUE LOCAL USER
consul kv put local/broker/rabbitmq/QUEUE/USER/CREATED "user.created"
consul kv put local/broker/rabbitmq/QUEUE/USER/UPDATED "user.updated"
consul kv put local/broker/rabbitmq/QUEUE/USER/LOGIN "user.login"
consul kv put local/broker/rabbitmq/QUEUE/USER/LOGOUT "user.logout"

#QUEUE PRODUCTION USER
consul kv put production/broker/rabbitmq/QUEUE/USER/CREATED "user.created"
consul kv put production/broker/rabbitmq/QUEUE/USER/UPDATED "user.updated"
consul kv put production/broker/rabbitmq/QUEUE/USER/LOGIN "user.login"
consul kv put production/broker/rabbitmq/QUEUE/USER/LOGOUT "user.logout"


# QUEUE LOCAL PRODUCT
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/CREATED "product.created"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/UPDATED "product.updated"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/DELETED "product.deleted"

# QUEUE PRODUCTION PRODUCT
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/CREATED "product.created"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/UPDATED "product.updated"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/DELETED "product.deleted"

# QUEUE LOCAL NOTIFICATION
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED "notification.email.otp.created"
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED "notification.email.payment.order.created"

# QUEUE PRODUCTION NOTIFICATION
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED "notification.email.otp.created"
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED "notification.email.payment.order.created"

# QUEUE LOCAL EVENT
consul kv put local/broker/rabbitmq/QUEUE/EVENT/CREATED "event.created"

# QUEUE PRODUCTION EVENT
consul kv put production/broker/rabbitmq/QUEUE/EVENT/CREATED "event.created"

# QUEUE LOCAL COMMERCE
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED "cart.created"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED "cart.updated"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED "cart.deleted"

# QUEUE PRODUCTION COMMERCE
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED "cart.created"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED "cart.updated"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED "cart.deleted"

# QUEUE LOCAL PAYMENT
consul kv put local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED "payment.order.created"
consul kv put local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED "payment.order.delayed.cancelled"

# QUEUE PRODUCTION PAYMENT
consul kv put production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED "payment.order.created"
consul kv put production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED "payment.order.delayed.cancelled"

#QUEUE LOCAL SHIPPING
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/CREATED "shipping.created"
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/UPDATED "shipping.updated"


#QUEUE PRODUCTION SHIPPING
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/CREATED "shipping.created"
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/UPDATED "shipping.updated"

echo "DONE INIT CONFIG RABBITMQ QUEUE"
