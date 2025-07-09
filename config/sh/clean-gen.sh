#!/bin/sh

echo "Cleaning proto generated files..."
rm -rf api-gateway/model/rpc/gen \
         commerce-service/src/model/rpc/gen \
         event-store-service/model/rpc/gen \
         notification-service/model/rpc/gen \
         payment-service/model/rpc/gen \
         product-service/model/rpc/gen \
         user-service/model/rpc/gen
echo "Done cleaning."
