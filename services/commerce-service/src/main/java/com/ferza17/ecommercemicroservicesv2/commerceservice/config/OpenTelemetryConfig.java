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
    public io.opentelemetry.api.OpenTelemetry setupOpenTelemetry() {
        JaegerGrpcSpanExporter jaegerExporter = JaegerGrpcSpanExporter.builder().setEndpoint("http://localhost:4317").build();
        Resource serviceNameResource = Resource.getDefault().merge(Resource.create(io.opentelemetry.api.common.Attributes.of(SERVICE_NAME, this.serviceName)));
        SdkTracerProvider tracerProvider = SdkTracerProvider.builder().addSpanProcessor(BatchSpanProcessor.builder(jaegerExporter).build()).setResource(serviceNameResource).build();
        return OpenTelemetrySdk.builder().setTracerProvider(tracerProvider).build();
    }
}
