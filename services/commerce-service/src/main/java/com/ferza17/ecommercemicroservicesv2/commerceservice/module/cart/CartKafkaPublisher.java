package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import io.opentelemetry.sdk.OpenTelemetrySdk;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class CartKafkaPublisher {
    private final KafkaTemplate<String, String> kafkaTemplate;
    private final OpenTelemetrySdk openTelemetrySdk;

    public CartKafkaPublisher(KafkaTemplate<String, String> kafkaTemplate, OpenTelemetrySdk openTelemetrySdk) {
        this.kafkaTemplate = kafkaTemplate;
        this.openTelemetrySdk = openTelemetrySdk;
    }
}
