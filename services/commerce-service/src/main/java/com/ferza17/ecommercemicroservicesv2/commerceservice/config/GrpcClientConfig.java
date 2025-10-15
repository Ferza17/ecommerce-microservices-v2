package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import com.ferza17.ecommercemicroservicesv2.proto.v1.product.ProductServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.user.UserServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.user.AuthServiceGrpc;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.grpc.client.GrpcChannelFactory;

@Configuration
public class GrpcClientConfig {

    @Bean
    ProductServiceGrpc.ProductServiceBlockingStub productServiceBlockingStub(GrpcChannelFactory grpcChannelFactory) {
        return ProductServiceGrpc.newBlockingStub(grpcChannelFactory.createChannel("product"));
    }

    @Bean
    UserServiceGrpc.UserServiceBlockingStub userServiceBlockingStub(GrpcChannelFactory grpcChannelFactory) {
        return UserServiceGrpc.newBlockingStub(grpcChannelFactory.createChannel("user"));
    }

    @Bean
    AuthServiceGrpc.AuthServiceBlockingStub authServiceBlockingStub(GrpcChannelFactory grpcChannelFactory) {
        return AuthServiceGrpc.newBlockingStub(grpcChannelFactory.createChannel("user"));
    }

}
