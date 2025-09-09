#!/bin/sh

echo "Cleaning proto generated files..."
rm -rf services/api-gateway/model/rpc/gen \
         services/commerce-service/src/model/rpc/gen \
         services/event-store-service/model/rpc/gen \
         services/notification-service/model/rpc/gen \
         services/payment-service/model/rpc/gen \
         services/product-service/model/rpc/gen \
         services/user-service/model/rpc/gen
echo "Done cleaning."
