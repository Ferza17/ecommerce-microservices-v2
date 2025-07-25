use opentelemetry::global;
use opentelemetry::propagation::Injector;
use opentelemetry::trace::TraceContextExt;
use tonic::{
    Request,
    metadata::{MetadataKey, MetadataMap, MetadataValue},
};

pub struct MetadataInjector<'a> {
    metadata: &'a mut MetadataMap,
}

impl<'a> Injector for MetadataInjector<'a> {
    fn set(&mut self, key: &str, value: String) {
        if let Ok(val) = MetadataValue::try_from(value) {
            if let Ok(key) = MetadataKey::from_bytes(key.as_bytes()) {
                let _ = self.metadata.insert(key, val);
            }
        }
    }
}

pub fn grpc_inject_trace_context<T>(mut request: Request<T>) -> Request<T> {
    let context = opentelemetry::Context::current(); // get current trace context

    global::get_text_map_propagator(|propagator| {
        println!("Current span context: {:?}", context.span().span_context());

        let mut injector = MetadataInjector {
            metadata: request.metadata_mut(),
        };
        propagator.inject_context(&context, &mut injector);
    });

    request
}
