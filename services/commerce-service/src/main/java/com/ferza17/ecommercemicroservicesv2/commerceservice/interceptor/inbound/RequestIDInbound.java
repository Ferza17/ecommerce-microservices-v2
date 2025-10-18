package com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound;

import io.grpc.*;
import jakarta.servlet.*;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.slf4j.MDC;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.grpc.server.GlobalServerInterceptor;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.UUID;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.X_REQUEST_ID_CONTEXT_KEY;
import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.X_REQUEST_ID_METADATA;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE + 1)
@Component
public class RequestIDInbound extends OncePerRequestFilter implements ServerInterceptor {
    // GRPC
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

    // REST
    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        try {
            String requestId = request.getHeader(X_REQUEST_ID_CONTEXT_KEY);
            if (requestId == null || requestId.isBlank()) {
                requestId = UUID.randomUUID().toString();
            }
            MDC.put(X_REQUEST_ID_CONTEXT_KEY, requestId);
            response.setHeader(X_REQUEST_ID_CONTEXT_KEY, requestId);
            filterChain.doFilter(request, response);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}