#!/bin/sh

echo "INIT CONFIG RABBITMQ QUEUE"

# QUEUE LOCAL USER
consul kv put local/broker/rabbitmq/QUEUE/USER/CREATED "direct.user.exchange.user.user_created.queue"
consul kv put local/broker/rabbitmq/QUEUE/USER/CREATED/FAILED "direct.user.exchange.user.user_created.failed.queue"
consul kv put local/broker/rabbitmq/QUEUE/USER/UPDATED "direct.user.exchange.user.user_updated.queue"
consul kv put local/broker/rabbitmq/QUEUE/USER/UPDATED/FAILED "direct.user.exchange.user.user_updated.failed.queue"
consul kv put local/broker/rabbitmq/QUEUE/USER/LOGIN "direct.user.exchange.user.user_login.queue"
consul kv put local/broker/rabbitmq/QUEUE/USER/LOGIN/FAILED "direct.user.exchange.user.user_login.failed.queue"
consul kv put local/broker/rabbitmq/QUEUE/USER/LOGOUT "direct.user.exchange.user.user_logout.queue"
consul kv put local/broker/rabbitmq/QUEUE/USER/LOGOUT/FAILED "direct.user.exchange.user.user_logout.failed.queue"

#QUEUE PRODUCTION USER
consul kv put production/broker/rabbitmq/QUEUE/USER/CREATED "direct.user.exchange.user.user_created.queue"
consul kv put production/broker/rabbitmq/QUEUE/USER/CREATED/FAILED "direct.user.exchange.user.user_created.failed.queue"
consul kv put production/broker/rabbitmq/QUEUE/USER/UPDATED "direct.user.exchange.user.user_updated.queue"
consul kv put production/broker/rabbitmq/QUEUE/USER/UPDATED/FAILED "direct.user.exchange.user.user_updated.failed.queue"
consul kv put production/broker/rabbitmq/QUEUE/USER/LOGIN "direct.user.exchange.user.user_login.queue"
consul kv put production/broker/rabbitmq/QUEUE/USER/LOGIN/FAILED "direct.user.exchange.user.user_login.failed.queue"
consul kv put production/broker/rabbitmq/QUEUE/USER/LOGOUT "direct.user.exchange.user.user_logout.queue"
consul kv put production/broker/rabbitmq/QUEUE/USER/LOGOUT/FAILED "direct.user.exchange.user.user_logout.failed.queue"

# QUEUE LOCAL PRODUCT
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/CREATED "direct.product.exchange.product.product_created.queue"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/CREATED/FAILED "direct.product.exchange.product.product_created.failed.queue"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/UPDATED "direct.product.exchange.product.product_updated.queue"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/UPDATED/FAILED "direct.product.exchange.product.product_updated.failed.queue"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/DELETED "direct.product.exchange.product.product_deleted.queue"
consul kv put local/broker/rabbitmq/QUEUE/PRODUCT/DELETED/FAILED "direct.product.exchange.product.product_deleted.failed.queue"

# QUEUE PRODUCTION PRODUCT
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/CREATED "direct.product.exchange.product.product_created.queue"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/CREATED/FAILED "direct.product.exchange.product.product_created.failed.queue"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/UPDATED "direct.product.exchange.product.product_updated.queue"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/UPDATED/FAILED "direct.product.exchange.product.product_updated.failed.queue"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/DELETED "direct.product.exchange.product.product_deleted.queue"
consul kv put production/broker/rabbitmq/QUEUE/PRODUCT/DELETED/FAILED "direct.product.exchange.product.product_deleted.failed.queue"

# QUEUE LOCAL NOTIFICATION
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED "direct.notification.exchange.notification.notification_email_otp_created.queue"
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED/FAILED "direct.notification.exchange.notification.notification_email_otp_created.failed.queue"
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED "direct.notification.exchange.notification.notification_email_payment_order_created.queue"
consul kv put local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED/FAILED "direct.notification.exchange.notification.notification_email_payment_order_created.failed.queue"

# QUEUE PRODUCTION NOTIFICATION
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED "direct.notification.exchange.notification.notification_email_otp_created.queue"
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED/FAILED "direct.notification.exchange.notification.notification_email_otp_created.failed.queue"
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED "direct.notification.exchange.notification.notification_email_payment_order_created.queue"
consul kv put production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/PAYMENT/ORDER/CREATED/FAILED "direct.notification.exchange.notification.notification_email_payment_order_created.failed.queue"

# QUEUE LOCAL EVENT
consul kv put local/broker/rabbitmq/QUEUE/EVENT/EVENT/EVENT-CREATED "fanout.event.exchange.event.event_created.queue"
consul kv put local/broker/rabbitmq/QUEUE/EVENT/API-GATEWAY/EVENT-CREATED "fanout.event.exchange.api_gateway.event_created.queue"

# QUEUE PRODUCTION EVENT
consul kv put production/broker/rabbitmq/QUEUE/EVENT/EVENT/EVENT-CREATED "fanout.event.exchange.event.event_created.queue"
consul kv put production/broker/rabbitmq/QUEUE/EVENT/API-GATEWAY/EVENT-CREATED "fanout.event.exchange.api_gateway.event_created.queue"


# QUEUE LOCAL COMMERCE
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED "direct.commerce.exchange.commerce.cart_created.queue"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED/FAILED "direct.commerce.exchange.commerce.cart_created.failed.queue"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED "direct.commerce.exchange.commerce.cart_updated.queue"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED/FAILED "direct.commerce.exchange.commerce.cart_updated.queue.failed"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED "direct.commerce.exchange.commerce.cart_deleted.queue"
consul kv put local/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED/FAILED "direct.commerce.exchange.commerce.cart_deleted.failed.queue"

# QUEUE PRODUCTION COMMERCE
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED "direct.commerce.exchange.commerce.cart_created.queue"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED/FAILED "direct.commerce.exchange.commerce.cart_created.failed.queue"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED "direct.commerce.exchange.commerce.cart_updated.queue"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED/FAILED "direct.commerce.exchange.commerce.cart_updated.queue.failed"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED "direct.commerce.exchange.commerce.cart_deleted.queue"
consul kv put production/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED/FAILED "direct.commerce.exchange.commerce.cart_deleted.failed.queue"

# QUEUE LOCAL PAYMENT
consul kv put local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED "direct.payment.exchange.payment.payment_order_created.queue"
consul kv put local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED/FAILED "direct.payment.exchange.payment.payment_order_created.failed.queue"
consul kv put local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED "delayed.payment.exchange.payment.payment_order_created.queue"
consul kv put local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED/FAILED "delayed.payment.exchange.payment.payment_order_created.failed.queue"


# QUEUE PRODUCTION PAYMENT
consul kv put production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED "direct.payment.exchange.payment.payment_order_created.queue"
consul kv put production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED/FAILED "direct.payment.exchange.payment.payment_order_created.failed.queue"
consul kv put production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED "delayed.payment.exchange.payment.payment_order_created.queue"
consul kv put production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED/FAILED "delayed.payment.exchange.payment.payment_order_created.failed.queue"


#QUEUE LOCAL SHIPPING
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/CREATED "direct.shipping.exchange.shipping.shipping_created.queue"
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/CREATED/FAILED "direct.shipping.exchange.shipping.shipping_created.failed.queue"
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/UPDATED "direct.shipping.exchange.shipping.shipping_updated.queue"
consul kv put local/broker/rabbitmq/QUEUE/SHIPPING/UPDATED/FAILED "direct.shipping.exchange.shipping.shipping_updated.failed.queue"

#QUEUE PRODUCTION SHIPPING
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/CREATED "direct.shipping.exchange.shipping.shipping_created.queue"
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/CREATED/FAILED "direct.shipping.exchange.shipping.shipping_created.failed.queue"
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/UPDATED "direct.shipping.exchange.shipping.shipping_updated.queue"
consul kv put production/broker/rabbitmq/QUEUE/SHIPPING/UPDATED/FAILED "direct.shipping.exchange.shipping.shipping_updated.failed.queue"

echo "DONE INIT CONFIG RABBITMQ QUEUE"
