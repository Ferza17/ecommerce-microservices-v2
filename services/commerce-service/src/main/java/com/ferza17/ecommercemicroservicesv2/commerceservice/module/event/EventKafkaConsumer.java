package com.ferza17.ecommercemicroservicesv2.commerceservice.module.event;

import io.opentelemetry.sdk.OpenTelemetrySdk;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class EventKafkaConsumer {
    private final OpenTelemetrySdk openTelemetrySdk;
    private final EventUseCase eventUseCase;

    public EventKafkaConsumer(OpenTelemetrySdk openTelemetrySdk, EventUseCase eventUseCase) {
        this.openTelemetrySdk = openTelemetrySdk;
        this.eventUseCase = eventUseCase;
    }


    @KafkaListener(topics = "dlq-sink-mongo-events-commerce_event_stores", groupId = "commerce-service")
    public void handleDlqSinkMongoEventsCommerceEventStores(String message) {

    }
}
