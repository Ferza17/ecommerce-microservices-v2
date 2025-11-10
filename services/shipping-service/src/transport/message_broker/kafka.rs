use crate::config::config::AppConfig;
use crate::infrastructure::database::async_postgres::get_connection;
use crate::infrastructure::message_broker::kafka::KafkaInfrastructure;
use crate::infrastructure::services::payment::PaymentServiceGrpcClient;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::model::rpc::event::EventEnvelope;
use crate::module::shipping::repository_postgres::ShippingPostgresRepositoryImpl;
use crate::module::shipping::usecase::ShippingUseCaseImpl;
use crate::module::shipping_provider::repository_postgres::ShippingProviderPostgresRepositoryImpl;
use crate::package::worker_pool::worker_pool::WorkerPool;
use futures::StreamExt;
use prost::Message as prostMessage;
use rdkafka::Message;
use std::sync::Arc;
use tracing::{error, info};

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
            Ok(v) => Arc::new(v),
            Err(e) => {
                eprintln!("Failed to create Kafka consumer: {:?}", e);
                return Err(e.into());
            }
        };

        let stream = kafka_consumer.stream();
        tokio::pin!(stream);

        while let Some(message) = stream.next().await {
            match message {
                Ok(m) => match m.topic() {
                    // SHIPPING CREATED
                    topic
                        if topic
                            == self
                                .config
                                .message_broker_kafka_topic_connector_mongo_event
                                .source_connector_event_envelopes
                                .as_str() =>
                    {
                        let consumer = shipping_consumer.clone();
                        match m.detach().payload() {
                            None => {
                                error!("no payload found");
                                continue;
                            }
                            Some(p) => {
                                let mut request = EventEnvelope::default();

                                // list topic
                                let event_type_shipping_created = self
                                    .config
                                    .message_broker_kafka_topic_shipping
                                    .snapshot_shippings_shipping_created
                                    .as_str();
                                let event_type_shipping_updated = self
                                    .config
                                    .message_broker_kafka_topic_shipping
                                    .snapshot_shippings_shipping_updated
                                    .as_str();

                                match EventEnvelope::decode(&*p) {
                                    Ok(v) => match v.event_type {
                                        event_type_shipping_created => {
                                            // TODO: Handle Me
                                            continue;
                                        }
                                        event_type_shipping_updated => {
                                            //TODO: Handle Me
                                            continue;
                                        }
                                        _ => {
                                            error!(
                                                "unregistered aggregate type : {}",
                                                v.aggregate_type
                                            );
                                            continue;
                                        }
                                    },
                                    Err(err) => {
                                        error!(
                                            "[consume_snapshot_shippings_shipping_created] consume_snapshot_shippings_shipping_created : {}",
                                            err
                                        );
                                    }
                                }
                            }
                        }
                    }
                    _ => {
                        eprintln!(
                            "Kafka unregistered topic {} , topic should be inserted into outbox",
                            m.topic()
                        );
                    }
                },
                Err(e) => {
                    eprintln!("Kafka message error: {:?}", e);
                    continue;
                }
            }
        }

        Ok(())
    }
}
