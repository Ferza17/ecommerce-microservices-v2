use crate::config::config::AppConfig;
use rdkafka::message::ToBytes;
use rdkafka::{
    config::ClientConfig,
    producer::{FutureProducer, FutureRecord},
};
use std::time::Duration;

#[derive(Debug, Clone)]
pub struct KafkaInfrastructure {
    pub kafka_config: ClientConfig,
}

impl KafkaInfrastructure {
    pub fn new(config: AppConfig) -> Self {
        let mut kafka_config = ClientConfig::new();
        kafka_config
            .set("bootstrap.servers", config.message_broker_kafka.broker_1)
            .set("message.timeout.ms", "5000")
            .set("client.id", config.service_shipping.name)
            .set("heartbeat.interval.ms", "3000");

        Self { kafka_config }
    }

    pub async fn publish<K, P>(&self, record: FutureRecord<'_, K, P>) -> Result<(), anyhow::Error>
    where
        K: ToBytes + ?Sized,
        P: ToBytes + ?Sized,
    {
        let producer: FutureProducer = self.kafka_config.create()?;
        match producer.send(record, Duration::from_secs(5)).await {
            Ok(_) => Ok(()),
            Err(_) => Err(anyhow::Error::msg("Error sending message")),
        }
    }
}
