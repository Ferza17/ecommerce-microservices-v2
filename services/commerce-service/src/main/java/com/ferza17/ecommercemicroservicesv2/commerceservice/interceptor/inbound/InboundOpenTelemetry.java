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

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.*;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE)
@Component
public class InboundOpenTelemetry<K, V> implements ServerInterceptor, HandlerInterceptor, ConsumerInterceptor<K, V> {
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
            String traceparent = metadata.get(TRACEPARENT_METADATA);
            if (traceparent == null || traceparent.isBlank()) {
                traceparent = "";
            }
            MDC.put(TRACEPARENT_CONTEXT_KEY, traceparent);
            ServerCall<ReqT, RespT> wrappedCall = new ForwardingServerCall.SimpleForwardingServerCall<>(call) {
                @Override
                public void sendHeaders(Metadata responseHeaders) {
                    responseHeaders.put(TRACEPARENT_METADATA, MDC.get(TRACEPARENT_CONTEXT_KEY));
                    super.sendHeaders(responseHeaders);
                }
            };

            return Contexts.interceptCall(Context.current().withValue(Context.key(TRACEPARENT_CONTEXT_KEY), traceparent),
                    wrappedCall, metadata, next);
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
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        String traceparent = request.getHeader(TRACEPARENT_CONTEXT_KEY);
        if (traceparent == null || traceparent.isBlank()) {
            traceparent = "";
        }
        MDC.put(TRACEPARENT_CONTEXT_KEY, traceparent);
        return true;
    }

    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler, @Nullable ModelAndView modelAndView) throws Exception {
        response.addHeader(TRACEPARENT_CONTEXT_KEY, MDC.get(TRACEPARENT_CONTEXT_KEY));
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, @Nullable Exception ex) throws Exception {
        MDC.remove(TRACEPARENT_CONTEXT_KEY);
    }

    /*===============================
     *
     *              KAFKA
     *
     * ==============================*/
    @Override
    public ConsumerRecords<K, V> onConsume(ConsumerRecords<K, V> records) {
        records.forEach(record -> {
            if (record.key() == TRACEPARENT_CONTEXT_KEY) {
                MDC.put(TRACEPARENT_CONTEXT_KEY, record.value().toString());
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