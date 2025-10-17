package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.exporter.jaeger.thrift.JaegerThriftSpanExporter;
import io.opentelemetry.sdk.trace.SdkTracerProvider;
import io.opentelemetry.sdk.trace.export.BatchSpanProcessor;
import io.opentelemetry.semconv.ResourceAttributes;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import io.opentelemetry.sdk.resources.Resource;
import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.sdk.OpenTelemetrySdk;



@Configuration
public class OpenTelemetryConfig {
    @Value("${spring.application.name}")
    private String appName;
    @Value("${spring.application.version}")
    private String appVersion;

    @Bean
    public JaegerThriftSpanExporter jaegerThriftSpanExporter() {
        return JaegerThriftSpanExporter.builder()
                .setEndpoint("http://localhost:14268/api/traces")
                .build();
    }

    @Bean
    public SdkTracerProvider sdkTracerProvider(JaegerThriftSpanExporter jaegerExporter) {
        Resource resource = Resource.getDefault()
                .merge(Resource.builder()
                        .put(ResourceAttributes.SERVICE_NAME, this.appName)
                        .put(ResourceAttributes.SERVICE_VERSION, this.appVersion)
                        .build());

        SdkTracerProvider tracerProvider = SdkTracerProvider.builder()
                .addSpanProcessor(BatchSpanProcessor.builder(jaegerExporter).build())
                .setResource(resource)
                .build();

        return tracerProvider;
    }

    @Bean
    public OpenTelemetry openTelemetry(SdkTracerProvider sdkTracerProvider) {
        return OpenTelemetrySdk.builder()
                .setTracerProvider(sdkTracerProvider)
                .build();
    }

    @Bean
    public Tracer tracer(OpenTelemetry openTelemetry) {
        return openTelemetry.getTracer("com.ferza17.ecommercemicroservicev2.commerceservice");
    }
}
