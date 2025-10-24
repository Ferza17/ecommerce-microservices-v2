package com.ferza17.ecommercemicroservicesv2.commerceservice.util;

import io.grpc.Metadata;
import io.opentelemetry.context.propagation.TextMapGetter;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import io.opentelemetry.api.OpenTelemetry;


@Component
public class OpenTelemetryPropagator {
    @Autowired
    private OpenTelemetry openTelemetry;
    private final TextMapGetter<HttpServletRequest> httpServletRequestTextMapGetter = new TextMapGetter<>() {
        @Override
        public Iterable<String> keys(HttpServletRequest carrier) {
            return carrier.getHeaderNames() != null ? java.util.Collections.list(carrier.getHeaderNames()) : java.util.Collections.emptyList();
        }

        @Override
        public String get(HttpServletRequest carrier, String key) {
            return carrier.getHeader(key);
        }
    };
    private final TextMapGetter<io.grpc.Metadata> grpcMetadataTextMapGetter = new TextMapGetter<>() {
        @Override
        public Iterable<String> keys(Metadata carrier) {
            return carrier.keys();
        }

        @Override
        public String get(Metadata carrier, String key) {
            Metadata.Key<String> metadataKey = Metadata.Key.of(key, Metadata.ASCII_STRING_MARSHALLER);
            return carrier.get(metadataKey);
        }
    };

    public io.opentelemetry.context.Context extractContextFromHttpHeader(HttpServletRequest request) {
        var propagator = this.openTelemetry.getPropagators().getTextMapPropagator();
        return propagator.extract(io.opentelemetry.context.Context.current(), request, this.httpServletRequestTextMapGetter);
    }

    public io.opentelemetry.context.Context extractContextFromGrpcMetadata(io.grpc.Metadata metadata) {
        var propagator = this.openTelemetry.getPropagators().getTextMapPropagator();
        return propagator.extract(io.opentelemetry.context.Context.current(), metadata, this.grpcMetadataTextMapGetter);
    }

}
