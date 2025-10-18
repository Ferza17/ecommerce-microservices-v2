package com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound;

import io.grpc.*;
import org.slf4j.MDC;
import org.springframework.stereotype.Component;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.*;

@Component
public class RequestIDClientOutbound implements ClientInterceptor {
    @Override
    public <ReqT, RespT> ClientCall<ReqT, RespT> interceptCall(MethodDescriptor<ReqT, RespT> method, CallOptions callOptions, Channel next) {
        try {
            return new ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(
                    next.newCall(method, callOptions)) {

                @Override
                public void start(Listener<RespT> responseListener, Metadata headers) {
                    headers.put(X_REQUEST_ID_METADATA, MDC.get(X_REQUEST_ID_CONTEXT_KEY));
                    super.start(responseListener, headers);
                }
            };
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
