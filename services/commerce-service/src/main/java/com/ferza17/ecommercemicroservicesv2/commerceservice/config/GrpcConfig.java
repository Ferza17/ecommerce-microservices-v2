package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart.CartGrpcService;
import com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist.WishlistGrpcService;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.protobuf.services.HealthStatusManager;
import io.grpc.protobuf.services.ProtoReflectionService;
import io.grpc.protobuf.services.ProtoReflectionServiceV1;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;


@Configuration
public class GrpcConfig {
    @Value("${grpc.server.port:50051}")
    private int grpcServerPort;

    @Bean
    public Server grpcServer(
            CartGrpcService cartPresenterGrpc,
            WishlistGrpcService wishlistPresenterGrpc
    ) throws IOException {
        Server server = ServerBuilder.
                forPort(grpcServerPort).
                addService(cartPresenterGrpc).
                addService(wishlistPresenterGrpc).
                addService(new HealthStatusManager().getHealthService()).
                addService(ProtoReflectionService.newInstance()).
                build();

        server.start();
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("Shutting down gRPC server");
            server.shutdown();
            System.out.println("gRPC server shut down");
        }));


        return server;
    }
}
