package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import io.opentelemetry.sdk.OpenTelemetrySdk;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class WishlistKafkaConsumer {
    private final WishlistUseCase wishlistUseCase;
    private final OpenTelemetrySdk openTelemetrySdk;


    public WishlistKafkaConsumer(WishlistUseCase wishlistUseCase, OpenTelemetrySdk openTelemetrySdk) {
        this.wishlistUseCase = wishlistUseCase;
        this.openTelemetrySdk = openTelemetrySdk;
    }

    @KafkaListener(topics = "snapshot-commerce-wishlist_created", groupId = "commerce-service")
    public void handleSnapshotCommerceWishlistCreated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "snapshot-commerce-wishlist_updated", groupId = "commerce-service")
    public void handleSnapshotCommerceWishlistUpdated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "snapshot-commerce-wishlist_deleted", groupId = "commerce-service")
    public void handleSnapshotCommerceWishlistDeleted(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "dlq-sink-mongo-commerce-wishlists", groupId = "")
    public void handleDlqSinkMongoCommerceWishlists(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

}
