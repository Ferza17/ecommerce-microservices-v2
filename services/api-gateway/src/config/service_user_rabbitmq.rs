use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct ServiceUserRabbitMQ {
    pub exchange_user: String,
    pub queue_user_login: String,
    pub queue_user_created: String,
}

impl Default for ServiceUserRabbitMQ {
    fn default() -> Self {
        Self {
            exchange_user: "".to_string(),
            queue_user_login: "".to_string(),
            queue_user_created: "".to_string(),
        }
    }
}

impl ServiceUserRabbitMQ {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.exchange_user =
            crate::config::config::get_kv(client, format!("{}/broker/rabbitmq/EXCHANGE/USER", env))
                .await;
        self.queue_user_login = crate::config::config::get_kv(
            client,
            format!("{}/broker/rabbitmq/QUEUE/USER/LOGIN", env),
        )
        .await;
        self.queue_user_created = crate::config::config::get_kv(
            client,
            format!("{}/broker/rabbitmq/QUEUE/USER/CREATED", env),
        )
        .await;

        Ok(self.clone())
    }
}
