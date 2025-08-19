use crate::config::config::AppConfig;
use lapin::{
    Connection, ConnectionProperties, Consumer, ExchangeKind, options::*, types::FieldTable,
};
use std::sync::Arc;

#[derive(Clone)]
pub struct RabbitMQInfrastructure {
    pub connection: Arc<Connection>,
}

impl RabbitMQInfrastructure {
    pub async fn new(config: AppConfig) -> Self {
        let addr = format!(
            "amqp://{}:{}@{}:{}",
            config.rabbitmq_username,
            config.rabbitmq_password,
            config.rabbitmq_host,
            config.rabbitmq_port
        );
        let conn = Connection::connect(&addr, ConnectionProperties::default())
            .await
            .map_err(|e| panic!("Cannot connect to RabbitMQ: {}", e))
            .unwrap();
        Self {
            connection: Arc::new(conn),
        }
    }

    pub async fn binding(&self, queue: &str, exchange: &str) -> &RabbitMQInfrastructure {
        self.connection
            .create_channel()
            .await
            .unwrap()
            .exchange_declare(
                exchange,
                ExchangeKind::Direct,
                ExchangeDeclareOptions {
                    passive: true,
                    durable: true,
                    auto_delete: true,
                    internal: true,
                    nowait: true,
                },
                FieldTable::default(),
            )
            .await
            .map_err(|e| panic!("Cannot Declare Exchange: {}", e))
            .unwrap();

        self.connection
            .create_channel()
            .await
            .unwrap()
            .queue_declare(
                queue,
                QueueDeclareOptions {
                    passive: true,
                    durable: true,
                    exclusive: true,
                    auto_delete: true,
                    nowait: true,
                },
                FieldTable::default(),
            )
            .await
            .map_err(|e| panic!("Cannot declare Queue: {}", e))
            .unwrap();

        self.connection
            .create_channel()
            .await
            .unwrap()
            .queue_bind(
                queue,
                exchange,
                "",
                QueueBindOptions {
                    nowait: true,
                    ..Default::default()
                },
                FieldTable::default(),
            )
            .await
            .map_err(|e| panic!("Cannot binding queue: {}", e))
            .unwrap();
        self
    }
    pub async fn setup_consumer(&self, queue: &str) -> Consumer {
        self.connection
            .create_channel()
            .await
            .unwrap()
            .basic_consume(queue, "", Default::default(), Default::default())
            .await
            .unwrap()
    }
}
