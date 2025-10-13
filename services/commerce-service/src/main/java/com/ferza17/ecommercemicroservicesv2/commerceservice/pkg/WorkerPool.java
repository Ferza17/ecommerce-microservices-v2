package com.ferza17.ecommercemicroservicesv2.commerceservice.pkg;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class WorkerPool {
    private static final int WORKER_COUNT = 2; // one for grpc, one for http
    private static final ExecutorService executor = Executors.newFixedThreadPool(WORKER_COUNT);

    public static void submit(Runnable task) {
        executor.submit(task);
    }

    public static void shutdown() {
        executor.shutdown();
    }
}

