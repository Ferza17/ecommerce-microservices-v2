use crate::config::config::AppConfig;
use std::net::TcpListener;
use tracing::info;

pub async fn serve_metric_http_collector(config: AppConfig) {
    let end_point = format!(
        "{}:{}",
        config.shipping_service_service_http_host, config.shipping_service_service_metric_http_port
    )
    .to_string();

    let listener = TcpListener::bind(end_point.as_str()).unwrap();

    info!("HTTP METRIC COLLECTOR STARTED {}", end_point);

    for stream in listener.incoming() {
        let _stream = stream.unwrap();
        info!("Connection established!");
    }
}
