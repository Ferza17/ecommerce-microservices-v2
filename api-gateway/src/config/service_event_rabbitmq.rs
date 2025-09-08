use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct ServiceEventRabbitMQ {
    pub exchange_fanout_event: String,
    pub queue_event_api_gateway_event_created: String,
    key_prefix: String,
}

impl Default for ServiceEventRabbitMQ {
    fn default() -> Self {
        Self {
            exchange_fanout_event: "".to_string(),
            queue_event_api_gateway_event_created: "".to_string(),
            key_prefix: "/broker/rabbitmq".to_string(),
        }
    }
}

impl ServiceEventRabbitMQ {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.exchange_fanout_event = crate::config::config::get_kv(
            client,
            format!("{}{}/EXCHANGE/EVENT", env, self.key_prefix),
        )
        .await;
        self.queue_event_api_gateway_event_created = crate::config::config::get_kv(
            client,
            format!(
                "{}{}/QUEUE/EVENT/API-GATEWAY/EVENT-CREATED",
                env, self.key_prefix
            ),
        )
        .await;

        Ok(self.clone())
    }
}
