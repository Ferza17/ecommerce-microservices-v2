package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.protobuf.services.ProtoReflectionServiceV1;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;

import java.io.IOException;


public class GrpcConfig {
    @Value("${grpc.server.port:50051}")
    private int grpcServerPort;

    @Bean
    public Server grpcServer(
            com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart.PresenterGrpc cartPresenterGrpc,
            com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist.PresenterGrpc wishlistPresenterGrpc
    ) throws IOException {
        Server server = ServerBuilder.
                forPort(grpcServerPort).
                addService(cartPresenterGrpc).
                addService(wishlistPresenterGrpc).
                addService(ProtoReflectionServiceV1.newInstance()).
                build();
        server.start();

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("Shutting down gRPC server");
            server.shutdown();
            System.out.println("gRPC server shut down");
        }));

        System.out.println("gRPC Server started on port: " + grpcServerPort);

        return server;
    }
}
