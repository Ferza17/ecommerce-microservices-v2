package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class WishlistKafkaPublisher {
    private final KafkaTemplate<String, String> kafkaTemplate;
    private final OpenTelemetrySdk openTelemetrySdk;


    public WishlistKafkaPublisher(KafkaTemplate<String, String> kafkaTemplate, OpenTelemetrySdk openTelemetrySdk) {
        this.kafkaTemplate = kafkaTemplate;
        this.openTelemetrySdk = openTelemetrySdk;
    }

    public void publishSinkMongoCommerceWishlistItem(Model.WishlistItem wishlistItem) {
        //TODO: Implement Me
    }
}
