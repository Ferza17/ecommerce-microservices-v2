package com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound;

import io.grpc.*;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.apache.kafka.clients.consumer.ConsumerInterceptor;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.OffsetAndMetadata;
import org.apache.kafka.common.TopicPartition;
import org.slf4j.MDC;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.grpc.server.GlobalServerInterceptor;
import org.springframework.lang.Nullable;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import java.util.Map;
import java.util.UUID;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.X_REQUEST_ID_CONTEXT_KEY;
import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.X_REQUEST_ID_METADATA;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE + 1)
@Component
public class InboundRequestID<K, V> implements ServerInterceptor, HandlerInterceptor, ConsumerInterceptor<K, V> {
    /*===============================
     *
     *              GRPC
     *
     * ==============================*/
    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(
            ServerCall<ReqT, RespT> call,
            Metadata metadata,
            ServerCallHandler<ReqT, RespT> next
    ) {
        try {
            String requestId = metadata.get(X_REQUEST_ID_METADATA);
            if (requestId == null || requestId.isBlank()) {
                requestId = UUID.randomUUID().toString();
            }
            MDC.put(X_REQUEST_ID_CONTEXT_KEY, requestId);
            ServerCall<ReqT, RespT> wrappedCall = new ForwardingServerCall.SimpleForwardingServerCall<>(call) {
                @Override
                public void sendHeaders(Metadata responseHeaders) {
                    responseHeaders.put(X_REQUEST_ID_METADATA, MDC.get(X_REQUEST_ID_CONTEXT_KEY));
                    super.sendHeaders(responseHeaders);
                }
            };


            return Contexts.interceptCall(Context.current().withValue(Context.key(X_REQUEST_ID_CONTEXT_KEY), requestId),
                    wrappedCall, metadata, next);
        } catch (Exception e) {
            throw e;
        }
    }

    /*===============================
     *
     *              REST
     *
     * ==============================*/
    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        String requestId = request.getHeader(X_REQUEST_ID_CONTEXT_KEY);
        if (requestId == null || requestId.isBlank()) {
            requestId = UUID.randomUUID().toString();
        }
        MDC.put(X_REQUEST_ID_CONTEXT_KEY, requestId);
        return true;
    }

    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler, @Nullable ModelAndView modelAndView) throws Exception {
        response.addHeader(X_REQUEST_ID_CONTEXT_KEY, MDC.get(X_REQUEST_ID_CONTEXT_KEY));
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, @Nullable Exception ex) throws Exception {
        MDC.remove(X_REQUEST_ID_CONTEXT_KEY);
    }

    /*===============================
     *
     *              KAFKA
     *
     * ==============================*/
    @Override
    public ConsumerRecords<K, V> onConsume(ConsumerRecords<K, V> records) {
        records.forEach(record -> {
            if (record.key() == X_REQUEST_ID_CONTEXT_KEY) {
                MDC.put(X_REQUEST_ID_CONTEXT_KEY, record.value().toString());
            }
        });
        return records;
    }

    @Override
    public void onCommit(Map<TopicPartition, OffsetAndMetadata> offsets) {

    }

    @Override
    public void close() {

    }

    @Override
    public void configure(Map<String, ?> configs) {

    }
}