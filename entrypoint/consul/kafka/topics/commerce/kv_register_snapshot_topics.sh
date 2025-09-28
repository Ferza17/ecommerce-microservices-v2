#!/bin/sh

echo "INIT CONFIG KAFKA NAMESPACE COMMERCE TOPICS"

## LOCAL
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_CREATED "snapshot-commerce-cart_created"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CONFIRM/CART_CREATED "confirm-snapshot-commerce-cart_created"
consul kv put local/broker/kafka/TOPICS/COMMERCE/COMPENSATE/CART_CREATED "compensate-snapshot-commerce-cart_created"

consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_UPDATED "snapshot-commerce-cart_updated"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CONFIRM/CART_UPDATED "confirm-snapshot-commerce-cart_updated"
consul kv put local/broker/kafka/TOPICS/COMMERCE/COMPENSATE/CART_UPDATED "compensate-snapshot-commerce-cart_updated"

consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_DELETED "snapshot-commerce-cart_deleted"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CONFIRM/CART_DELETED "confirm-snapshot-commerce-cart_deleted"
consul kv put local/broker/kafka/TOPICS/COMMERCE/COMPENSATE/CART_DELETED "compensate-snapshot-commerce-cart_deleted"

consul kv put local/broker/kafka/TOPICS/COMMERCE/WISHLIST_CREATED "snapshot-commerce-wishlist_created"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CONFIRM/WISHLIST_CREATED "confirm-snapshot-commerce-wishlist_created"
consul kv put local/broker/kafka/TOPICS/COMMERCE/COMPENSATE/WISHLIST_CREATED "compensate-snapshot-commerce-wishlist_created"

consul kv put local/broker/kafka/TOPICS/COMMERCE/WISHLIST_UPDATED "snapshot-commerce-wishlist_updated"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CONFIRM/WISHLIST_UPDATED "confirm-snapshot-commerce-wishlist_updated"
consul kv put local/broker/kafka/TOPICS/COMMERCE/COMPENSATE/WISHLIST_UPDATED "compensate-snapshot-commerce-wishlist_updated"

consul kv put local/broker/kafka/TOPICS/COMMERCE/WISHLIST_DELETED "snapshot-commerce-wishlist_deleted"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CONFIRM/WISHLIST_DELETED "confirm-snapshot-commerce-wishlist_deleted"
consul kv put local/broker/kafka/TOPICS/COMMERCE/COMPENSATE/WISHLIST_DELETED "compensate-snapshot-commerce-wishlist_deleted"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_CREATED "snapshot-commerce-cart_created"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CONFIRM/CART_CREATED "confirm-snapshot-commerce-cart_created"
consul kv put production/broker/kafka/TOPICS/COMMERCE/COMPENSATE/CART_CREATED "compensate-snapshot-commerce-cart_created"

consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_UPDATED "snapshot-commerce-cart_updated"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CONFIRM/CART_UPDATED "confirm-snapshot-commerce-cart_updated"
consul kv put production/broker/kafka/TOPICS/COMMERCE/COMPENSATE/CART_UPDATED "compensate-snapshot-commerce-cart_updated"

consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_DELETED "snapshot-commerce-cart_deleted"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CONFIRM/CART_DELETED "confirm-snapshot-commerce-cart_deleted"
consul kv put production/broker/kafka/TOPICS/COMMERCE/COMPENSATE/CART_DELETED "compensate-snapshot-commerce-cart_deleted"

consul kv put production/broker/kafka/TOPICS/COMMERCE/WISHLIST_CREATED "snapshot-commerce-wishlist_created"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CONFIRM/WISHLIST_CREATED "confirm-snapshot-commerce-wishlist_created"
consul kv put production/broker/kafka/TOPICS/COMMERCE/COMPENSATE/WISHLIST_CREATED "compensate-snapshot-commerce-wishlist_created"

consul kv put production/broker/kafka/TOPICS/COMMERCE/WISHLIST_UPDATED "snapshot-commerce-wishlist_updated"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CONFIRM/WISHLIST_UPDATED "confirm-snapshot-commerce-wishlist_updated"
consul kv put production/broker/kafka/TOPICS/COMMERCE/COMPENSATE/WISHLIST_UPDATED "compensate-snapshot-commerce-wishlist_updated"

consul kv put production/broker/kafka/TOPICS/COMMERCE/WISHLIST_DELETED "snapshot-commerce-wishlist_deleted"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CONFIRM/WISHLIST_DELETED "confirm-snapshot-commerce-wishlist_deleted"
consul kv put production/broker/kafka/TOPICS/COMMERCE/COMPENSATE/WISHLIST_DELETED "compensate-snapshot-commerce-wishlist_deleted"

echo "DONE INIT CONFIG KAFKA NAMESPACE COMMERCE TOPICS"
