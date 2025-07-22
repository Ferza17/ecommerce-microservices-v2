use config::{Config, ConfigError, Environment, File};
use consulrs::client::{ConsulClient, ConsulClientSettingsBuilder};
use consulrs::kv;

use serde::Deserialize;
use std::env;

#[derive(Clone)]
pub struct AppConfig {
    // FROM CONFIG/CONFIG.TOML
    pub env: String,
    pub consul_host: String,
    pub consul_port: String,

    // FROM CONSUL DATABASE POSTGRES
    pub database_postgres_host: String,
    pub database_postgres_port: String,
    pub database_postgres_username: String,
    pub database_postgres_password: String,
    pub database_postgres_database: String,

    // FROM CONSUL KV SERVICES/SHIPPING
    pub shipping_service_service_name: String,
    pub shipping_service_service_rpc_host: String,
    pub shipping_service_service_rpc_port: String,
    pub shipping_service_service_http_host: String,
    pub shipping_service_service_http_port: String,
    pub shipping_service_service_metric_http_port: String,
}

#[derive(Deserialize)]
struct ConfigEnv {
    env: String,
    consul_host: String,
    consul_port: String,
}

impl AppConfig {
    pub async fn new(app_env: &str) -> Result<Self, ConfigError> {
        dotenv::dotenv().ok();
        let run_env = env::var("ENV").unwrap_or_else(|_| app_env.parse().unwrap());

        let builder = Config::builder()
            .add_source(File::with_name("config/config.toml"))
            .add_source(Environment::with_prefix("APP").separator("__"))
            .set_override("env", run_env.clone())?;

        let cfg_env = builder.build()?.get::<ConfigEnv>(&run_env)?;
        let mut app_config = AppConfig {
            env: cfg_env.env.clone(),
            consul_host: cfg_env.consul_host.clone(),
            consul_port: cfg_env.consul_port.clone(),
            database_postgres_host: "".to_string(),
            database_postgres_port: "".to_string(),
            database_postgres_username: "".to_string(),
            database_postgres_password: "".to_string(),
            database_postgres_database: "".to_string(),
            shipping_service_service_name: "".to_string(),
            shipping_service_service_rpc_host: "".to_string(),
            shipping_service_service_rpc_port: "".to_string(),
            shipping_service_service_http_host: "".to_string(),
            shipping_service_service_http_port: "".to_string(),
            shipping_service_service_metric_http_port: "".to_string(),
        };

        // Create a Consul Client
        let client = ConsulClient::new(
            ConsulClientSettingsBuilder::default()
                .address(format!(
                    "http://{}:{}",
                    cfg_env.consul_host, cfg_env.consul_port
                ))
                .build()
                .map_err(|e| eprintln!(" Error Consul :  {:?}", e))
                .unwrap(),
        )
        .unwrap();

        // GET CONSUL CONFIG
        app_config.get_config_shipping_service(&client).await;
        app_config.get_config_database_postgres(&client).await;

        // Merge section based on environment, e.g. [local], [production]
        Ok(app_config)
    }

    async fn get_config_shipping_service(&mut self, client: &ConsulClient) {
        self.shipping_service_service_name = Self::get_kv(
            client,
            format!("{}/services/shipping/SERVICE_NAME", self.env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.shipping_service_service_rpc_host =
            Self::get_kv(client, format!("{}/services/shipping/RPC_HOST", self.env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.shipping_service_service_rpc_port =
            Self::get_kv(client, format!("{}/services/shipping/RPC_PORT", self.env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.shipping_service_service_http_host =
            Self::get_kv(client, format!("{}/services/shipping/HTTP_HOST", self.env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.shipping_service_service_http_port =
            Self::get_kv(client, format!("{}/services/shipping/HTTP_PORT", self.env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.shipping_service_service_metric_http_port = Self::get_kv(
            client,
            format!("{}/services/shipping/METRIC_HTTP_PORT", self.env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());
    }

    async fn get_config_database_postgres(&mut self, client: &ConsulClient) {
        self.database_postgres_host = Self::get_kv(
            client,
            format!("{}/database/postgres/POSTGRES_HOST", self.env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.database_postgres_port = Self::get_kv(
            client,
            format!("{}/database/postgres/POSTGRES_PORT", self.env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.database_postgres_username = Self::get_kv(
            client,
            format!("{}/database/postgres/POSTGRES_USERNAME", self.env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.database_postgres_password = Self::get_kv(
            client,
            format!("{}/database/postgres/POSTGRES_PASSWORD", self.env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.database_postgres_database = Self::get_kv(
            client,
            format!(
                "{}/database/postgres/POSTGRES_DATABASE_NAME/SHIPPINGS",
                self.env
            ),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());
    }

    async fn get_kv(client: &ConsulClient, formatted_key: String) -> String {
        kv::read(client, &*formatted_key, None)
            .await
            .map_err(|e| eprintln!(" Error Consul GET {} :  {:?}", formatted_key, e))
            .unwrap()
            .response
            .pop()
            .unwrap()
            .value
            .unwrap()
            .try_into()
            .unwrap()
    }
}
