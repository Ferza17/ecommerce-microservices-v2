use once_cell::sync::Lazy;
use prometheus::{Counter, Histogram, register_counter, register_histogram};

// gRPC metrics
pub static GRPC_REQUESTS_TOTAL: Lazy<Counter> = Lazy::new(|| {
    register_counter!("grpc_requests_total", "Total number of gRPC requests")
        .expect("Failed to register grpc_requests_total metric")
});

pub static GRPC_REQUEST_DURATION: Lazy<Histogram> = Lazy::new(|| {
    register_histogram!(
        "grpc_request_duration_seconds",
        "Duration of gRPC requests in seconds"
    )
    .expect("Failed to register grpc_request_duration_seconds metric")
});

// HTTP metrics
pub static HTTP_REQUESTS_TOTAL: Lazy<Counter> = Lazy::new(|| {
    register_counter!("http_requests_total", "Total number of HTTP requests")
        .expect("Failed to register http_requests_total metric")
});

pub static HTTP_REQUEST_DURATION: Lazy<Histogram> = Lazy::new(|| {
    register_histogram!(
        "http_request_duration_seconds",
        "Duration of HTTP requests in seconds"
    )
    .expect("Failed to register http_request_duration_seconds metric")
});

// RabbitMQ metrics
pub static RABBITMQ_MESSAGES_PUBLISHED: Lazy<Counter> = Lazy::new(|| {
    register_counter!(
        "rabbitmq_messages_published_total",
        "Total number of messages published to RabbitMQ"
    )
    .expect("Failed to register rabbitmq_messages_published_total metric")
});

pub static RABBITMQ_MESSAGES_CONSUMED: Lazy<Counter> = Lazy::new(|| {
    register_counter!(
        "rabbitmq_messages_consumed_total",
        "Total number of messages consumed from RabbitMQ"
    )
    .expect("Failed to register rabbitmq_messages_consumed_total metric")
});

// Helper functions for metrics
pub fn record_grpc_request() {
    GRPC_REQUESTS_TOTAL.inc();
}

pub fn record_grpc_duration(duration: f64) {
    GRPC_REQUEST_DURATION.observe(duration);
}

pub fn record_http_request() {
    HTTP_REQUESTS_TOTAL.inc();
}

pub fn record_http_duration(duration: f64) {
    HTTP_REQUEST_DURATION.observe(duration);
}

pub fn record_rabbitmq_published() {
    RABBITMQ_MESSAGES_PUBLISHED.inc();
}

pub fn record_rabbitmq_consumed() {
    RABBITMQ_MESSAGES_CONSUMED.inc();
}
