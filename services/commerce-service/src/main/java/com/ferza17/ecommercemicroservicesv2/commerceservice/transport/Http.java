package com.ferza17.ecommercemicroservicesv2.commerceservice.transport;

import com.ferza17.ecommercemicroservicesv2.commerceservice.CommerceServiceApplication;
import org.springframework.boot.SpringApplication;
import org.springframework.context.ConfigurableApplicationContext;

public class Http implements Runnable {
    private ConfigurableApplicationContext context;

    public void run() {
        System.out.println("🌐 Starting HTTP server on port 40051...");
        context = SpringApplication.run(CommerceServiceApplication.class);
    }

    public void shutdown() {
        if (context != null) {
            context.close();
        }
    }
}
