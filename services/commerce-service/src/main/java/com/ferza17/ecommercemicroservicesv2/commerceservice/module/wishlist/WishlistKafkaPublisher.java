package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class WishlistKafkaPublisher {
    private final KafkaTemplate<String, String> kafkaTemplate;

    public WishlistKafkaPublisher(KafkaTemplate<String, String> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    public void publishSinkMongoCommerceWishlistItem(Model.WishlistItem wishlistItem) {
        //TODO: Implement Me
    }
}
