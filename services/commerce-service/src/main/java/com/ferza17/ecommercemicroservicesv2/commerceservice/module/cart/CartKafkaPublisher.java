package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class CartKafkaPublisher {
    private final KafkaTemplate<String, String> kafkaTemplate;

    public CartKafkaPublisher(KafkaTemplate<String, String> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }
}
