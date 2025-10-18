package com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound;

import io.grpc.*;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.apache.kafka.clients.consumer.ConsumerInterceptor;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.OffsetAndMetadata;
import org.apache.kafka.common.TopicPartition;
import org.jboss.logging.MDC;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.grpc.server.GlobalServerInterceptor;
import org.springframework.lang.Nullable;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import javax.security.sasl.AuthenticationException;
import java.util.Map;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.*;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE + 2)
@Component
public class AuthorizationInbound<K, V> implements ServerInterceptor, HandlerInterceptor, ConsumerInterceptor<K, V> {
    /*===============================
     *
     *              GRPC
     *
     * ==============================*/
    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata metadata, ServerCallHandler<ReqT, RespT> next) {
        try {
            String tokenFromHeader = metadata.get(AUTHORIZATION_METADATA);

            if (tokenFromHeader == null || tokenFromHeader.isBlank()) {
                throw new AuthenticationException("Authorization header missing");
            }

            String sanitizedToken = tokenFromHeader.replaceAll("(?i)^Bearer\\s+", "");
            if (sanitizedToken.isBlank()) {
                throw new AuthenticationException("Invalid Token");
            }
            MDC.put(AUTHORIZATION_CONTEXT_KEY, tokenFromHeader);
            ServerCall<ReqT, RespT> wrappedCall = new ForwardingServerCall.SimpleForwardingServerCall<>(call) {
                @Override
                public void sendHeaders(Metadata responseHeaders) {
                    responseHeaders.put(AUTHORIZATION_METADATA, MDC.get(AUTHORIZATION_CONTEXT_KEY).toString());
                    super.sendHeaders(responseHeaders);
                }
            };

            return Contexts.interceptCall(Context.current().withValue(Context.key(AUTHORIZATION_CONTEXT_KEY), tokenFromHeader),
                    wrappedCall, metadata, next);

        } catch (AuthenticationException e) {
            call.close(Status.UNAUTHENTICATED.withDescription(e.getMessage()), new Metadata());
            return new ServerCall.Listener<>() {
            }; // return empty listener to stop call
        } catch (Exception e) {
            call.close(Status.INTERNAL.withDescription(e.getMessage()), new Metadata());
            return new ServerCall.Listener<>() {
            };
        }
    }

    /*===============================
     *
     *              REST
     *
     * ==============================*/
    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        try {
            String tokenFromHeader = request.getHeader(AUTHORIZATION_CONTEXT_KEY);
            if (tokenFromHeader == null || tokenFromHeader.isBlank()) {
                throw new AuthenticationException("Authorization header missing");
            }
            String sanitizedToken = tokenFromHeader.replaceAll("(?i)^Bearer\\s+", "");
            if (sanitizedToken.isBlank()) {
                throw new AuthenticationException("Invalid Token");
            }
            MDC.put(AUTHORIZATION_CONTEXT_KEY, tokenFromHeader);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
        return true;
    }

    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler, @Nullable ModelAndView modelAndView) throws Exception {
        response.addHeader(AUTHORIZATION_CONTEXT_KEY, MDC.get(AUTHORIZATION_CONTEXT_KEY).toString());
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, @Nullable Exception ex) throws Exception {
        MDC.remove(AUTHORIZATION_CONTEXT_KEY);
    }

    /*===============================
     *
     *              KAFKA
     *
     * ==============================*/
    @Override
    public ConsumerRecords<K, V> onConsume(ConsumerRecords<K, V> records) {
        records.forEach(record -> {
            if (record.key() == AUTHORIZATION_CONTEXT_KEY) {
                org.slf4j.MDC.put(AUTHORIZATION_CONTEXT_KEY, record.value().toString());
            }
        });
        return records;
    }

    @Override
    public void onCommit(Map<TopicPartition, OffsetAndMetadata> offsets) {

    }

    @Override
    public void close() {
        MDC.remove(AUTHORIZATION_CONTEXT_KEY);
    }

    @Override
    public void configure(Map<String, ?> configs) {

    }
}
