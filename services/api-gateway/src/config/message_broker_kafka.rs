use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct MessageBrokerKafka {
    pub broker_1: String,
}

impl Default for MessageBrokerKafka {
    fn default() -> Self {
        Self {
            broker_1: "".to_string(),
        }
    }
}

impl MessageBrokerKafka {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.broker_1 =
            crate::config::config::get_kv(client, format!("{}/broker/kafka/BROKER_1", env)).await;
        Ok(self.clone())
    }
}
