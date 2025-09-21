use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct MessageBrokerKafkaTopicSinkShipping {
    pub pg_shippings_shippings: String,
    pub pg_shippings_shipping_providers: String,
}

impl Default for MessageBrokerKafkaTopicSinkShipping {
    fn default() -> Self {
        Self {
            pg_shippings_shippings: "".to_string(),
            pg_shippings_shipping_providers: "".to_string(),
        }
    }
}

impl MessageBrokerKafkaTopicSinkShipping {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.pg_shippings_shippings =
            crate::config::config::get_kv(client, format!("{}/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPINGS", env)).await;
        self.pg_shippings_shipping_providers =
            crate::config::config::get_kv(client, format!("{}/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPING-PROVIDERS", env)).await;
        Ok(self.clone())
    }
}
