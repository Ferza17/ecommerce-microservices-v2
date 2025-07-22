use crate::config::config::AppConfig;
use crate::transport::{
    grpc::grpc::GrpcTransport,
    http::{http::HttpTransport, metric::serve_metric_http_collector},
    message_broker::rabbitmq::RabbitMQTransport,
};
use std::sync::Arc;

use crate::package::worker_pool::typed_worker_pool::TypedWorkerPool;
use clap::Args;

#[derive(Args, Debug)]
pub struct RunArgs {
    #[arg(short, long, help = "run direction: 'local' or 'production'")]
    pub direction: String,
}
pub async fn handle_run_command(args: RunArgs) {
    // Create specialized worker pools
    let pools = Arc::new(TypedWorkerPool::new(5, 3, 4, 2));

    let cfg = AppConfig::new(&*args.direction)
        .await
        .map_err(|e| {
            eprintln!("Failed to load configuration: {}", e);
            std::process::exit(1);
        })
        .unwrap();

    let cfg = Arc::new(cfg);
    let mut handles = Vec::new();

    // HTTP with dedicated pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = Arc::clone(&cfg);
        let handle = pool
            .http_pool
            .spawn(move || async move {
                let http_transport = HttpTransport::new((*cfg_clone).clone());
                http_transport.serve().await
            })
            .await;
        handles.push(handle);
    }

    // GRPC with dedicated pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = Arc::clone(&cfg);
        let handle = pool
            .grpc_pool
            .spawn(move || async move {
                let grpc_transport = GrpcTransport::new((*cfg_clone).clone());
                grpc_transport.serve().await.expect("GRPC service failed")
            })
            .await;
        handles.push(handle);
    }

    // RabbitMQ with messaging pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = Arc::clone(&cfg);
        let handle = pool
            .messaging_pool
            .spawn(move || async move {
                let rabbitmq_transport = RabbitMQTransport::new((*cfg_clone).clone());
                rabbitmq_transport.serve().await
            })
            .await;
        handles.push(handle);
    }

    // Metrics with dedicated pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = Arc::clone(&cfg);
        let handle = pool
            .metrics_pool
            .spawn(move || async move { serve_metric_http_collector((*cfg_clone).clone()).await })
            .await;
        handles.push(handle);
    }

    println!("All services started with typed worker pools");

    // Wait for all services
    for handle in handles {
        if let Err(e) = handle.await {
            eprintln!("Service task failed: {}", e);
        }
    }
}
