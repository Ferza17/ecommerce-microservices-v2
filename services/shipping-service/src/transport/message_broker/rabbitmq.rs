use crate::config::config::AppConfig;
use crate::infrastructure::database::async_postgres::get_connection;
use crate::infrastructure::message_broker::rabbitmq::RabbitMQInfrastructure;
use crate::infrastructure::services::payment::PaymentServiceGrpcClient;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::module::shipping::consumer_rabbitmq::ShippingRabbitMQConsumer;
use crate::module::shipping::repository_postgres::ShippingPostgresRepositoryImpl;
use crate::module::shipping::usecase::ShippingUseCaseImpl;
use crate::module::shipping_provider::repository_postgres::ShippingProviderPostgresRepositoryImpl;
use crate::package::worker_pool::worker_pool::{WorkerPool, WorkerPoolError};
use futures::StreamExt;
use std::sync::Arc;
use tokio::task::JoinHandle;
use tracing::{error, info, warn};

pub struct RabbitMQTransport {
    config: AppConfig,
    pool: WorkerPool,
}

impl RabbitMQTransport {
    pub fn new(config: AppConfig, pool: WorkerPool) -> Self {
        RabbitMQTransport { config, pool }
    }

    pub async fn serve(&self) -> std::result::Result<(), Box<dyn std::error::Error>> {
        // infrastructure Layer
        let postgres_pool = get_connection(&self.config.clone()).await;
        let user_service = UserServiceGrpcClient::new(self.config.clone()).await;
        let payment_service = PaymentServiceGrpcClient::new(self.config.clone()).await;
        let rabbitmq_infrastructure =
            Arc::new(RabbitMQInfrastructure::new(self.config.clone()).await);

        // Repository Layer
        let shipping_provider_postgres_repository =
            ShippingProviderPostgresRepositoryImpl::new(postgres_pool.clone());
        let shipping_postgres_repository =
            ShippingPostgresRepositoryImpl::new(postgres_pool.clone());

        // UseCase Layer
        let shipping_use_case = ShippingUseCaseImpl::new(
            shipping_postgres_repository,
            shipping_provider_postgres_repository,
            user_service,
            payment_service,
        );

        // Wrap in Arc for sharing across tasks
        let shipping_rabbitmq_consumer = Arc::new(ShippingRabbitMQConsumer::new(shipping_use_case));
        let mut handles: Vec<(String, String, Result<JoinHandle<()>, WorkerPoolError>)> =
            Vec::new();
        {
            let queue_shipping_created = self
                .config
                .service_shipping_rabbitmq
                .queue_shipping_created
                .clone();
            let exchange_shipping = self
                .config
                .service_shipping_rabbitmq
                .exchange_shipping
                .clone();
            let rabbitmq_infra_clone = Arc::clone(&rabbitmq_infrastructure);
            let consumer_clone = Arc::clone(&shipping_rabbitmq_consumer);
            handles.push((
                self.config.clone().service_shipping_rabbitmq.exchange_shipping,
                self.config.clone().service_shipping_rabbitmq.queue_shipping_created,
                Ok(self
                    .pool
                    .spawn(move || async move {
                        let mut messages = rabbitmq_infra_clone
                            .binding(&queue_shipping_created, &exchange_shipping)
                            .await
                            .setup_consumer(&queue_shipping_created)
                            .await;

                        while let Some(delivery) = messages.next().await {
                            match delivery {
                                Ok(delivery) => {
                                    match consumer_clone.consumer_shipping_created(delivery).await {
                                        Ok(_) => {}
                                        Err(e) => {
                                            error!("[RABBITMQ] Error consumer_shipping_created message: {}", e);
                                        }
                                    }
                                }
                                Err(e) => {
                                    error!("[RABBITMQ] Error receiving message: {}", e);
                                }
                            }
                        }
                    })
                    .await?),
            ));
        }
        {
            let queue_shipping_updated = self
                .config
                .service_shipping_rabbitmq
                .queue_shipping_updated
                .clone();
            let exchange_shipping = self
                .config
                .service_shipping_rabbitmq
                .exchange_shipping
                .clone();
            let rabbitmq_infra_clone = Arc::clone(&rabbitmq_infrastructure);
            let consumer_clone = Arc::clone(&shipping_rabbitmq_consumer);

            let handle = self
                .pool
                .spawn(move || async move {
                    let mut messages = rabbitmq_infra_clone
                        .binding(&queue_shipping_updated, &exchange_shipping)
                        .await
                        .setup_consumer(&queue_shipping_updated)
                        .await;

                    while let Some(delivery) = messages.next().await {
                        match delivery {
                            Ok(delivery) => {
                                match consumer_clone.consumer_shipping_updated(delivery).await {
                                    Ok(_) => {}
                                    Err(e) => {
                                        error!("[RABBITMQ] Error consumer_shipping_updated message: {}", e);
                                    }
                                }
                            }
                            Err(e) => {
                                error!("[RABBITMQ] Error receiving message: {}", e);
                            }
                        }
                    }
                })
                .await?;
            handles.push((
                self.config
                    .clone()
                    .service_shipping_rabbitmq
                    .exchange_shipping,
                self.config
                    .clone()
                    .service_shipping_rabbitmq
                    .queue_shipping_created,
                Ok(handle),
            ));
        }
        for (exchange, queue, handle) in handles {
            if let Err(e) = handle.unwrap().await {
                error!("[RABBITMQ] Error: {}", e);
            }
        }

        Ok(())
    }
}
