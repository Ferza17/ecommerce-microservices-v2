package com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context;

import io.grpc.Metadata;

import static com.google.common.net.HttpHeaders.AUTHORIZATION;
import static com.google.common.net.HttpHeaders.X_REQUEST_ID;


// TO GET THIS VALUE
// USING LIBRARY import org.slf4j.MDC;
public class BaseContext {
    public static final String TRACEPARENT_CONTEXT_KEY = "traceparent";
    public static final String SPAN_ID_CONTEXT_KEY = "span_id";
    public static final String X_REQUEST_ID_CONTEXT_KEY = X_REQUEST_ID;
    public static final String AUTHORIZATION_CONTEXT_KEY = AUTHORIZATION;

    public static final Metadata.Key<String> TRACEPARENT_METADATA = Metadata.Key.of(TRACEPARENT_CONTEXT_KEY, Metadata.ASCII_STRING_MARSHALLER);
    public static final Metadata.Key<String> X_REQUEST_ID_METADATA = Metadata.Key.of(X_REQUEST_ID_CONTEXT_KEY, Metadata.ASCII_STRING_MARSHALLER);
    public static final Metadata.Key<String> AUTHORIZATION_METADATA = Metadata.Key.of(AUTHORIZATION_CONTEXT_KEY, Metadata.ASCII_STRING_MARSHALLER);
}
