use crate::config::config::AppConfig;
use crate::transport::{
    grpc::grpc::GrpcTransport,
    http::{http::HttpTransport, metric::serve_metric_http_collector},
    message_broker::rabbitmq::RabbitMQTransport,
};
use std::sync::Arc;

use crate::infrastructure::telemetry::jaeger::init_tracing;
use crate::package::worker_pool::typed_worker_pool::TypedWorkerPool;
use crate::package::worker_pool::worker_pool::WorkerPoolError;
use clap::Args;
use tokio::task::JoinHandle;

#[derive(Args, Debug)]
pub struct RunArgs {
    #[arg(short, long, help = "run direction: 'local' or 'production'")]
    pub direction: String,
}
pub async fn handle_run_command(args: RunArgs) {
    // Init config
    let cfg = AppConfig::new(&*args.direction)
        .await
        .map_err(|e| {
            eprintln!("Failed to load configuration: {}", e);
            std::process::exit(1);
        })
        .unwrap();

    // ======= WORKER POOLS ===========
    // Create specialized worker pools
    let pools = Arc::new(TypedWorkerPool::new(5, 5, 1000, 1));
    let mut handles: Vec<(
        String,
        Result<JoinHandle<Result<(), anyhow::Error>>, WorkerPoolError>,
    )> = Vec::new();
    // HTTP Transport with dedicated pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = cfg.clone();
        let handle: Result<JoinHandle<Result<(), anyhow::Error>>, WorkerPoolError> = pool
            .spawn_http_task(move || async move {
                let transport = HttpTransport::new(cfg_clone);
                match transport.serve().await {
                    Ok(_) => {
                        tracing::info!("HTTP service completed successfully");
                        Ok(())
                    }
                    Err(e) => {
                        tracing::error!("HTTP service failed: {}", e);
                        Err(anyhow::anyhow!("HTTP service failed: {}", e))
                    }
                }
            })
            .await;

        handles.push((format!("{:?}", pool.http_pool.worker_type()), handle));
    }
    // GRPC Transport with dedicated pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = cfg.clone();
        let handle: Result<JoinHandle<Result<(), anyhow::Error>>, WorkerPoolError> = pool
            .spawn_grpc_task(move || async move {
                let transport = GrpcTransport::new(cfg_clone);
                match transport.serve().await {
                    Ok(_) => {
                        tracing::info!("GRPC service completed successfully");
                        Ok(())
                    }
                    Err(e) => {
                        tracing::error!("GRPC service failed: {}", e);
                        Err(anyhow::anyhow!("GRPC service failed: {}", e))
                    }
                }
            })
            .await;
        handles.push((format!("{:?}", pool.grpc_pool.worker_type()), handle));
    }
    // METRIC HTTP Transport with dedicated pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = cfg.clone();
        let handle: Result<JoinHandle<Result<(), anyhow::Error>>, WorkerPoolError> = pool
            .spawn_metrics_task(move || async move {
                match serve_metric_http_collector(cfg_clone).await {
                    Ok(_) => {
                        tracing::info!("Metric service completed successfully");
                        Ok(())
                    }
                    Err(e) => {
                        tracing::error!("Metric service failed: {}", e);
                        Err(anyhow::anyhow!("Metric service failed: {}", e))
                    }
                }
            })
            .await;
        handles.push((format!("{:?}", pool.metrics_pool.worker_type()), handle));
    }
    // RabbitMQ with messaging pool
    {
        let pool = Arc::clone(&pools);
        let cfg_clone = cfg.clone();
        let pools_clone = Arc::clone(&pools);
        let handle: Result<JoinHandle<Result<(), anyhow::Error>>, WorkerPoolError> = pool
            .spawn_rabbitmq_task(move || async move {
                let transport =
                    RabbitMQTransport::new(cfg_clone, pools_clone.messaging_pool.clone());
                match transport.serve().await {
                    Ok(_) => {
                        tracing::info!("RABBITMQ service completed successfully");
                        Ok(())
                    }
                    Err(e) => {
                        tracing::error!("RABBITMQ service failed: {}", e);
                        Err(anyhow::anyhow!("RABBITMQ service failed: {}", e))
                    }
                }
            })
            .await;
        handles.push((format!("{:?}", pool.messaging_pool.worker_type()), handle));
    }
    // Wait for all services
    for (service_name, handle) in handles {
        handle
            .unwrap()
            .await
            .unwrap()
            .map_err(|e| {
                eprintln!("Service {} failed: {}", service_name, e);
                anyhow::anyhow!("Service {} failed: {}", service_name, e)
            })
            .expect("TODO: panic message");
    }
}
