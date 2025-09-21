use config::{Config, ConfigError, Environment, File};
use consulrs::client::{ConsulClient, ConsulClientSettingsBuilder};
use consulrs::{kv, service};
use serde::Deserialize;

use crate::config::database_postgres::DatabasePostgres;
use crate::config::message_broker_kafka::MessageBrokerKafka;
use crate::config::message_broker_kafka_topic_sink_shipping::MessageBrokerKafkaTopicSinkShipping;
use crate::config::message_broker_rabbitmq::MessageBrokerRabbitMQ;
use crate::config::service_payment::ServicePayment;
use crate::config::service_shipping::ServiceShipping;
use crate::config::service_shipping_rabbitmq::ServiceShippingRabbitMQ;
use crate::config::service_user::ServiceUser;
use crate::config::telemetry_jaeger::TelemetryJaeger;
use consulrs::api::check::common::AgentServiceCheckBuilder;
use consulrs::api::service::common::AgentServiceConnect;
use consulrs::api::service::requests::RegisterServiceRequest;
use std::env;

#[derive(Clone, Debug)]
pub struct AppConfig {
    // FROM CONFIG/CONFIG.TOML
    pub config_env: ConfigEnv,
    // FROM CONSUL DATABASE POSTGRES
    pub database_postgres: DatabasePostgres,
    // FROM CONSUL KV SERVICES/SHIPPING
    pub service_shipping: ServiceShipping,
    pub service_shipping_rabbitmq: ServiceShippingRabbitMQ,
    // FROM CONSUL KV SERVICE/USER
    pub service_user: ServiceUser,
    // FROM CONSUL KV SERVICE/USER
    pub service_payment: ServicePayment,
    // FROM CONSUL JAEGER TELEMETRY
    pub telemetry_jaeger: TelemetryJaeger,
    // FROM CONSUL RABBITMQ
    pub message_broker_rabbitmq: MessageBrokerRabbitMQ,
    pub message_broker_kafka: MessageBrokerKafka,
    pub message_broker_kafka_topic_sink_shipping: MessageBrokerKafkaTopicSinkShipping,
}

impl Default for AppConfig {
    fn default() -> Self {
        Self {
            config_env: ConfigEnv::default(),
            database_postgres: DatabasePostgres::default(),
            service_shipping: ServiceShipping::default(),
            service_shipping_rabbitmq: ServiceShippingRabbitMQ::default(),
            service_user: ServiceUser::default(),
            service_payment: ServicePayment::default(),
            telemetry_jaeger: TelemetryJaeger::default(),
            message_broker_rabbitmq: MessageBrokerRabbitMQ::default(),
            message_broker_kafka: MessageBrokerKafka::default(),
            message_broker_kafka_topic_sink_shipping: MessageBrokerKafkaTopicSinkShipping::default(
            ),
        }
    }
}

#[derive(Deserialize, Clone, Debug)]
pub struct ConfigEnv {
    pub env: String,
    pub consul_host: String,
    pub consul_port: String,
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
            .with_config_env(cfg_env))


            // .with_database_postgres_from_consul(&client)
            // .with_message_broker_rabbitmq_from_consul(&client)
            // .with_service_payment_from_consul(&client)
            // .with_service_shipping_from_consul(&client)
            // .with_service_shipping_rabbitmq_from_consul(&client)
            // .with_service_user_from_consul(&client)
            // .with_telemetry_jaeger_from_consul(&client)
            // .with_message_broker_kafka_from_consul(&client)
            // .with_message_broker_kafka_topic_sink_shipping_from_consul(&client)
            // .with_register_consul_service(&client)
    }

    fn with_config_env(mut self, env: ConfigEnv) -> Self {
        self.config_env = env;
        self
    }

    pub fn with_database_postgres_from_consul(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.database_postgres = DatabasePostgres::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_database_postgres_from_consul :  {:?}", e);
                    });
            });
        });
        self
    }

    pub fn with_message_broker_rabbitmq_from_consul(mut self, client: &ConsulClient) -> Self {
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

    pub fn with_service_payment_from_consul(mut self, client: &ConsulClient) -> Self {
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

    pub fn with_service_shipping_from_consul(mut self, client: &ConsulClient) -> Self {
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

    pub fn with_service_shipping_rabbitmq_from_consul(mut self, client: &ConsulClient) -> Self {
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

    pub fn with_service_user_from_consul(mut self, client: &ConsulClient) -> Self {
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

    pub fn with_telemetry_jaeger_from_consul(mut self, client: &ConsulClient) -> Self {
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

    pub fn with_register_consul_service(self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                let svc_addr = format!(
                    "{}:{}",
                    self.service_shipping.http_host, self.service_shipping.http_port
                )
                .to_string();

                service::register(
                    client,
                    self.service_shipping.name.as_str(),
                    Some(
                        RegisterServiceRequest::builder()
                            .name(self.service_shipping.name.as_str())
                            .address(svc_addr.as_str())
                            .port(self.service_shipping.http_port.parse::<u64>().unwrap())
                            .tags(vec!["service".to_string(), "rabbitmq-client".to_string()])
                            .id(format!(
                                "{}:{}",
                                self.service_shipping.http_host, self.service_shipping.http_port
                            )
                            .as_str())
                            .check(
                                AgentServiceCheckBuilder::default()
                                    .name("health_check")
                                    .interval("30s")
                                    .timeout("5s")
                                    .deregister_critical_service_after("40s")
                                    .grpc(
                                        format!(
                                            "{}:{}",
                                            self.service_shipping.name, self.service_user.rpc_port
                                        )
                                        .as_str(),
                                    )
                                    .build()
                                    .unwrap(),
                            )
                            .connect(AgentServiceConnect {
                                native: Option::from(true),
                                sidecar_service: None,
                            }),
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

    pub fn with_message_broker_kafka_from_consul(mut self, client: &ConsulClient) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.message_broker_kafka = MessageBrokerKafka::default()
                    .with_consul_client(self.config_env.env.clone(), client)
                    .await
                    .unwrap_or_else(|e| {
                        panic!("Error with_message_broker_kafka :  {:?}", e);
                    });
            });
        });
        self
    }

    pub fn with_message_broker_kafka_topic_sink_shipping_from_consul(
        mut self,
        client: &ConsulClient,
    ) -> Self {
        tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async {
                self.message_broker_kafka_topic_sink_shipping =
                    MessageBrokerKafkaTopicSinkShipping::default()
                        .with_consul_client(self.config_env.env.clone(), client)
                        .await
                        .unwrap_or_else(|e| {
                            panic!(
                                "Error with_message_broker_kafka_topic_sink_shipping :  {:?}",
                                e
                            );
                        });
            });
        });
        self
    }
}

pub async fn get_kv(client: &ConsulClient, formatted_key: String) -> String {
    kv::read(client, &*formatted_key, None)
        .await
        .map_err(|e| panic!("Error Consul :  {:?}", e))
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
