use crate::config::config::AppConfig;
use lapin::types::FieldTable;
use lapin::{BasicProperties, Channel, Connection, ConnectionProperties, ExchangeKind, options::*};
use std::sync::Arc;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct RabbitMQInfrastructure {
    pub ch: Arc<Channel>,
}

impl RabbitMQInfrastructure {
    pub async fn new(config: AppConfig) -> Self {
        let addr = format!(
            "amqp://{}:{}@{}:{}",
            config.message_broker_rabbitmq.username,
            config.message_broker_rabbitmq.password,
            config.message_broker_rabbitmq.host,
            config.message_broker_rabbitmq.port
        );
        let conn = Connection::connect(&addr, ConnectionProperties::default())
            .await
            .map_err(|e| panic!("Cannot connect to RabbitMQ: {}", e))
            .unwrap();
        Self {
            ch: Arc::new(conn.create_channel().await.unwrap()),
        }
    }

    #[instrument("message_broker.rabbitmq.publish")]
    pub async fn publish(
        &self,
        exchange: &str,
        queue: &str,
        properties: BasicProperties,
        message: &[u8],
    ) -> Result<(), anyhow::Error> {
        match self
            .ch
            .basic_publish(
                exchange,
                queue,
                BasicPublishOptions::default(),
                message,
                properties,
            )
            .await
        {
            Ok(_) => Ok(()),
            Err(_) => Err(anyhow::Error::msg(format!(
                "Cannot publish message to RabbitMQ exchange:{} , queue:{}",
                exchange, queue
            ))),
        }
    }

    #[instrument("message_broker.rabbitmq.binding")]
    pub async fn binding(&self, queue: &str, exchange: &str, exchange_kind: ExchangeKind) -> &RabbitMQInfrastructure {
        self.ch
            .exchange_declare(
                exchange,
                exchange_kind,
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

        self.ch
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

        self.ch
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

    pub async fn setup_consumer(&self, queue: &str) -> lapin::Consumer {
        self.ch
            .basic_consume(queue, "", Default::default(), Default::default())
            .await
            .unwrap()
    }
}
