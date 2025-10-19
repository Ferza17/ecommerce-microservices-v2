package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound.InboundAuthorization;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound.InboundOpenTelemetry;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound.InboundRequestID;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound.OutboundAuthorization;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound.OutboundOpenTelemetry;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound.OutboundRequestID;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.common.serialization.StringDeserializer;
import org.apache.kafka.common.serialization.StringSerializer;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;
import org.springframework.kafka.core.*;
import org.springframework.kafka.config.ConcurrentKafkaListenerContainerFactory;

import java.util.HashMap;
import java.util.Map;

@Configuration
public class KafkaConfig {
    @Value("${spring.kafka.bootstrap-servers}")
    private String bootstrapServers;
    @Value("${spring.kafka.properties.schema.registry.url}")
    private String schemaRegistryUrl;
    @Value("${spring.application.name}")
    private String applicationName;

    @Bean
    public KafkaAdmin.NewTopics setupKafkaTopics() {
        return new KafkaAdmin.NewTopics(
                TopicBuilder.name("snapshot-commerce-cart_added").partitions(3).replicas(1).build(),
                TopicBuilder.name("snapshot-commerce-cart_deleted").partitions(3).replicas(1).build(),

                TopicBuilder.name("snapshot-commerce-wishlist_added").partitions(3).replicas(1).build(),
                TopicBuilder.name("compensate-snapshot-commerce-wishlist_deleted").partitions(3).replicas(1).build()
        );
    }

    /* ============================================================
    *
    *       PUBLISH CONFIG WITH PROTOBUF SERIALIZER
    *       WITH OUTBOUND HEADERS
    *
    ================================================================ */
    @Bean
    public KafkaTemplate<String, Model.CartItem> kafkaTemplateSinkMongoCart() {
        Map<String, Object> properties = new HashMap<>();
        properties.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers);
        properties.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
        properties.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, "io.confluent.kafka.serializers.protobuf.KafkaProtobufSerializer");
        properties.put("schema.registry.url", this.schemaRegistryUrl);
        properties.put("specific.protobuf.value.type", Model.CartItem.class);
        properties.put(ProducerConfig.INTERCEPTOR_CLASSES_CONFIG, String.format("%s,%s,%s", OutboundAuthorization.class.getName(), OutboundRequestID.class.getName(), OutboundOpenTelemetry.class.getName()));
        return new KafkaTemplate<>(new DefaultKafkaProducerFactory<>(properties));
    }

    @Bean
    public KafkaTemplate<String, Model.WishlistItem> kafkaTemplateSinkMongoWishlist() {
        Map<String, Object> properties = new HashMap<>();
        properties.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers);
        properties.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
        properties.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, "io.confluent.kafka.serializers.protobuf.KafkaProtobufSerializer");
        properties.put("schema.registry.url", this.schemaRegistryUrl);
        properties.put("specific.protobuf.value.type", Model.WishlistItem.class);
        properties.put(ProducerConfig.INTERCEPTOR_CLASSES_CONFIG, String.format("%s,%s,%s", OutboundAuthorization.class.getName(), OutboundRequestID.class.getName(), OutboundOpenTelemetry.class.getName()));
        return new KafkaTemplate<>(new DefaultKafkaProducerFactory<>(properties));
    }

    /* ============================================================
    *
    *       CONSUMER CONFIG WITH PROTOBUF DESERIALIZER
    *       WITH INBOUND HEADERS
    *
    ================================================================ */
    @Bean
    public ConcurrentKafkaListenerContainerFactory<String, Request.AddToCartRequest> kafkaListenerContainerAddToCartFactory() {
        Map<String, Object> properties = new HashMap<>();
        properties.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers);
        properties.put(ConsumerConfig.GROUP_ID_CONFIG, applicationName);
        properties.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        properties.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, "io.confluent.kafka.serializers.protobuf.KafkaProtobufDeserializer");
        properties.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        properties.put("schema.registry.url", this.schemaRegistryUrl);
        properties.put("specific.protobuf.value.type", Request.AddToCartRequest.class);
        properties.put(ConsumerConfig.INTERCEPTOR_CLASSES_CONFIG, String.format("%s,%s,%s", InboundAuthorization.class.getName(), InboundRequestID.class.getName(), InboundOpenTelemetry.class.getName()));
        ConcurrentKafkaListenerContainerFactory<String, Request.AddToCartRequest> factory = new ConcurrentKafkaListenerContainerFactory<>();
        factory.setConsumerFactory(new DefaultKafkaConsumerFactory<>(properties));
        return factory;
    }

    @Bean
    public ConcurrentKafkaListenerContainerFactory<String, Request.DeleteCartItemByIdRequest> kafkaListenerContainerDeleteCartItemByIdFactory() {
        Map<String, Object> properties = new HashMap<>();
        properties.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers);
        properties.put(ConsumerConfig.GROUP_ID_CONFIG, applicationName);
        properties.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        properties.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, "io.confluent.kafka.serializers.protobuf.KafkaProtobufDeserializer");
        properties.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        properties.put("schema.registry.url", this.schemaRegistryUrl);
        properties.put("specific.protobuf.value.type", Request.DeleteCartItemByIdRequest.class);
        properties.put(ConsumerConfig.INTERCEPTOR_CLASSES_CONFIG, String.format("%s,%s,%s", InboundAuthorization.class.getName(), InboundRequestID.class.getName(), InboundOpenTelemetry.class.getName()));
        ConcurrentKafkaListenerContainerFactory<String, Request.DeleteCartItemByIdRequest> factory = new ConcurrentKafkaListenerContainerFactory<>();
        factory.setConsumerFactory(new DefaultKafkaConsumerFactory<>(properties));
        return factory;
    }

    @Bean
    public ConcurrentKafkaListenerContainerFactory<String, Request.AddToWishlistRequest> kafkaListenerContainerAddToWishlistFactory() {
        Map<String, Object> properties = new HashMap<>();
        properties.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers);
        properties.put(ConsumerConfig.GROUP_ID_CONFIG, applicationName);
        properties.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        properties.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, "io.confluent.kafka.serializers.protobuf.KafkaProtobufDeserializer");
        properties.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        properties.put("schema.registry.url", this.schemaRegistryUrl);
        properties.put("specific.protobuf.value.type", Request.AddToWishlistRequest.class);
        properties.put(ConsumerConfig.INTERCEPTOR_CLASSES_CONFIG, String.format("%s,%s,%s", InboundAuthorization.class.getName(), InboundRequestID.class.getName(), InboundOpenTelemetry.class.getName()));
        ConcurrentKafkaListenerContainerFactory<String, Request.AddToWishlistRequest> factory = new ConcurrentKafkaListenerContainerFactory<>();
        factory.setConsumerFactory(new DefaultKafkaConsumerFactory<>(properties));
        return factory;
    }

    @Bean
    public ConcurrentKafkaListenerContainerFactory<String, Request.DeleteWishlistItemByIdRequest> kafkaListenerContainerDeleteWishlistItemByIdFactory() {
        Map<String, Object> properties = new HashMap<>();
        properties.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers);
        properties.put(ConsumerConfig.GROUP_ID_CONFIG, applicationName);
        properties.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        properties.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, "io.confluent.kafka.serializers.protobuf.KafkaProtobufDeserializer");
        properties.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        properties.put("schema.registry.url", this.schemaRegistryUrl);
        properties.put("specific.protobuf.value.type", Request.DeleteWishlistItemByIdRequest.class);
        properties.put(ConsumerConfig.INTERCEPTOR_CLASSES_CONFIG, String.format("%s,%s,%s", InboundAuthorization.class.getName(), InboundRequestID.class.getName(), InboundOpenTelemetry.class.getName()));
        ConcurrentKafkaListenerContainerFactory<String, Request.DeleteWishlistItemByIdRequest> factory = new ConcurrentKafkaListenerContainerFactory<>();
        factory.setConsumerFactory(new DefaultKafkaConsumerFactory<>(properties));
        return factory;
    }

}
