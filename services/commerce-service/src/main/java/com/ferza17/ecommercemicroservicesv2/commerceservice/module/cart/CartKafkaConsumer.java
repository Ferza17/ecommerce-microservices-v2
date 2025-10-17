package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class CartKafkaConsumer {
    @Autowired
    private CartUseCase cartUseCase;
    @Autowired
    private Tracer tracer;

    // TODO: Start Span From Kafka Header
    @KafkaListener(topics = "snapshot-commerce-cart_created", groupId = "commerce-service")
    public void handleSnapshotCommerceCartCreated(String message) {
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

    @KafkaListener(topics = "snapshot-commerce-cart_updated", groupId = "commerce-service")
    public void handleSnapshotCommerceCartUpdated(String message) {
        Span span = this.tracer.spanBuilder("CartKafkaConsumer.handleSnapshotCommerceCartUpdated").startSpan();
        try (Scope scope = span.makeCurrent()) {
            System.out.println("📥 Received Kafka message: " + message);
        } catch (Exception e) {
            span.recordException(e);
            throw new RuntimeException(e);
        } finally {
            span.end();
        }
    }

    @KafkaListener(topics = "snapshot-commerce-cart_deleted", groupId = "commerce-service")
    public void handleSnapshotCommerceCartDeleted(String message) {
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

    @KafkaListener(topics = "dlq-sink-mongo-commerce-carts", groupId = "")
    public void handleDlqSinkMongoCommerceCarts(String message) {
        Span span = this.tracer.spanBuilder("CartKafkaConsumer.handleDlqSinkMongoCommerceCarts").startSpan();
        try (Scope scope = span.makeCurrent()) {
            System.out.println("📥 Received Kafka message: " + message);
        } catch (Exception e) {
            span.recordException(e);
            throw new RuntimeException(e);
        } finally {
            span.end();
        }
    }

}
