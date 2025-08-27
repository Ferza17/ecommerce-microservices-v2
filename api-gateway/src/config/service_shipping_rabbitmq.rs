use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct ServiceShippingRabbitMQ {
    pub exchange_shipping: String,
    pub queue_shipping_created: String,
    pub queue_shipping_updated: String,
}

impl Default for ServiceShippingRabbitMQ {
    fn default() -> Self {
        Self {
            exchange_shipping: "".to_string(),
            queue_shipping_created: "".to_string(),
            queue_shipping_updated: "".to_string(),
        }
    }
}

impl ServiceShippingRabbitMQ {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.exchange_shipping = crate::config::config::get_kv(
            client,
            format!("{}/broker/rabbitmq/EXCHANGE/SHIPPING", env),
        )
        .await;
        self.queue_shipping_created = crate::config::config::get_kv(
            client,
            format!("{}/broker/rabbitmq/QUEUE/SHIPPING/CREATED", env),
        )
        .await;
        self.queue_shipping_updated = crate::config::config::get_kv(
            client,
            format!("{}/broker/rabbitmq/QUEUE/SHIPPING/UPDATED", env),
        )
        .await;

        Ok(self.clone())
    }
}
