package com.ferza17.ecommercemicroservicesv2.commerceservice.module.event;

import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class EventKafkaPublisher {
    private final KafkaTemplate<String, String> kafkaTemplate;

    public EventKafkaPublisher(KafkaTemplate<String, String> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }
}
