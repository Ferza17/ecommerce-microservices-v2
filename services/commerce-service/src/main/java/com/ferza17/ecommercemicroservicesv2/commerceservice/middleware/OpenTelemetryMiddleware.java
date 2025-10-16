package com.ferza17.ecommercemicroservicesv2.commerceservice.middleware;

import io.grpc.*;
import jakarta.servlet.*;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.slf4j.MDC;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.grpc.server.GlobalServerInterceptor;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.UUID;

import static com.google.common.net.HttpHeaders.X_REQUEST_ID;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE)
@Component
public class RequestIDMiddleware extends OncePerRequestFilter implements ServerInterceptor {
    private static final Metadata.Key<String> REQUEST_ID_METADATA =
            Metadata.Key.of(X_REQUEST_ID, Metadata.ASCII_STRING_MARSHALLER);

    // GRPC
    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(
            ServerCall<ReqT, RespT> call,
            Metadata metadata,
            ServerCallHandler<ReqT, RespT> next
    ) {
        this.logger.info("RequestIDMiddleware");

        String requestId = metadata.get(REQUEST_ID_METADATA);
        if (requestId == null || requestId.isBlank()) {
            requestId = UUID.randomUUID().toString();
        }

        MDC.put(X_REQUEST_ID, requestId);


        // Add the request ID to response headers for downstream clients
        String finalRequestId = requestId;
        ServerCall<ReqT, RespT> wrappedCall = new ForwardingServerCall.SimpleForwardingServerCall<>(call) {
            @Override
            public void sendHeaders(Metadata responseHeaders) {
                responseHeaders.put(REQUEST_ID_METADATA, finalRequestId);
                super.sendHeaders(responseHeaders);
            }
        };


        return Contexts.interceptCall(Context.current().withValue(Context.key(X_REQUEST_ID), requestId),
                wrappedCall, metadata, next);
    }

    // REST
    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {

        this.logger.info("RequestIDMiddleware");

        String requestId = request.getHeader(X_REQUEST_ID);
        if (requestId == null || requestId.isBlank()) {
            requestId = UUID.randomUUID().toString();
        }

        // Add to MDC (for logging) and response header
        MDC.put(X_REQUEST_ID, requestId);
        response.setHeader(X_REQUEST_ID, requestId);

        try {
            filterChain.doFilter(request, response);
        } finally {
            MDC.remove(X_REQUEST_ID);
        }
    }
}