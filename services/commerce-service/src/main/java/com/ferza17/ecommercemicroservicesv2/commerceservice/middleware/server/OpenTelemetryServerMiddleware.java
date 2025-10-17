package com.ferza17.ecommercemicroservicesv2.commerceservice.middleware.server;

import io.grpc.*;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.slf4j.MDC;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.grpc.server.GlobalServerInterceptor;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.*;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE)
@Component
public class OpenTelemetryServerMiddleware extends OncePerRequestFilter implements ServerInterceptor {
    // GRPC
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

    // REST
    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        try {
            String traceparent = request.getHeader(TRACEPARENT_CONTEXT_KEY);
            if (traceparent == null || traceparent.isBlank()) {
                traceparent = "";
            }
            MDC.put(TRACEPARENT_CONTEXT_KEY, traceparent);
            response.setHeader(TRACEPARENT_CONTEXT_KEY, traceparent);
            filterChain.doFilter(request, response);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}