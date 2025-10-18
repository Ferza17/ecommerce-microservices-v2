package com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound;

import org.apache.kafka.clients.producer.ProducerInterceptor;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;
import org.slf4j.MDC;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.X_REQUEST_ID_CONTEXT_KEY;
import java.util.Map;

public class RequestIDOutboundKafka<K, V> implements ProducerInterceptor<K, V> {
    @Override
    public ProducerRecord<K, V> onSend(ProducerRecord<K, V> record) {
        record.headers().add(X_REQUEST_ID_CONTEXT_KEY, MDC.get(X_REQUEST_ID_CONTEXT_KEY).getBytes());
        return record;
    }

    @Override
    public void onAcknowledgement(RecordMetadata metadata, Exception exception) {

    }

    @Override
    public void close() {

    }

    @Override
    public void configure(Map<String, ?> configs) {

    }
}
