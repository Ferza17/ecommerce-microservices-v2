use crate::config::{
    message_broker_rabbitmq::MessageBrokerRabbitMQ, open_policy_agent::OpenPolicyAgent,
    service_api_gateway::ServiceApiGateway, service_event::ServiceEvent,
    service_payment::ServicePayment, service_product::ServiceProduct,
    service_shipping::ServiceShipping, service_shipping_rabbitmq::ServiceShippingRabbitMQ,
    service_user::ServiceUser, service_user_rabbitmq::ServiceUserRabbitMQ,
    telemetry_jaeger::TelemetryJaeger,
};
use config::{Config, ConfigError, Environment, File};
use consulrs::api::check::common::AgentServiceCheckBuilder;
use consulrs::api::service::requests::RegisterServiceRequest;
use consulrs::client::{ConsulClient, ConsulClientSettingsBuilder};
use consulrs::{kv, service};
use serde::Deserialize;
use std::env;

#[derive(Clone, Debug)]
pub struct AppConfig {
    // FROM CONFIG/CONFIG.TOML
    pub config_env: ConfigEnv,
    // FROM CONSUL KV SERVICE/API-GATEWAY
    pub service_api_gateway: ServiceApiGateway,
    // FROM CONSUL KV SERVICES/SHIPPING
    pub service_shipping: ServiceShipping,
    pub service_shipping_rabbitmq: ServiceShippingRabbitMQ,
    // FROM CONSUL KV SERVICE/USER
    pub service_user: ServiceUser,
    pub service_user_rabbitmq: ServiceUserRabbitMQ,
    // FROM CONSUL KV SERVICE/PRODUCT
    pub service_product: ServiceProduct,
    // FROM CONSUL KV SERVICE/PAYMENT
    pub service_payment: ServicePayment,
    // FROM CONSUL KV SERVICE/EVENT-STORE
    pub service_event: ServiceEvent,
    // FROM CONSUL JAEGER TELEMETRY
    pub telemetry_jaeger: TelemetryJaeger,
    // FROM CONSUL RABBITMQ
    pub message_broker_rabbitmq: MessageBrokerRabbitMQ,
    // FROM CONSUL OPEN POLICY CONFIG
    pub open_policy_agent: OpenPolicyAgent,
}

impl Default for AppConfig {
    fn default() -> Self {
        Self {
            // FROM CONFIG/CONFIG.TOML
            config_env: ConfigEnv::default(),
            service_shipping: ServiceShipping::default(),
            service_shipping_rabbitmq: ServiceShippingRabbitMQ::default(),
            service_api_gateway: ServiceApiGateway::default(),
            service_user: ServiceUser::default(),
            service_user_rabbitmq: ServiceUserRabbitMQ::default(),
            service_product: ServiceProduct::default(),
            service_payment: ServicePayment::default(),
            service_event: ServiceEvent::default(),
            telemetry_jaeger: TelemetryJaeger::default(),
            message_broker_rabbitmq: MessageBrokerRabbitMQ::default(),
            open_policy_agent: OpenPolicyAgent::default(),
        }
    }
}

#[derive(Deserialize, Clone, Debug)]
pub struct ConfigEnv {
    env: String,
    consul_host: String,
    consul_port: String,
}

impl Default for ConfigEnv {
    fn default() -> Self {
        Self {
            env: "".to_string(),
            consul_host: "".to_string(),
            consul_port: "".to_string(),
        }
    }
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

        Ok(AppConfig::default()
            .with_config_env(cfg_env)
            .with_service_api_gateway(&client)
            .with_service_shipping(&client)
            .with_open_policy_agent(&client)
            .with_service_shipping_rabbitmq(&client)
            .with_service_user(&client)
            .with_service_user_rabbitmq(&client)
            .with_service_payment(&client)
            .with_service_event(&client)
            .with_service_product(&client)
            .with_message_broker_rabbitmq(&client)
            .with_telemetry_jaeger(&client)
            .with_register_consul_service(&client))
    }

    fn with_config_env(mut self, env: ConfigEnv) -> Self {
        self.config_env = env;
        self
    }

    fn with_service_shipping(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.service_shipping = ServiceShipping::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_service_shipping :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_service_shipping_rabbitmq(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.service_shipping_rabbitmq = ServiceShippingRabbitMQ::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_service_shipping_rabbitmq : {:?}", e);
                    });
            });
        });
        self
    }

    fn with_service_user(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.service_user = ServiceUser::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_service_user :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_service_user_rabbitmq(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.service_user_rabbitmq = ServiceUserRabbitMQ::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_service_user_rabbitmq :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_service_api_gateway(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.service_api_gateway = ServiceApiGateway::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_service_user :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_service_payment(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.service_payment = ServicePayment::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_service_payment :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_service_event(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.service_event = ServiceEvent::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_service_payment :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_telemetry_jaeger(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.telemetry_jaeger = TelemetryJaeger::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_telemetry_jaeger :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_service_product(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.service_product = ServiceProduct::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_service_user :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_message_broker_rabbitmq(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.message_broker_rabbitmq = MessageBrokerRabbitMQ::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_message_broker_rabbitmq :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_open_policy_agent(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.open_policy_agent = OpenPolicyAgent::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_open_policy_agent :  {:?}", e);
                    });
            });
        });
        self
    }

    fn with_register_consul_service(self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                let svc_addr = format!(
                    "{}:{}",
                    self.service_api_gateway.http_host, self.service_api_gateway.http_port
                )
                .to_string();

                service::register(
                    client,
                    self.service_api_gateway.name.as_str(),
                    Some(
                        RegisterServiceRequest::builder()
                            .address(svc_addr.as_str())
                            .port(self.service_api_gateway.http_port.parse::<u64>().unwrap())
                            .check(
                                AgentServiceCheckBuilder::default()
                                    .name("health_check")
                                    .interval("30s")
                                    .http(format!("{}/api/v1/checks", svc_addr.as_str()).as_str())
                                    .build()
                                    .unwrap(),
                            ),
                    ),
                )
                .await
                .unwrap_or_else(|e| {
                    panic!("Error Register Consul :  {:?}", e);
                });
            });
        });
        self
    }
}

pub async fn get_kv(client: &ConsulClient, formatted_key: String) -> String {
    kv::read(client, &*formatted_key, None)
        .await
        .map_err(|e| panic!("Error key {} Consul :  {:?}", formatted_key, e))
        .unwrap()
        .response
        .pop()
        .unwrap()
        .value
        .unwrap()
        .try_into()
        .map_err(|e| panic!("Error Consul :  {:?}", e))
        .unwrap()
}
