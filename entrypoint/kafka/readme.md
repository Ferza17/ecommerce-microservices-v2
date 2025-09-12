# Naming Convention for RabbitMQ
---
## SNAPSHOT /  CHANGELOG TOPIC
- naming convention for snapshot / changelog :`<domain>.<type>.snapshot`

### Example : 
- `users.login.snapshot`
- `users.created.snapshot`
- `users.updated.snapshot`
---

## SINK CONNECTOR
- naming convention for snapshot / changelog :`<connector type>.<database provider>.<databases>.<table or collection>`

### Example :
- `sink.pg.users.users`
- `sink.pg.users.roles`
- `sink.pg.products.products`
- `sink.pg.shippings.shippings`
- `sink.pg.shippings.shipping_providers`
- `sink.pg.payments.payments`
- `sink.pg.payments.payment_providers`
- `source.mongo.notification.templates`
---
