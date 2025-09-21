use opentelemetry::global;
use opentelemetry::propagation::{Extractor, Injector};
use tonic::metadata::{MetadataKey, MetadataMap, MetadataValue};
use tracing::warn;

pub struct MetadataInjector<'a>(pub &'a mut MetadataMap);

impl Injector for MetadataInjector<'_> {
    fn set(&mut self, key: &str, value: String) {
        match MetadataKey::from_bytes(key.as_bytes()) {
            Ok(key) => match MetadataValue::try_from(&value) {
                Ok(value) => {
                    self.0.insert(key, value);
                }
                Err(error) => warn!(value, error = format!("{error:#}"), "parse metadata value"),
            },

            Err(error) => warn!(key, error = format!("{error:#}"), "parse metadata key"),
        }
    }
}

pub struct HeaderExtractor<'a>(pub &'a hyper::HeaderMap);

impl<'a> Extractor for HeaderExtractor<'a> {
    fn get(&self, key: &str) -> Option<&str> {
        self.0.get(key).and_then(|v| v.to_str().ok())
    }
    fn keys(&self) -> Vec<&str> {
        self.0.keys().map(|k| k.as_str()).collect()
    }
}



struct KafkaHeaderInjector<'a> {
    headers: &'a mut rdkafka::message::OwnedHeaders,
}

impl<'a> Injector for KafkaHeaderInjector<'a> {
    fn set(&mut self, key: &str, value: String) {
        // Kafka headers are key-value, values must be &str or &[u8]
        let value_bytes = value.into_bytes();
        *self.headers = std::mem::replace(self.headers, rdkafka::message::OwnedHeaders::new())
            .insert(rdkafka::message::Header {
                key,
                value: Some(&value_bytes),
            });
    }
}

/// Function to inject trace context into Kafka headers
pub fn inject_trace_context_to_kafka_headers(
    mut headers: rdkafka::message::OwnedHeaders,
    ctx: &opentelemetry::Context,
) -> rdkafka::message::OwnedHeaders {
    global::get_text_map_propagator(|prop| {
        let mut injector = KafkaHeaderInjector {
            headers: &mut headers,
        };
        prop.inject_context(ctx, &mut injector);
    });
    headers
}