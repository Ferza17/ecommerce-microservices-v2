package com.ferza17.ecommercemicroservicesv2.commerceservice;

import org.springframework.boot.Banner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.scheduling.annotation.EnableAsync;

import java.util.concurrent.Executor;
import java.util.concurrent.Executors;

@SpringBootApplication
@EnableAsync
public class CommerceServiceApplication {

    public static void main(String[] args) {
        SpringApplication app = new SpringApplication(CommerceServiceApplication.class);
//        app.setBanner(Banner.Mode.OFF);
        app.run(args);
    }

    @Bean(name = "grpcExecutor")
    public Executor grpcExecutor() {
        return Executors.newVirtualThreadPerTaskExecutor();
    }

}
