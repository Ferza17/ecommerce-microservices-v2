package com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound;

import io.grpc.*;
import org.slf4j.MDC;
import org.springframework.stereotype.Component;
import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.AUTHORIZATION_CONTEXT_KEY;
import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.AUTHORIZATION_METADATA;

@Component
public class AuthorizationOutbound implements ClientInterceptor {
    @Override
    public <ReqT, RespT> ClientCall<ReqT, RespT> interceptCall(MethodDescriptor<ReqT, RespT> method, CallOptions callOptions, Channel next) {
        try {
            return new ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(
                    next.newCall(method, callOptions)) {

                @Override
                public void start(Listener<RespT> responseListener, Metadata headers) {
                    headers.put(AUTHORIZATION_METADATA, MDC.get(AUTHORIZATION_CONTEXT_KEY));
                    super.start(responseListener, headers);
                }
            };
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}