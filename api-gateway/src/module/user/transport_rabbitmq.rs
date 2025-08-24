use crate::config::config::AppConfig;

#[derive(Debug, Clone)]
pub struct Transport {
    app_config: AppConfig,
}

impl Transport {
    pub fn new(app_config: AppConfig) -> Self {
        Self { app_config }
    }
}
