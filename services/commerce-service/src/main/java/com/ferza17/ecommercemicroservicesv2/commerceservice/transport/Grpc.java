package com.ferza17.ecommercemicroservicesv2.commerceservice.transport;

import io.grpc.Server;
import io.grpc.netty.NettyServerBuilder;
import io.grpc.protobuf.services.ProtoReflectionServiceV1;



public class Grpc implements Runnable {
    private Server server;

    @Override
    public void run() {
        try {
            server = NettyServerBuilder
                    .forPort(50051)
                    .addService(new com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart.Presenter())
                    .addService(new com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist.Presenter())
                    .addService(ProtoReflectionServiceV1.newInstance())
                    .build()
                    .start();

            System.out.println("✅ gRPC server started on port 50051");
            server.awaitTermination();
        } catch (Exception e) {
            System.err.println("❌ Failed to start gRPC server: " + e.getMessage());
            Thread.currentThread().interrupt();
        }
    }

    public void shutdown() throws InterruptedException {
        if (server != null) {
            server.shutdown().awaitTermination();
        }
    }
}
