package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;
import org.springframework.kafka.core.KafkaAdmin;

@Configuration
public class KafkaConfig {
    @Bean
    public KafkaAdmin.NewTopics snapshotCommerceCreateCartItem() {
        return new KafkaAdmin.NewTopics(
                TopicBuilder.name("snapshot-commerce-create_cart_item").partitions(3).replicas(1).build(),
                TopicBuilder.name("dlq-sink-mongo-event-commerce_event_stores").partitions(3).replicas(1).build(),
                TopicBuilder.name("snapshot-commerce-create_wishlist_item").partitions(3).replicas(1).build()
        );
    }
}
