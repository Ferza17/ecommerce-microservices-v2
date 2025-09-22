use crate::config::config::AppConfig;
use crate::infrastructure::database::async_postgres::get_connection;
use crate::infrastructure::message_broker::kafka::KafkaInfrastructure;
use crate::infrastructure::message_broker::rabbitmq::RabbitMQInfrastructure;
use crate::infrastructure::services::payment::PaymentServiceGrpcClient;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::module::shipping::repository_postgres::ShippingPostgresRepositoryImpl;
use crate::module::shipping::usecase::ShippingUseCaseImpl;
use crate::module::shipping_provider::repository_postgres::ShippingProviderPostgresRepositoryImpl;
use crate::package::worker_pool::worker_pool::{WorkerPool, WorkerPoolError};
use anyhow::Error;
use futures::StreamExt;
use rdkafka::Message;
use rdkafka::consumer::StreamConsumer;
use rdkafka::error::KafkaResult;
use serde_json::to_string;
use std::sync::Arc;
use tokio::task::JoinHandle;
use tracing::error;

pub struct Transport {
    config: AppConfig,
    pool: WorkerPool,
}

impl Transport {
    pub fn new(config: AppConfig, pool: WorkerPool) -> Self {
        Self { config, pool }
    }

    pub async fn serve(&self) -> std::result::Result<(), Box<dyn std::error::Error>> {
        // infrastructure Layer
        let postgres_pool = get_connection(&self.config.clone()).await;
        let user_service = UserServiceGrpcClient::new(self.config.clone()).await;
        let payment_service = PaymentServiceGrpcClient::new(self.config.clone()).await;
        let rabbitmq_infrastructure =
            Arc::new(RabbitMQInfrastructure::new(self.config.clone()).await);
        let kafka_infrastructure = KafkaInfrastructure::new(self.config.clone());

        // Repository Layer
        let shipping_provider_postgres_repository =
            ShippingProviderPostgresRepositoryImpl::new(postgres_pool.clone());
        let shipping_postgres_repository =
            ShippingPostgresRepositoryImpl::new(postgres_pool.clone());

        // UseCase Layer
        let shipping_use_case = ShippingUseCaseImpl::new(
            self.config.clone(),
            shipping_postgres_repository,
            shipping_provider_postgres_repository,
            user_service,
            payment_service,
            kafka_infrastructure.clone(),
        );

        let shipping_consumer = Arc::new(crate::module::shipping::consumer_kafka::Consumer::new(
            shipping_use_case,
        ));

        let kafka_consumer = match kafka_infrastructure
            .consume(&[
                self.config
                    .message_broker_kafka_topic_shipping
                    .snapshot_shippings_shipping_created
                    .as_str(),
                self.config
                    .message_broker_kafka_topic_shipping
                    .snapshot_shippings_shipping_updated
                    .as_str(),
            ])
            .await
        {
            Ok(v) => Arc::new(v), // Wrap in Arc
            Err(e) => {
                eprintln!("Failed to create Kafka consumer: {:?}", e);
                return Err(e.into());
            }
        };

        let kafka_consumer = match kafka_infrastructure
            .consume(&[
                self.config
                    .message_broker_kafka_topic_shipping
                    .snapshot_shippings_shipping_created
                    .as_str(),
                self.config
                    .message_broker_kafka_topic_shipping
                    .snapshot_shippings_shipping_updated
                    .as_str(),
            ])
            .await
        {
            Ok(v) => Arc::new(v),
            Err(e) => {
                eprintln!("Failed to create Kafka consumer: {:?}", e);
                return Err(e.into());
            }
        };

        let mut handles: Vec<(String, Result<JoinHandle<()>, WorkerPoolError>)> = Vec::new();

        // Move the entire consumer into the async block
        let kafka_consumer = Arc::clone(&kafka_consumer);
        let shipping_consumer = Arc::clone(&shipping_consumer);
        let config = self.config.clone();

        // Create a separate task for the stream processing
        let mut stream = kafka_consumer.stream();

        while let Some(message) = stream.next().await {
            match message {
                Ok(m) => {
                    let shipping_consumer = Arc::clone(&shipping_consumer);
                    let message_topic = m.topic().clone().to_string();

                    if message_topic
                        == config
                            .message_broker_kafka_topic_shipping
                            .snapshot_shippings_shipping_created
                            .as_str()
                    {
                        let handler = Ok(self
                            .pool
                            .spawn(move || async move {
                                match shipping_consumer
                                    .consume_snapshot_shippings_shipping_created(m)
                                    .await
                                {
                                    Ok(_) => {}
                                    Err(e) => {
                                        error!("{:?}", e);
                                    }
                                }
                            })
                            .await
                            .expect("Failed to spawn worker"));

                        handles.push((
                            self.config
                                .message_broker_kafka_topic_shipping
                                .snapshot_shippings_shipping_created
                                .clone(),
                            handler,
                        ));
                    }

                    if message_topic
                        == config
                            .message_broker_kafka_topic_shipping
                            .snapshot_shippings_shipping_updated
                            .as_str()
                    {}

                    continue;
                }
                Err(e) => {
                    error!("Kafka message error: {:?}", e);
                }
            }
        }

        // Wait for the stream processing task to complete
        for (topic, handle) in handles {
            if let Err(e) = handle.unwrap().await {
                error!("[Kafka] Error: {}", e);
            }
        }

        Ok(())
    }
}
