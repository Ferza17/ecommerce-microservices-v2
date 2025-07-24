use crate::config::config::AppConfig;

pub struct RabbitMQTransport {
    config: AppConfig,
}

impl RabbitMQTransport {
    pub fn new(cfg: AppConfig) -> Self {
        RabbitMQTransport { config: cfg }
    }

    pub async fn serve(&self) {
        eprintln!("RABBITMQ {} ", self.config.shipping_service_service_name,);
    }
}
