package com.ferza17.ecommercemicroservicesv2.commerceservice.module.event;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class EventKafkaConsumer {
    @KafkaListener(topics = "dlq-sink-mongo-event-commerce_event_stores", groupId = "commerce-service")
    public void handleDqlSinkMongoEventCommerceEventStores(String message) {

    }

}
