use crate::config::config::AppConfig;

#[derive(Debug, Clone)]
pub struct UserTransportRabbitMQ {
    app_config: AppConfig,
}

impl UserTransportRabbitMQ {
    pub fn new(app_config: AppConfig) -> Self {
        Self { app_config }
    }
}
