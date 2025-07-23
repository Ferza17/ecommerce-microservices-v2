use crate::config::config::AppConfig;
use crate::transport::{
    grpc::grpc::GrpcTransport,
    http::{http::HttpTransport, metric::serve_metric_http_collector},
    message_broker::rabbitmq::RabbitMQTransport,
};
use std::sync::Arc;

use crate::package::worker_pool::typed_worker_pool::TypedWorkerPool;
use clap::Args;
use tracing::{error, info};

#[derive(Args, Debug)]
pub struct RunArgs {
    #[arg(short, long, help = "run direction: 'local' or 'production'")]
    pub direction: String,
}
pub async fn handle_run_command(args: RunArgs) {
    // Create specialized worker pools
    let pools = Arc::new(TypedWorkerPool::new(5, 5, 10, 1));

    let cfg = AppConfig::new(&*args.direction)
        .await
        .map_err(|e| {
            error!("Failed to load configuration: {}", e);
            std::process::exit(1);
        })
        .unwrap();

    let cfg = Arc::new(cfg);
    let mut handles = Vec::new();

    // HTTP with dedicated pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = Arc::clone(&cfg);
        let handle: tokio::task::JoinHandle<Result<(), anyhow::Error>> = pool
            .http_pool
            .spawn(move || async move {
                let http_transport = HttpTransport::new((*cfg_clone).clone());
                Ok(http_transport.serve().await.expect("HTTP service failed"))
            })
            .await;
        handles.push(handle);
    }

    // GRPC with dedicated pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = Arc::clone(&cfg);
        let handle: tokio::task::JoinHandle<Result<(), anyhow::Error>> = pool
            .grpc_pool
            .spawn(move || async move {
                let grpc_transport = GrpcTransport::new((*cfg_clone).clone());
                Ok(grpc_transport.serve().await.expect("GRPC service failed"))
            })
            .await;
        handles.push(handle);
    }

    // RabbitMQ with messaging pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = Arc::clone(&cfg);
        let handle: tokio::task::JoinHandle<Result<(), anyhow::Error>> = pool
            .messaging_pool
            .spawn(move || async move {
                let rabbitmq_transport = RabbitMQTransport::new((*cfg_clone).clone());
                Ok(rabbitmq_transport.serve().await)
            })
            .await;
        handles.push(handle);
    }

    // Metrics with dedicated pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = Arc::clone(&cfg);
        let handle: tokio::task::JoinHandle<Result<(), anyhow::Error>> = pool
            .metrics_pool
            .spawn(
                move || async move { Ok(serve_metric_http_collector((*cfg_clone).clone()).await) },
            )
            .await;
        handles.push(handle);
    }

    info!("All services started with typed worker pools");

    // Wait for all services
    for handle in handles {
        if let Err(e) = handle.await {
            error!("Service task failed: {}", e);
        }
    }
}
