use crate::config::config::AppConfig;
use std::net::TcpListener;

pub async fn serve_metric_http_collector(config: AppConfig) {
    let end_point: String = config.shipping_service_service_http_host
        + ":"
        + config
            .shipping_service_service_http_port
            .to_string()
            .as_str();
    let listener = TcpListener::bind(end_point).unwrap();

    println!(
        "{} metric collector is listening at port {}",
        config.shipping_service_service_name, config.shipping_service_service_metric_http_port
    );

    for stream in listener.incoming() {
        let _stream = stream.unwrap();
        println!("Connection established!");
    }
}
