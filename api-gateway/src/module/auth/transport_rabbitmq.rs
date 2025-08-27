use crate::config::config::AppConfig;
use crate::infrastructure::message_broker::rabbitmq::RabbitMQInfrastructure;

#[derive(Debug, Clone)]
pub struct Transport {
    rabbitmq_infrastructure: RabbitMQInfrastructure
}

impl Transport {
    pub fn new (
        rabbitmq_infrastructure: RabbitMQInfrastructure
    ) -> Self {
        Self {
            rabbitmq_infrastructure
        }
    }
}