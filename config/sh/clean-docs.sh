#!/bin/sh

echo "Cleaning generated docs..."
  rm -rf commerce-service/src/docs \
           event-store-service/docs \
           notification-service/docs \
           payment-service/docs \
           product-service/docs \
           user-service/docs       
  
echo "Done cleaning docs."


