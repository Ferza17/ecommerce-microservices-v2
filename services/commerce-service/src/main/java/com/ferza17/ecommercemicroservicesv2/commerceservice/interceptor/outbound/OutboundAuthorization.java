package com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound;

import io.grpc.*;
import org.apache.kafka.clients.producer.ProducerInterceptor;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;
import org.slf4j.MDC;
import org.springframework.http.HttpRequest;
import org.springframework.http.client.ClientHttpRequestExecution;
import org.springframework.http.client.ClientHttpRequestInterceptor;
import org.springframework.http.client.ClientHttpResponse;
import org.springframework.stereotype.Component;

import java.io.IOException;
import java.util.Map;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.AUTHORIZATION_CONTEXT_KEY;
import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.AUTHORIZATION_METADATA;

@Component
public class OutboundAuthorization<K, V> implements ClientInterceptor, ClientHttpRequestInterceptor, ProducerInterceptor<K, V> {
    /*===============================
     *
     *              GRPC
     *
     * ==============================*/

    @Override
    public <ReqT, RespT> ClientCall<ReqT, RespT> interceptCall(MethodDescriptor<ReqT, RespT> method, CallOptions callOptions, Channel next) {
        try {
            return new ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(
                    next.newCall(method, callOptions)) {

                @Override
                public void start(Listener<RespT> responseListener, Metadata headers) {
                    headers.put(AUTHORIZATION_METADATA, MDC.get(AUTHORIZATION_CONTEXT_KEY));
                    super.start(responseListener, headers);
                }
            };
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    /*===============================
     *
     *              REST
     *
     * ==============================*/
    @Override
    public ClientHttpResponse intercept(HttpRequest request, byte[] body, ClientHttpRequestExecution execution) throws IOException {
        request.getHeaders().add(AUTHORIZATION_CONTEXT_KEY, MDC.get(AUTHORIZATION_CONTEXT_KEY));
        return execution.execute(request, body);
    }

    /*===============================
     *
     *              KAFKA
     *
     * ==============================*/
    @Override
    public ProducerRecord<K, V> onSend(ProducerRecord<K, V> record) {
        record.headers().add(AUTHORIZATION_CONTEXT_KEY, MDC.get(AUTHORIZATION_CONTEXT_KEY).getBytes());
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