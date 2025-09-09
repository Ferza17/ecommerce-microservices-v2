use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct MessageBrokerRabbitMQ {
    pub username: String,
    pub password: String,
    pub host: String,
    pub port: String,
}

impl Default for MessageBrokerRabbitMQ {
    fn default() -> Self {
        Self {
            username: "".to_string(),
            password: "".to_string(),
            host: "".to_string(),
            port: "".to_string(),
        }
    }
}

impl MessageBrokerRabbitMQ {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.username = crate::config::config::get_kv(
            client,
            format!("{}/broker/rabbitmq/RABBITMQ_USERNAME", env),
        )
            .await;
        self.password = crate::config::config::get_kv(
            client,
            format!("{}/broker/rabbitmq/RABBITMQ_PASSWORD", env),
        )
            .await;
        self.host = crate::config::config::get_kv(
            client,
            format!("{}/broker/rabbitmq/RABBITMQ_HOST", env),
        )
            .await;
        self.port = crate::config::config::get_kv(
            client,
            format!("{}/broker/rabbitmq/RABBITMQ_PORT", env),
        )
            .await;
        Ok((self.clone()))
    }
}
