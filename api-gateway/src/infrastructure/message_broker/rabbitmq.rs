use crate::config::config::AppConfig;
use lapin::{BasicProperties, Channel, Connection, ConnectionProperties, options::*};
use std::sync::Arc;

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

    pub async fn publish(
        &self,
        exchange: &str,
        queue: &str,
        message: tonic::Request<tonic::body::BoxBody>,
    ) -> Result<(), anyhow::Error> {
        let body = message.into_inner();

        match self
            .ch
            .basic_publish(
                exchange,
                queue,
                BasicPublishOptions::default(),
                b"",
                BasicProperties::default(),
            )
            .await
        {
            Ok(_) => Ok(()),
            Err(_) => Err(anyhow::Error::msg("Error")),
        }
    }
}
