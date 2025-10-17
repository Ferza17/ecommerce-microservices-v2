package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import io.grpc.protobuf.services.ProtoReflectionService;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.grpc.autoconfigure.server.GrpcServerFactoryCustomizer;

@Configuration
public class GrpcServerConfig {
    @Bean
    public GrpcServerFactoryCustomizer reflectionAndHealthConfigurer() {
        return serverBuilder -> {
            // Enable gRPC reflection
            serverBuilder.addService(ProtoReflectionService.newInstance().bindService());
        };
    }
}
