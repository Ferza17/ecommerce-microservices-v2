package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class GrpcGatewayConfig {
    @Value("${grpc.server.host:localhost}")
    private String host;

    @Value("${grpc.server.port:50051}")
    private int port;

    @Bean
    public ManagedChannel grpcChannel() {
        return ManagedChannelBuilder.forAddress(host, port).usePlaintext().build();
    }
}
