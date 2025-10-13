package com.ferza17.ecommercemicroservicesv2.commerceservice.middleware;

import io.grpc.Metadata;
import io.grpc.ServerCall;
import io.grpc.ServerCallHandler;
import io.grpc.ServerInterceptor;
import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.SpanKind;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import jakarta.servlet.*;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.grpc.server.GlobalServerInterceptor;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE)
@Component
public class TracingMiddleware extends OncePerRequestFilter implements ServerInterceptor {
    private static final Logger logger = LoggerFactory.getLogger(TracingMiddleware.class);
    private final Tracer tracer = GlobalOpenTelemetry
            .getTracer("com.ferza17.ecommercemicroservicesv2.commerceservice");

    // GRPC
    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata headers, ServerCallHandler<ReqT, RespT> next) {
        return null;
    }

    // REST
    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        String path = request.getRequestURI();
        String method = request.getMethod();

        // Create span
        Span span = tracer.spanBuilder(method + " " + path)
                .setSpanKind(SpanKind.SERVER)
                .startSpan();

        // Add useful attributes
        span.setAttribute("http.method", method);
        span.setAttribute("http.url", request.getRequestURL().toString());

        try (Scope scope = span.makeCurrent()) {
            filterChain.doFilter(request, response);
            span.setAttribute("http.status_code", response.getStatus());
        } catch (Exception e) {
            span.recordException(e);
            span.setStatus(io.opentelemetry.api.trace.StatusCode.ERROR);
            throw e;
        } finally {
            span.end();
        }
    }
}
