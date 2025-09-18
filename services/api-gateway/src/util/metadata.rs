use opentelemetry::global;
use opentelemetry::propagation::{Extractor, Injector};
use std::str::FromStr;
use tonic::metadata::{MetadataKey, MetadataMap, MetadataValue};
use tracing::warn;

pub struct MetadataInjector<'a>(pub &'a mut MetadataMap);

impl<'a> Injector for MetadataInjector<'a> {
    fn set(&mut self, key: &str, value: String) {
        // gRPC metadata keys must be lowercase ASCII
        if let Ok(key) = MetadataKey::from_str(&key.to_ascii_lowercase()) {
            match MetadataValue::from_str(&value) {
                Ok(val) => {
                    self.0.insert(key, val);
                }
                Err(err) => warn!(?value, ?err, "failed to set metadata value"),
            }
        } else {
            warn!(key, "invalid metadata key for tonic");
        }
    }
}
pub fn inject_trace_context<T>(
    mut request: tonic::Request<T>,
    ctx: opentelemetry::Context,
) -> tonic::Request<T> {
    global::get_text_map_propagator(|propagator| {
        propagator.inject_context(&ctx, &mut MetadataInjector(request.metadata_mut()))
    });
    request
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

pub struct AmqpInjector<'a>(&'a mut std::collections::HashMap<String, String>);
impl<'a> Injector for AmqpInjector<'a> {
    fn set(&mut self, key: &str, value: String) {
        self.0.insert(key.to_string(), value);
    }
}

pub fn inject_trace_context_to_lapin_table(
    mut request: lapin::types::FieldTable,
    ctx: opentelemetry::Context,
) -> lapin::types::FieldTable {
    let mut headers = std::collections::HashMap::new();
    global::get_text_map_propagator(|propagator| {
        propagator.inject_context(&ctx, &mut AmqpInjector(&mut headers));
    });

    for (k, v) in headers {
        request.insert(k.into(), lapin::types::AMQPValue::LongString(v.into()));
    }
    request
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
