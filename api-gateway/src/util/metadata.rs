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