package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

@Component
public class CartKafkaConsumer {
    @Autowired
    private CartUseCase cartUseCase;
    @Autowired
    private Tracer tracer;

    // TODO: Start Span From Kafka Header
    @KafkaListener(topics = "snapshot-commerce-cart_added", groupId = "commerce-service", containerFactory = "kafkaListenerContainerAddToCartFactory")
    public void handleSnapshotCommerceCartCreated(Request.AddToCartRequest message) {
        Span span = this.tracer.spanBuilder("CartKafkaConsumer.handleSnapshotCommerceCartCreated").startSpan();
        try (Scope scope = span.makeCurrent()) {
            System.out.println("📥 Received Kafka message: " + message);
        } catch (Exception e) {
            span.recordException(e);
            throw new RuntimeException(e);
        } finally {
            span.end();
        }
    }

    @KafkaListener(topics = "snapshot-commerce-cart_deleted", groupId = "commerce-service", containerFactory = "kafkaListenerContainerDeleteCartItemByIdFactory")
    public void handleSnapshotCommerceCartDeleted(Request.DeleteCartItemByIdRequest message) {
        Span span = this.tracer.spanBuilder("CartKafkaConsumer.handleSnapshotCommerceCartDeleted").startSpan();
        try (Scope scope = span.makeCurrent()) {
            System.out.println("📥 Received Kafka message: " + message);
        } catch (Exception e) {
            span.recordException(e);
            throw new RuntimeException(e);
        } finally {
            span.end();
        }
    }

//    @KafkaListener(topics = "dlq-sink-mongo-commerce-carts", groupId = "")
//    public void handleDlqSinkMongoCommerceCarts(String message) {
//        Span span = this.tracer.spanBuilder("CartKafkaConsumer.handleDlqSinkMongoCommerceCarts").startSpan();
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
