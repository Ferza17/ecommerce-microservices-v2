use config::{Config, ConfigError, Environment, File};
use consulrs::client::{ConsulClient, ConsulClientSettingsBuilder};
use consulrs::{kv, service};

use consulrs::api::check::common::AgentServiceCheckBuilder;
use consulrs::api::service::requests::RegisterServiceRequest;
use serde::Deserialize;
use std::env;
use tracing::error;

#[derive(Clone, Debug)]
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

    // FROM CONSUL KV SERVICE/USER
    pub user_service_service_name: String,
    pub user_service_service_rpc_host: String,
    pub user_service_service_rpc_port: String,

    // FROM CONSUL KV SERVICE/USER
    pub payment_service_service_name: String,
    pub payment_service_service_rpc_host: String,
    pub payment_service_service_rpc_port: String,

    // FROM CONSUL JAEGER TELEMETRY
    pub jaeger_telemetry_host: String,
    pub jaeger_telemetry_rpc_port: String,

    // FROM CONSUL RABBITMQ
    pub rabbitmq_username: String,
    pub rabbitmq_password: String,
    pub rabbitmq_host: String,
    pub rabbitmq_port: String,

    // FROM CONSUL RABBITMQ EXCHANGE
    pub exchange_shipping: String,

    // FROM CONSUL RABBITMQ QUEUE
    pub queue_shipping_created: String,
    pub queue_shipping_updated: String,
}

impl Default for AppConfig {
    fn default() -> Self {
        Self {
            env: "".to_string(),
            consul_host: "".to_string(),
            consul_port: "".to_string(),
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
            user_service_service_name: "".to_string(),
            user_service_service_rpc_host: "".to_string(),
            user_service_service_rpc_port: "".to_string(),
            payment_service_service_name: "".to_string(),
            payment_service_service_rpc_host: "".to_string(),
            payment_service_service_rpc_port: "".to_string(),
            jaeger_telemetry_host: "".to_string(),
            jaeger_telemetry_rpc_port: "".to_string(),
            rabbitmq_username: "".to_string(),
            rabbitmq_password: "".to_string(),
            rabbitmq_host: "".to_string(),
            rabbitmq_port: "".to_string(),
            exchange_shipping: "".to_string(),
            queue_shipping_created: "".to_string(),
            queue_shipping_updated: "".to_string(),
        }
    }
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
        let mut app_config = AppConfig::default();
        app_config.env = cfg_env.env;
        app_config.get_config_shipping_service(&client).await;
        app_config.get_config_user_service(&client).await;
        app_config.get_config_payment_service(&client).await;
        app_config.get_config_database_postgres(&client).await;
        app_config.get_config_jaeger_telemetry(&client).await;
        app_config.get_config_rabbitmq(&client).await;
        app_config.get_config_rabbitmq_exchange(&client).await;
        app_config.get_config_rabbitmq_queue(&client).await;

        // Register Consul Config
        app_config
            .register_consul_service(&app_config.clone(), &client)
            .await
            .expect("TODO: panic message");

        eprintln!(
            "Done Loading Config {} From Consul {} : {} ",
            app_config.env, app_config.consul_host, app_config.consul_port
        );

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

    async fn get_config_user_service(&mut self, client: &ConsulClient) {
        self.user_service_service_name =
            Self::get_kv(client, format!("{}/services/user/SERVICE_NAME", self.env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.user_service_service_rpc_host =
            Self::get_kv(client, format!("{}/services/user/RPC_HOST", self.env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.user_service_service_rpc_port =
            Self::get_kv(client, format!("{}/services/user/RPC_PORT", self.env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());
    }

    async fn get_config_payment_service(&mut self, client: &ConsulClient) {
        self.payment_service_service_name = Self::get_kv(
            client,
            format!("{}/services/payment/SERVICE_NAME", self.env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.payment_service_service_rpc_host =
            Self::get_kv(client, format!("{}/services/payment/RPC_HOST", self.env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.payment_service_service_rpc_port =
            Self::get_kv(client, format!("{}/services/payment/RPC_PORT", self.env))
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

    async fn get_config_jaeger_telemetry(&mut self, client: &ConsulClient) {
        self.jaeger_telemetry_host = Self::get_kv(
            client,
            format!("{}/telemetry/jaeger/JAEGER_TELEMETRY_HOST", self.env),
        )
        .await;

        self.jaeger_telemetry_rpc_port = Self::get_kv(
            client,
            format!("{}/telemetry/jaeger/JAEGER_TELEMETRY_RPC_PORT", self.env),
        )
        .await;
    }

    async fn get_config_rabbitmq(&mut self, client: &ConsulClient) {
        self.rabbitmq_username = Self::get_kv(
            client,
            format!("{}/broker/rabbitmq/RABBITMQ_USERNAME", self.env),
        )
        .await;
        self.rabbitmq_password = Self::get_kv(
            client,
            format!("{}/broker/rabbitmq/RABBITMQ_PASSWORD", self.env),
        )
        .await;
        self.rabbitmq_host = Self::get_kv(
            client,
            format!("{}/broker/rabbitmq/RABBITMQ_HOST", self.env),
        )
        .await;
        self.rabbitmq_port = Self::get_kv(
            client,
            format!("{}/broker/rabbitmq/RABBITMQ_PORT", self.env),
        )
        .await;
    }

    async fn get_config_rabbitmq_exchange(&mut self, client: &ConsulClient) {
        self.exchange_shipping = Self::get_kv(
            client,
            format!("{}/broker/rabbitmq/EXCHANGE/SHIPPING", self.env),
        )
        .await;
    }

    async fn get_config_rabbitmq_queue(&mut self, client: &ConsulClient) {
        self.queue_shipping_created = Self::get_kv(
            client,
            format!("{}/broker/rabbitmq/QUEUE/SHIPPING/CREATED", self.env),
        )
        .await;
        self.queue_shipping_updated = Self::get_kv(
            client,
            format!("{}/broker/rabbitmq/QUEUE/SHIPPING/UPDATED", self.env),
        )
        .await;
    }

    async fn register_consul_service(
        &mut self,
        config: &AppConfig,
        client: &ConsulClient,
    ) -> Result<(), Box<dyn std::error::Error>> {
        let svc_addr = format!(
            "{}:{}",
            config.shipping_service_service_http_host, config.shipping_service_service_http_port
        )
        .to_string();

        service::register(
            client,
            config.shipping_service_service_name.as_str(),
            Some(
                RegisterServiceRequest::builder()
                    .address(svc_addr.as_str())
                    .port(
                        config
                            .shipping_service_service_http_port
                            .parse::<u64>()
                            .unwrap(),
                    )
                    .check(
                        AgentServiceCheckBuilder::default()
                            .name("health_check")
                            .interval("30s")
                            .http(format!("{}/v1/shipping/checks", svc_addr.as_str()).as_str())
                            .build()
                            .unwrap(),
                    ),
            ),
        )
        .await?;
        Ok(())
    }
    async fn get_kv(client: &ConsulClient, formatted_key: String) -> String {
        kv::read(client, &*formatted_key, None)
            .await
            .map_err(|e| error!("Error Consul :  {:?}", e))
            .unwrap()
            .response
            .pop()
            .unwrap()
            .value
            .unwrap()
            .try_into()
            .map_err(|e| error!("Error Consul :  {:?}", e))
            .unwrap()
    }
}
