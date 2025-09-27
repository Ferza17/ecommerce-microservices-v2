#!/bin/sh

echo ">>> Generating descriptor.pb files with buf <<<"
buf build --as-file-descriptor-set --output services/api-gateway/descriptor.pb;
buf build --as-file-descriptor-set --output services/commerce-service/descriptor.pb;
buf build --as-file-descriptor-set --output services/notification-service/descriptor.pb;
buf build --as-file-descriptor-set --output services/payment-service/descriptor.pb;
buf build --as-file-descriptor-set --output services/product-service/descriptor.pb;
buf build --as-file-descriptor-set --output services/user-service/descriptor.pb;
buf build --as-file-descriptor-set --output services/shipping-service/descriptor.pb;
echo ">>> Done."
