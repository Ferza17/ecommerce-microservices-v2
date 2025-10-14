package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class WishlistKafkaConsumer {
    @KafkaListener(topics = "snapshot-commerce-wishlist_created", groupId = "commerce-service")
    public void handleSnapshotCommerceWishlistCreated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "confirm-snapshot-commerce-wishlist_created", groupId = "commerce-service")
    public void handleConfirmSnapshotCommerceWishlistCreated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "compensate-snapshot-commerce-wishlist_created", groupId = "commerce-service")
    public void handleCompensateSnapshotCommerceWishlistCreated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "snapshot-commerce-wishlist_updated", groupId = "commerce-service")
    public void handleSnapshotCommerceWishlistUpdated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "confirm-snapshot-commerce-wishlist_updated", groupId = "commerce-service")
    public void handleConfirmSnapshotCommerceWishlistUpdated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "compensate-snapshot-commerce-wishlist_updated", groupId = "commerce-service")
    public void handleCompensateSnapshotCommerceWishlistUpdated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "snapshot-commerce-wishlist_deleted", groupId = "commerce-service")
    public void handleSnapshotCommerceWishlistDeleted(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "confirm-snapshot-commerce-wishlist_deleted", groupId = "commerce-service")
    public void handleConfirmSnapshotCommerceWishlistDeleted(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "compensate-snapshot-commerce-wishlist_deleted", groupId = "commerce-service")
    public void handleCompensateSnapshotCommerceWishlistDeleted(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "dlq-sink-mongo-commerce-wishlists", groupId = "")
    public void handleDlqSinkMongoCommerceWishlists(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

}
