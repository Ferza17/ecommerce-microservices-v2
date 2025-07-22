use crate::config::config::AppConfig;

pub struct HttpTransport {
    config: AppConfig,
}

impl HttpTransport {
    pub fn new(config: AppConfig) -> Self {
        HttpTransport { config }
    }

    pub async fn serve(&self) {
        eprintln!(
            "HTTP {} : {} : {}",
            self.config.shipping_service_service_name,
            self.config.shipping_service_service_http_host,
            self.config.shipping_service_service_http_port
        );
    }
}
