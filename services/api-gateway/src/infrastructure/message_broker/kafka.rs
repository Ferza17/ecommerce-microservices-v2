use crate::config::config::AppConfig;
use rdkafka::consumer::{Consumer, StreamConsumer};
use rdkafka::message::ToBytes;
use rdkafka::{
    config::ClientConfig,
    producer::{FutureProducer, FutureRecord},
};
use std::time::Duration;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct KafkaInfrastructure {
    client_config: ClientConfig,
}

impl KafkaInfrastructure {
    pub fn new(config: AppConfig) -> Self {
        let mut client_config = ClientConfig::new();
        client_config.set(
            "bootstrap.servers",
            config.message_broker_kafka.broker_1.clone(),
        );
        // Published Config
        client_config
            .set("client.id", config.service_shipping.name.clone())
            .set("message.timeout.ms", "2000");

        Self { client_config }
    }

    #[instrument("KafkaInfrastructure.publish")]
    pub async fn publish<K, P>(&self, record: FutureRecord<'_, K, P>) -> Result<(), anyhow::Error>
    where
        K: ToBytes + ?Sized + std::fmt::Debug,
        P: ToBytes + ?Sized + std::fmt::Debug,
    {
        let producer: FutureProducer = self.client_config.create()?;
        match producer.send(record, Duration::from_secs(5)).await {
            Ok(_) => Ok(()),
            Err(_) => Err(anyhow::Error::msg("Error sending message")),
        }
    }
}
