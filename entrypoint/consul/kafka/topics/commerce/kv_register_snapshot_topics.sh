#!/bin/sh

echo "INIT CONFIG KAFKA NAMESPACE COMMERCE TOPICS"

## LOCAL
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_ADDEDD "snapshot-commerce-cart_added"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_DELETED "snapshot-commerce-cart_deleted"

consul kv put local/broker/kafka/TOPICS/COMMERCE/WISHLIST_ADDEDD "snapshot-commerce-wishlist_added"
consul kv put local/broker/kafka/TOPICS/COMMERCE/WISHLIST_DELETED "snapshot-commerce-wishlist_deleted"
## PRODUCTION
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_ADDEDD "snapshot-commerce-cart_added"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_DELETED "snapshot-commerce-cart_deleted"

consul kv put production/broker/kafka/TOPICS/COMMERCE/WISHLIST_ADDEDD "snapshot-commerce-wishlist_added"
consul kv put production/broker/kafka/TOPICS/COMMERCE/WISHLIST_DELETED "snapshot-commerce-wishlist_deleted"
echo "DONE INIT CONFIG KAFKA NAMESPACE COMMERCE TOPICS"
