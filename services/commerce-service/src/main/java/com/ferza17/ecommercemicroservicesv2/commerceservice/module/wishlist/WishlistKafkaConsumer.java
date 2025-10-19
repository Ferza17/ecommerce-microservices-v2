package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

@Component
public class WishlistKafkaConsumer {
    @Autowired
    private WishlistUseCase wishlistUseCase;
    @Autowired
    private Tracer tracer;

    @KafkaListener(
            topics = "snapshot-commerce-wishlist_added",
            groupId = "commerce-service",
            containerFactory = "kafkaListenerContainerAddToWishlistFactory",
            errorHandler = "KafkaGlobalException"
    )
    public void handleSnapshotCommerceWishlistCreated(
            @Payload Request.AddToWishlistRequest message
    ) {
        Span span = this.tracer.spanBuilder("WishlistKafkaConsumer.handleSnapshotCommerceWishlistCreated").startSpan();
        try (Scope scope = span.makeCurrent()) {
            System.out.println("📥 Received Kafka message: " + message);
        } catch (Exception e) {
            span.recordException(e);
            throw new RuntimeException(e);
        } finally {
            span.end();
        }
    }

    @KafkaListener(
            topics = "snapshot-commerce-wishlist_deleted",
            groupId = "commerce-service",
            containerFactory = "kafkaListenerContainerDeleteWishlistItemByIdFactory",
            errorHandler = "KafkaGlobalException"
    )
    public void handleSnapshotCommerceWishlistDeleted(
            @Payload Request.DeleteWishlistItemByIdRequest message
    ) {
        Span span = this.tracer.spanBuilder("WishlistKafkaConsumer.handleSnapshotCommerceWishlistDeleted").startSpan();
        try (Scope scope = span.makeCurrent()) {
            System.out.println("📥 Received Kafka message: " + message);
        } catch (Exception e) {
            span.recordException(e);
            throw new RuntimeException(e);
        } finally {
            span.end();
        }
    }

//    @KafkaListener(topics = "dlq-sink-mongo-commerce-wishlists", groupId = "")
//    public void handleDlqSinkMongoCommerceWishlists(String message) {
//        Span span = this.tracer.spanBuilder("WishlistKafkaConsumer.handleDlqSinkMongoCommerceWishlists").startSpan();
//        try (Scope scope = span.makeCurrent()) {
//            System.out.println("📥 Received Kafka message: " + message);
//        } catch (Exception e) {
//            span.recordException(e);
//            throw new RuntimeException(e);
//        } finally {
//            span.end();
//        }
//    }

}
