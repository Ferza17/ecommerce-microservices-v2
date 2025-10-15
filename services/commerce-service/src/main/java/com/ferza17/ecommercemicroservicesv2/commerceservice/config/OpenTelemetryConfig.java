package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import io.opentelemetry.exporter.jaeger.JaegerGrpcSpanExporter;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.trace.SdkTracerProvider;
import io.opentelemetry.sdk.trace.export.BatchSpanProcessor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import io.opentelemetry.sdk.resources.Resource;

import static io.opentelemetry.semconv.ServiceAttributes.SERVICE_NAME;


@Configuration
public class OpenTelemetryConfig {
    @Value("${spring.application.name:commerce-service}")
    private String serviceName;

    @Bean
    public io.opentelemetry.sdk.OpenTelemetrySdk setupOpenTelemetry() {
        return OpenTelemetrySdk
                .builder()
                .setTracerProvider(SdkTracerProvider
                        .builder()
                        .addSpanProcessor(BatchSpanProcessor
                                .builder(JaegerGrpcSpanExporter.builder().setEndpoint("http://localhost:4317").build())
                                .build()
                        )
                        .setResource(Resource.getDefault().merge(Resource.create(io.opentelemetry.api.common.Attributes.of(SERVICE_NAME, this.serviceName))))
                        .build()
                )
                .build();
    }
}
