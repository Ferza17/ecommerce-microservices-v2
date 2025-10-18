package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class WishlistKafkaPublisher {
//    private final KafkaTemplate<String, String> kafkaTemplate;
//    @Autowired
//    private Tracer tracer;
//
//
//    public WishlistKafkaPublisher(KafkaTemplate<String, String> kafkaTemplate) {
//        this.kafkaTemplate = kafkaTemplate;
//    }
//
//    public void publishSinkMongoCommerceWishlistItem(Model.WishlistItem wishlistItem) {
//        Span span = this.tracer.spanBuilder("WishlistKafkaPublisher.publishSinkMongoCommerceWishlistItem").startSpan();
//        try (Scope scope = span.makeCurrent()) {
//            // TODO: Implement Me
//        } catch (Exception e) {
//            span.recordException(e);
//            throw new RuntimeException(e);
//        } finally {
//            span.end();
//        }
//    }
}
