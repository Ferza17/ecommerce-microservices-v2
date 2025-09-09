#!/bin/sh

echo "Cleaning generated docs..."
  rm -rf services/commerce-service/src/docs \
           services/event-store-service/docs \
           services/notification-service/docs \
           services/payment-service/docs \
           services/product-service/docs \
           services/user-service/docs
  
echo "Done cleaning docs."


