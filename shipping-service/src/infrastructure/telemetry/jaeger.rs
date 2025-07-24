use crate::config::config::AppConfig;

use opentelemetry::{
    KeyValue, runtime,
    sdk::{Resource, trace},
    trace::TraceError,
};
use opentelemetry_otlp::WithExportConfig;

pub fn init_tracer(config: AppConfig) -> Result<trace::Tracer, TraceError> {
    // Initialise OTLP Pipeline
    opentelemetry_otlp::new_pipeline()
        .tracing() // create OTLP tracing pipeline
        .with_exporter(
            opentelemetry_otlp::new_exporter()
                .tonic() // create GRPC layer
                .with_endpoint(
                    format!(
                        "http://{}:{}",
                        config.jaeger_telemetry_host, config.jaeger_telemetry_rpc_port
                    )
                        .to_string(),
                ), // GRPC OTLP Jaeger Endpoint
        )
        // Trace provider configuration
        .with_trace_config(trace::config().with_resource(Resource::new(vec![
            KeyValue::new(
                opentelemetry_semantic_conventions::resource::SERVICE_NAME,
                config.shipping_service_service_name.to_string(),
            ),
            KeyValue::new(
                opentelemetry_semantic_conventions::resource::SERVICE_VERSION,
                "v1.0.0".to_string(),
            ),
        ])))
        .install_batch(runtime::Tokio) // configure a span exporter
}
