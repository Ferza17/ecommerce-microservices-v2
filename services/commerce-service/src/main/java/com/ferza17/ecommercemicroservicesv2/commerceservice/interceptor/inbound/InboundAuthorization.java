package com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound;

import com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.exception.BaseErrorCode;
import com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.exception.BaseException;
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
import java.util.Map;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.*;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE + 2)
@Component
public class InboundAuthorization<K, V> implements ServerInterceptor, HandlerInterceptor, ConsumerInterceptor<K, V> {
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
                throw new BaseException(BaseErrorCode.UNAUTHENTICATED, String.format("%s header missing", AUTHORIZATION_CONTEXT_KEY));
            }

            String sanitizedToken = tokenFromHeader.replaceAll("(?i)^Bearer\\s+", "");
            if (sanitizedToken.isBlank()) {
                throw new BaseException(BaseErrorCode.UNAUTHENTICATED, "Invalid Token");
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

        } catch (BaseException e) {
            call.close(Status.fromCode(e.getBaseErrorCode().getCode()).withDescription(e.getMessage()), new Metadata());
            return new ServerCall.Listener<>() {
            }; // return empty listener to stop call
        } catch (Exception e) {
            call.close(Status.fromCode(BaseErrorCode.INTERNAL_ERROR.getCode()).withDescription(BaseErrorCode.INTERNAL_ERROR.getMessage()), new Metadata());
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
                throw new BaseException(BaseErrorCode.UNAUTHENTICATED, String.format("%s header missing", AUTHORIZATION_CONTEXT_KEY));
            }
            String sanitizedToken = tokenFromHeader.replaceAll("(?i)^Bearer\\s+", "");
            if (sanitizedToken.isBlank()) {
                throw new BaseException(BaseErrorCode.UNAUTHENTICATED, "Invalid Token");
            }
            MDC.put(AUTHORIZATION_CONTEXT_KEY, tokenFromHeader);
        } catch (BaseException e) {
            throw e;
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
