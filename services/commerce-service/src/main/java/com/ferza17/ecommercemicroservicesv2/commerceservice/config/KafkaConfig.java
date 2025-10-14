package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;
import org.springframework.kafka.core.KafkaAdmin;

@Configuration
public class KafkaConfig {
    @Bean
    public KafkaAdmin.NewTopics setupKafkaTopics() {
        return new KafkaAdmin.NewTopics(
                TopicBuilder.name("snapshot-commerce-cart_created").partitions(3).replicas(1).build(),
                TopicBuilder.name("confirm-snapshot-commerce-cart_created").partitions(3).replicas(1).build(),
                TopicBuilder.name("compensate-snapshot-commerce-cart_created").partitions(3).replicas(1).build(),

                TopicBuilder.name("snapshot-commerce-cart_updated").partitions(3).replicas(1).build(),
                TopicBuilder.name("confirm-snapshot-commerce-cart_updated").partitions(3).replicas(1).build(),
                TopicBuilder.name("compensate-snapshot-commerce-cart_updated").partitions(3).replicas(1).build(),

                TopicBuilder.name("snapshot-commerce-cart_deleted").partitions(3).replicas(1).build(),
                TopicBuilder.name("confirm-snapshot-commerce-cart_deleted").partitions(3).replicas(1).build(),
                TopicBuilder.name("compensate-snapshot-commerce-cart_deleted").partitions(3).replicas(1).build(),

                TopicBuilder.name("snapshot-commerce-wishlist_created").partitions(3).replicas(1).build(),
                TopicBuilder.name("confirm-snapshot-commerce-wishlist_created").partitions(3).replicas(1).build(),
                TopicBuilder.name("compensate-snapshot-commerce-wishlist_created").partitions(3).replicas(1).build(),

                TopicBuilder.name("snapshot-commerce-wishlist_updated").partitions(3).replicas(1).build(),
                TopicBuilder.name("confirm-snapshot-commerce-wishlist_updated").partitions(3).replicas(1).build(),
                TopicBuilder.name("compensate-snapshot-commerce-wishlist_updated").partitions(3).replicas(1).build(),

                TopicBuilder.name("snapshot-commerce-wishlist_deleted").partitions(3).replicas(1).build(),
                TopicBuilder.name("confirm-snapshot-commerce-wishlist_deleted").partitions(3).replicas(1).build(),
                TopicBuilder.name("compensate-snapshot-commerce-wishlist_deleted").partitions(3).replicas(1).build()
        );
    }
}
