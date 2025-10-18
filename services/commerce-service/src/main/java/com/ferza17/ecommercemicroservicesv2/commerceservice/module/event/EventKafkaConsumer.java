package com.ferza17.ecommercemicroservicesv2.commerceservice.module.event;

import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class EventKafkaConsumer {
    @Autowired
    private EventUseCase eventUseCase;
    @Autowired
    private Tracer tracer;

//    @KafkaListener(topics = "dlq-sink-mongo-events-commerce_event_stores", groupId = "commerce-service")
//    public void handleDlqSinkMongoEventsCommerceEventStores(String message) {
//        Span span = this.tracer.spanBuilder("EventKafkaConsumer.handleDlqSinkMongoEventsCommerceEventStores").startSpan();
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
