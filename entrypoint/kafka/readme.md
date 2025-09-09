# Naming Convention for RabbitMQ
---
## TOPIC
naming convention :`<topic_owner_name>.<type>`

example : 
- `topic.user_register.queue`
- `topic.user_register.failed.queue`
- `topic.user_login.queue`
- `topic.user_login.failed.queue`
- `direct.payment_service.exchange.payment_service.order_created.queue`
- `direct.payment_service.exchange.payment_service.order_created.failed.queue`
- `delayed.payment_service.exchange.payment_service.order_delayed_cancel.queue`
- `delayed.payment_service.exchange.payment_service.order_delayed_cancel.failed.queue`
- `direct.event_store_service.exchange.event_store_service.append.queue`
- `direct.event_store_service.exchange.event_store_service.append.failed.queue`
- `direct.event_store_service.exchange.api_gateway.append.queue`
