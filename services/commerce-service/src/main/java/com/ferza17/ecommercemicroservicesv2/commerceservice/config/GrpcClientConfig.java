package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound.AuthorizationOutbound;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound.OpenTelemetryOutbound;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.outbound.RequestIDClientOutbound;
import com.ferza17.ecommercemicroservicesv2.proto.v1.product.ProductServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.user.UserServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.user.AuthServiceGrpc;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.grpc.client.GrpcChannelFactory;

@Configuration
public class GrpcClientConfig {
    @Autowired
    private GrpcChannelFactory grpcChannelFactory;
    // ClientMiddleware
    @Autowired
    RequestIDClientOutbound requestIDClientMiddleware;
    @Autowired
    OpenTelemetryOutbound openTelemetryClientMiddleware;
    @Autowired
    AuthorizationOutbound authorizationClientMiddleware;



    @Bean
    ProductServiceGrpc.ProductServiceBlockingStub productServiceBlockingStub() {
        return ProductServiceGrpc
                .newBlockingStub(grpcChannelFactory.createChannel("product"))
                .withInterceptors(requestIDClientMiddleware, openTelemetryClientMiddleware, authorizationClientMiddleware);
    }

    @Bean
    UserServiceGrpc.UserServiceBlockingStub userServiceBlockingStub() {
        return UserServiceGrpc
                .newBlockingStub(grpcChannelFactory.createChannel("user"))
                .withInterceptors(requestIDClientMiddleware, openTelemetryClientMiddleware, authorizationClientMiddleware);
    }

    @Bean
    AuthServiceGrpc.AuthServiceBlockingStub authServiceBlockingStub() {
        return AuthServiceGrpc
                .newBlockingStub(grpcChannelFactory.createChannel("user"))
                .withInterceptors(requestIDClientMiddleware, openTelemetryClientMiddleware, authorizationClientMiddleware);
    }

}
