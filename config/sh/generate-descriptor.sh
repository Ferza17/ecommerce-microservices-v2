#!/bin/sh

echo ">>> Generating descriptor.pb files with buf <<<"
buf build --as-file-descriptor-set --output api-gateway/descriptor.pb;
buf build --as-file-descriptor-set --output commerce-service/descriptor.pb;
buf build --as-file-descriptor-set --output event-store-service/descriptor.pb;
buf build --as-file-descriptor-set --output notification-service/descriptor.pb;
buf build --as-file-descriptor-set --output payment-service/descriptor.pb;
buf build --as-file-descriptor-set --output product-service/descriptor.pb;
buf build --as-file-descriptor-set --output user-service/descriptor.pb;
buf build --as-file-descriptor-set --output shipping-service/descriptor.pb;
echo ">>> Done."
