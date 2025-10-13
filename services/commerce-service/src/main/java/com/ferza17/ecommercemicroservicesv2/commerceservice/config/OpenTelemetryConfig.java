package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.exporter.otlp.trace.OtlpGrpcSpanExporter;
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
    public io.opentelemetry.api.OpenTelemetry openTelemetry() {
        Resource serviceNameResource = Resource.getDefault()
                .merge(Resource.create(io.opentelemetry.api.common.Attributes.of(SERVICE_NAME, this.serviceName)));

        OtlpGrpcSpanExporter exporter = OtlpGrpcSpanExporter.builder()
                .setEndpoint("http://localhost:4317") // adjust for your OTLP collector
                .build();

        SdkTracerProvider tracerProvider = SdkTracerProvider.builder()
                .addSpanProcessor(BatchSpanProcessor.builder(exporter).build())
                .setResource(serviceNameResource)
                .build();

        OpenTelemetrySdk openTelemetry = OpenTelemetrySdk.builder()
                .setTracerProvider(tracerProvider)
                .buildAndRegisterGlobal();

        GlobalOpenTelemetry.set(openTelemetry);
        return openTelemetry;
    }
}
