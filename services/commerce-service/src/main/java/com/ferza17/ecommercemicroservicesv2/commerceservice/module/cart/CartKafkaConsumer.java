package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class CartKafkaConsumer {
    @KafkaListener(topics = "snapshot-commerce-cart_created", groupId = "commerce-service")
    public void handleSnapshotCommerceCartCreated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "confirm-snapshot-commerce-cart_created", groupId = "commerce-service")
    public void handleConfirmSnapshotCommerceCartCreated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "compensate-snapshot-commerce-cart_created", groupId = "commerce-service")
    public void handleCompensateSnapshotCommerceCartCreated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "snapshot-commerce-cart_updated", groupId = "commerce-service")
    public void handleSnapshotCommerceCartUpdated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "confirm-snapshot-commerce-cart_updated", groupId = "commerce-service")
    public void handleConfirmSnapshotCommerceCartUpdated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "compensate-snapshot-commerce-cart_updated", groupId = "commerce-service")
    public void handleCompensateSnapshotCommerceCartUpdated(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "snapshot-commerce-cart_deleted", groupId = "commerce-service")
    public void handleSnapshotCommerceCartDeleted(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "confirm-snapshot-commerce-cart_deleted", groupId = "commerce-service")
    public void handleConfirmSnapshotCommerceCartDeleted(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "compensate-snapshot-commerce-cart_deleted", groupId = "commerce-service")
    public void handleCompensateSnapshotCommerceCartDeleted(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

    @KafkaListener(topics = "dlq-sink-mongo-commerce-carts", groupId = "")
    public void handleDlqSinkMongoCommerceCarts(String message) {
        System.out.println("📥 Received Kafka message: " + message);
    }

}
