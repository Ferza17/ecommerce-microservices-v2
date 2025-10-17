package com.ferza17.ecommercemicroservicesv2.commerceservice.middleware.client;

import io.grpc.*;
import org.slf4j.MDC;
import org.springframework.stereotype.Component;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.TRACEPARENT_METADATA;
import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.TRACEPARENT_CONTEXT_KEY;

@Component
public class OpenTelemetryClientMiddleware implements ClientInterceptor {
    @Override
    public <ReqT, RespT> ClientCall<ReqT, RespT> interceptCall(MethodDescriptor<ReqT, RespT> method, CallOptions callOptions, Channel next) {
        try {
            return new ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(
                    next.newCall(method, callOptions)) {

                @Override
                public void start(Listener<RespT> responseListener, Metadata headers) {
                    String traceId = MDC.get(TRACEPARENT_CONTEXT_KEY);
                    headers.put(TRACEPARENT_METADATA, traceId);
                    super.start(responseListener, headers);
                }
            };
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
