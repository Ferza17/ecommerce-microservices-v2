use crate::config::config::AppConfig;
use rdkafka::config::RDKafkaLogLevel;
use rdkafka::consumer::{Consumer, DefaultConsumerContext, MessageStream, StreamConsumer};
use rdkafka::message::{OwnedHeaders, ToBytes};
use rdkafka::{
    config::ClientConfig,
    producer::{FutureProducer, FutureRecord},
};
use schema_registry_converter::async_impl::schema_registry::post_schema;
use schema_registry_converter::schema_registry_common::{
    SchemaType, SubjectNameStrategy, SuppliedSchema,
};
use std::time::Duration;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct KafkaInfrastructure {
    client_config: ClientConfig,
    schema_registry_settings: schema_registry_converter::async_impl::schema_registry::SrSettings,
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

        // Consumer Config
        client_config
            .set("group.id", config.service_shipping.name.clone().as_str())
            .set("session.timeout.ms", "10000")
            .set("enable.auto.commit", "true")
            .set("auto.offset.reset", "earliest")
            .set("heartbeat.interval.ms", "3000")
            .set_log_level(RDKafkaLogLevel::Debug);

        let schema_registry_settings =
            schema_registry_converter::async_impl::schema_registry::SrSettings::new(
                config.message_broker_kafka.schema_registry_url.clone(),
            );

        Self {
            client_config,
            schema_registry_settings,
        }
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

    #[instrument("KafkaInfrastructure.publish_with_json_schema")]
    pub async fn publish_with_json_schema(
        &self,
        topic: String,
        schema: crate::model::schema_registry::registry::Registry,
        payload: serde_json::Value,
        key: String,
        headers: Option<OwnedHeaders>,
    ) -> Result<(), anyhow::Error> {
        let subject = match SubjectNameStrategy::TopicNameStrategy(topic.clone().to_string(), false)
            .get_subject()
        {
            Ok(v) => v,
            Err(e) => Err(anyhow::Error::msg(e.error))?,
        };

        match post_schema(
            &self.schema_registry_settings.clone(),
            subject,
            SuppliedSchema {
                name: Option::from(String::from(topic.clone().to_string())),
                schema_type: SchemaType::Json,
                schema: schema.to_string(),
                references: vec![],
                properties: None,
                tags: None,
            },
        )
        .await
        {
            Ok(_) => {}
            Err(e) => Err(anyhow::Error::msg(e.error))?,
        }

        let encoded_payload = schema_registry_converter::async_impl::json::JsonEncoder::new(
            self.schema_registry_settings.clone(),
        )
        .encode(
            &payload,
            SubjectNameStrategy::TopicNameStrategy(topic.clone(), false),
        )
        .await;

        let producer: FutureProducer = self.client_config.create()?;
        match producer
            .send(
                FutureRecord::to(topic.as_str())
                    .key(&key)
                    .headers(headers.unwrap_or_default())
                    .payload(encoded_payload?.as_slice()),
                Duration::from_secs(5),
            )
            .await
        {
            Ok(_) => Ok(()),
            Err(_) => Err(anyhow::Error::msg("Error sending message")),
        }
    }

    pub async fn consume(&self, topics: &[&str]) -> Result<StreamConsumer, anyhow::Error> {
        let mut consumer: StreamConsumer = match self.client_config.create() {
            Ok(v) => v,
            Err(_) => Err(anyhow::Error::msg("Error creating consumer"))?,
        };

        consumer
            .subscribe(topics)
            .expect("Can't subscribe to specified topics");

        Ok(consumer)
    }
}
