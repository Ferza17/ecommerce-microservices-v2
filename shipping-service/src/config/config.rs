use config::{Config, ConfigError, Environment, File};
use consul::Client;
use serde::Deserialize;
use std::env;
use url::Url;

#[derive(Debug, Deserialize, Clone)]
pub struct AppConfig {
    pub env: String,
    pub consul_host: String,
    pub consul_port: String,

    pub shipping_service_service_name: String,
    pub shipping_service_service_rpc_host: String,
    pub shipping_service_service_rpc_port: String,
    pub shipping_service_service_http_host: String,
    pub shipping_service_service_http_port: String,
}

impl AppConfig {
    pub fn set_config(app_env: &str) -> Result<Self, ConfigError> {
        dotenv::dotenv().ok();
        let run_env = env::var("ENV").unwrap_or_else(|_| app_env.parse().unwrap());

        let builder = Config::builder()
            .add_source(File::with_name("config/config.toml"))
            .add_source(Environment::with_prefix("APP").separator("__"))
            .set_override("env", run_env.clone())?;
        
        
        // TODO: Consul Load Config
        // Merge section based on environment, e.g. [local], [production]
        Ok(builder.build()?.get::<AppConfig>(&run_env)?)
    }
}
