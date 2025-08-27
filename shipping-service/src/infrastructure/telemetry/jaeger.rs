use crate::config::config::AppConfig;
use opentelemetry::sdk::Resource;
use opentelemetry::sdk::trace as sdktrace;
use opentelemetry::{KeyValue, global};
use opentelemetry_otlp::WithExportConfig;
use tracing_opentelemetry::OpenTelemetryLayer;
use tracing_subscriber::Registry;
use tracing_subscriber::prelude::*;

pub fn init_tracing(config: AppConfig) -> Result<(), Box<dyn std::error::Error>> {
    global::set_text_map_propagator(opentelemetry::sdk::propagation::TraceContextPropagator::new());
    let tracer = opentelemetry_otlp::new_pipeline()
        .tracing()
        .with_exporter(
            opentelemetry_otlp::new_exporter().tonic().with_endpoint(
                format!(
                    "http://{}:{}",
                    config.telemetry_jaeger.host, config.telemetry_jaeger.rpc_port
                )
                .to_string(),
            ),
        )
        .with_trace_config(
            sdktrace::config().with_resource(Resource::new(vec![KeyValue::new(
                opentelemetry_semantic_conventions::resource::SERVICE_NAME,
                config.service_shipping.name.to_string(),
            )])),
        )
        .install_batch(opentelemetry::runtime::Tokio)?;

    let telemetry = OpenTelemetryLayer::new(tracer);
    let subscriber = Registry::default()
        .with(tracing_subscriber::fmt::layer())
        .with(telemetry);

    tracing::subscriber::set_global_default(subscriber).unwrap();
    Ok(())
}
