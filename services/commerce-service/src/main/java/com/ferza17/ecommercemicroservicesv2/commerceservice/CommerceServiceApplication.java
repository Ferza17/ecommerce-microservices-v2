package com.ferza17.ecommercemicroservicesv2.commerceservice;

import com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.WorkerPool;
import com.ferza17.ecommercemicroservicesv2.commerceservice.transport.Grpc;
import com.ferza17.ecommercemicroservicesv2.commerceservice.transport.Http;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class CommerceServiceApplication {

    public static void main(String[] args) {
        try {
            System.out.println("🚀 Starting Commerce Service...");

            WorkerPool.submit(new Grpc());
            WorkerPool.submit(new Http());

            Runtime.getRuntime().addShutdownHook(new Thread(() -> {
                System.out.println("🛑 Shutting down worker pool...");
                WorkerPool.shutdown();
            }));

        } catch (Exception e) {
            System.err.println("❌ Failed to start gRPC server: " + e.getMessage());
        }

    }

}
