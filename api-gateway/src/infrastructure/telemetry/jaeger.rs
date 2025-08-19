use crate::config::config::AppConfig;
use opentelemetry::KeyValue;
use opentelemetry::sdk::Resource;
use opentelemetry::sdk::trace as sdktrace;
use opentelemetry_otlp::WithExportConfig;
use tracing_opentelemetry::OpenTelemetryLayer;
use tracing_subscriber::Registry;
use tracing_subscriber::prelude::*;

pub fn init_tracing(config: AppConfig) -> Result<(), Box<dyn std::error::Error>> {
    let tracer = opentelemetry_otlp::new_pipeline()
        .tracing()
        .with_exporter(
            opentelemetry_otlp::new_exporter().tonic().with_endpoint(
                format!(
                    "http://{}:{}",
                    config.jaeger_telemetry_host, config.jaeger_telemetry_rpc_port
                )
                .to_string(),
            ), // GRPC OTLP Jaeger Endpoint
        )
        .with_trace_config(
            sdktrace::config().with_resource(Resource::new(vec![KeyValue::new(
                opentelemetry_semantic_conventions::resource::SERVICE_NAME,
                config.api_gateway_service_service_name,
            )])),
        )
        .install_batch(opentelemetry::runtime::Tokio)?;

    let otel_layer = OpenTelemetryLayer::new(tracer);
    Registry::default()
        .with(tracing_subscriber::fmt::layer())
        .with(otel_layer)
        .init();

    Ok(())
}
