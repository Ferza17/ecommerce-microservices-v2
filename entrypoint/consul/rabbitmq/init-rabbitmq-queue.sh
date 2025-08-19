#!/bin/sh

echo "INIT CONFIG RABBITMQ QUEUE"

# QUEUE LOCAL USER
consul kv put local/broker/rabbitmq/QUEUE/USER/CREATED "user.created"
consul kv put local/broker/rabbitmq/QUEUE/USER/CREATED/FAILED "user.created.failed"
consul kv put local/broker/rabbitmq/QUEUE/USER/UPDATED "user.updated"
consul kv put local/broker/rabbitmq/QUEUE/USER/UPDATED/FAILED "user.updated.failed"
consul kv put local/broker/rabbitmq/QUEUE/USER/LOGIN "user.login"
consul kv put local/broker/rabbitmq/QUEUE/USER/LOGOUT "user.logout"

#QUEUE PRODUCTION USER
consul kv put production/broker/rabbitmq/QUEUE/USER/CREATED "user.created"
consul kv put production/broker/rabbitmq/QUEUE/USER/CREATED/FAILED "user.created.failed"
consul kv put production/broker/rabbitmq/QUEUE/USER/UPDATED "user.updated"
consul kv put production/broker/rabbitmq/QUEUE/USER/UPDATED/FAILED "user.updated.failed"
consul kv put production/broker/rabbitmq/QUEUE/USER/LOGIN "user.login"
consul kv put production/broker/rabbitmq/QUEUE/USER/LOGOUT "user.logout"

# QUEUE LOCAL PRODUCT
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/CREATED "product.created"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/CREATED/FAILED "product.created.failed"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/UPDATED "product.updated"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/UPDATED/FAILED "product.updated.failed"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/DELETED "product.deleted"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/DELETED/FAILED "product.deleted.failed"

# QUEUE PRODUCTION PRODUCT
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/CREATED "product.created"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/CREATED/FAILED "product.created.failed"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/UPDATED "product.updated"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/UPDATED/FAILED "product.updated.failed"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/DELETED "product.deleted"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/DELETED/FAILED "product.deleted.failed"

# QUEUE LOCAL NOTIFICATION
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED "notification.email.otp.created"
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED/FAILED "notification.email.otp.created.failed"
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED "notification.email.payment.order.created"
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED/FAILED "notification.email.payment.order.created.failed"

# QUEUE PRODUCTION NOTIFICATION
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED "notification.email.otp.created"
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED/FAILED "notification.email.otp.created.failed"
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED "notification.email.payment.order.created"
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED/FAILED "notification.email.payment.order.created.failed"

# QUEUE LOCAL EVENT
consul kv put local/broker/rabbitmq/QUEUE/EVENT/CREATED "event.created"

# QUEUE PRODUCTION EVENT
consul kv put production/broker/rabbitmq/QUEUE/EVENT/CREATED "event.created"

# QUEUE LOCAL COMMERCE
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED "cart.created"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED/FAILED "cart.created.failed"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED "cart.updated"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED/FAILED "cart.updated.failed"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED "cart.deleted"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED/FAILED "cart.deleted.failed"

# QUEUE PRODUCTION COMMERCE
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED "cart.created"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED/FAILED "cart.created.failed"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED "cart.updated"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED/FAILED "cart.updated.failed"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED "cart.deleted"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED/FAILED "cart.deleted.failed"

# QUEUE LOCAL PAYMENT
consul kv put local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED "payment.order.created"
consul kv put local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED/FAILED "payment.order.created.failed"
consul kv put local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED "payment.order.delayed.cancelled"

# QUEUE PRODUCTION PAYMENT
consul kv put production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED "payment.order.created"
consul kv put production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED/FAILED "payment.order.created.failed"
consul kv put production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED "payment.order.delayed.cancelled"

#QUEUE LOCAL SHIPPING
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/CREATED "shipping.created"
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/CREATED/FAILED "shipping.created.failed"
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/UPDATED "shipping.updated"
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/UPDATED/FAILED "shipping.updated.failed"

#QUEUE PRODUCTION SHIPPING
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/CREATED "shipping.created"
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/CREATED/FAILED "shipping.created.failed"
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/UPDATED "shipping.updated"
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/UPDATED/FAILED "shipping.updated.failed"

echo "DONE INIT CONFIG RABBITMQ QUEUE"
