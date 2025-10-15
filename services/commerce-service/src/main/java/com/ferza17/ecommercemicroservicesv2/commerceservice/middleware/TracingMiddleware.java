package com.ferza17.ecommercemicroservicesv2.commerceservice.middleware;

import io.grpc.*;
import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.SpanKind;
import io.opentelemetry.api.trace.StatusCode;
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

import static com.google.common.net.HttpHeaders.X_REQUEST_ID;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE + 1)
@Component
public class TracingMiddleware extends OncePerRequestFilter implements ServerInterceptor {
    public static final String TRACEPARENT_HEADER = "traceparent";
    public static final Metadata.Key<String> TRACEPARENT_METADATA =
            Metadata.Key.of(TRACEPARENT_HEADER, Metadata.ASCII_STRING_MARSHALLER);
    private final Tracer tracer = GlobalOpenTelemetry.getTracer(TracingMiddleware.class.getSimpleName());

    // GRPC
    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(
            ServerCall<ReqT, RespT> call,
            Metadata metadata,
            ServerCallHandler<ReqT, RespT> next
    ) {
        this.logger.info("TracingMiddleware");


        // Extract useful info
        String fullMethodName = call.getMethodDescriptor().getFullMethodName();
        String serviceName = fullMethodName.substring(0, fullMethodName.lastIndexOf('/'));
        String methodName = fullMethodName.substring(fullMethodName.lastIndexOf('/') + 1);

        Span span = tracer.spanBuilder(fullMethodName)
                .setSpanKind(SpanKind.SERVER)
                .setAttribute("rpc.system", "grpc")
                .setAttribute("rpc.service", serviceName)
                .setAttribute("rpc.method", methodName)
                .startSpan();


        try (Scope scope = span.makeCurrent()) {
            // Add response header for request-id propagation
            ServerCall<ReqT, RespT> wrappedCall =
                    new ForwardingServerCall.SimpleForwardingServerCall<>(call) {
                        @Override
                        public void sendHeaders(Metadata responseHeaders) {
                            responseHeaders.put(TRACEPARENT_METADATA, span.getSpanContext().getTraceId());
                            super.sendHeaders(responseHeaders);
                        }

                        @Override
                        public void close(Status status, Metadata trailers) {
                            if (!status.isOk()) {
                                span.recordException(status.asRuntimeException());
                                span.setStatus(StatusCode.ERROR, status.getDescription());
                            }
                            super.close(status, trailers);
                            span.end();
                        }
                    };

            return next.startCall(wrappedCall, metadata);
        } catch (Exception e) {
            span.recordException(e);
            span.setStatus(StatusCode.ERROR);
            span.end();
            throw e;
        }
    }

    // REST
    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        this.logger.info("TracingMiddleware");


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
            span.setStatus(StatusCode.ERROR);
            throw e;
        } finally {
            span.end();
        }
    }
}
