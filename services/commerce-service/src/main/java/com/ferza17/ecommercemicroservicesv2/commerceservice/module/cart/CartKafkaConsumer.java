package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class CartKafkaConsumer {
    @KafkaListener(topics = "snapshot-commerce-create_cart_item", groupId = "commerce-service")
    public void handleSnapshotCommerceCreateCartItem(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }
}
