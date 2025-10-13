package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class WishlistKafkaConsumer {
    @KafkaListener(topics = "snapshot-commerce-create_wishlist_item", groupId = "commerce-service")
    public void handleSnapshotCommerceCreateWishlistItem(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

}
