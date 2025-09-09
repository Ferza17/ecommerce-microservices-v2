use crate::config::config::AppConfig;
use std::net::TcpListener;

pub async fn serve_metric_http_collector(
    config: AppConfig,
) -> std::result::Result<(), Box<dyn std::error::Error>> {
    let end_point = format!(
        "{}:{}",
        config.service_shipping.http_host, config.service_shipping.metric_http_port
    )
    .to_string();

    let listener = TcpListener::bind(end_point.as_str()).unwrap();

    eprintln!("HTTP METRIC COLLECTOR STARTED {}", end_point);

    for stream in listener.incoming() {
        let _stream = stream.unwrap();
        eprintln!("Connection established!");
    }

    Ok(())
}
