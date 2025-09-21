use crate::config::config::AppConfig;
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::inject_trace_context_to_kafka_headers;
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
use tracing::Span;
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Debug, Clone)]
pub struct KafkaInfrastructure {
    pub kafka_config: ClientConfig,
    schema_registry_settings: schema_registry_converter::async_impl::schema_registry::SrSettings,
}

impl KafkaInfrastructure {
    pub fn new(config: AppConfig) -> Self {
        let mut kafka_config = ClientConfig::new();
        kafka_config
            .set("bootstrap.servers", config.message_broker_kafka.broker_1)
            .set("message.timeout.ms", "5000")
            .set("client.id", config.service_shipping.name)
            .set("heartbeat.interval.ms", "3000");

        let schema_registry_settings =
            schema_registry_converter::async_impl::schema_registry::SrSettings::new(
                config.message_broker_kafka.schema_registry_url.clone(),
            );

        Self {
            kafka_config,
            schema_registry_settings,
        }
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

    pub async fn publish_with_json_schema(
        &self,
        topic: String,
        schema: crate::model::schema_registry::registry::Registry,
        payload: serde_json::Value,
        key: String,
        request_id: Option<String>,
        token: Option<String>,
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

        let mut headers =
            inject_trace_context_to_kafka_headers(OwnedHeaders::new(), &Span::current().context());

        if request_id.is_some() {
            headers = headers.insert(rdkafka::message::Header {
                key: X_REQUEST_ID_HEADER,
                value: Some(request_id.unwrap().as_bytes()),
            });
        }

        if token.is_some() {
            headers = headers.insert(rdkafka::message::Header {
                key: AUTHORIZATION_HEADER,
                value: Some(format!("Bearer {}", token.unwrap()).as_bytes()),
            })
        }

        let producer: FutureProducer = self.kafka_config.create()?;
        match producer
            .send(
                FutureRecord::to(topic.as_str())
                    .key(&key)
                    .headers(headers)
                    .payload(encoded_payload?.as_slice()),
                Duration::from_secs(5),
            )
            .await
        {
            Ok(_) => Ok(()),
            Err(_) => Err(anyhow::Error::msg("Error sending message")),
        }
    }
}
