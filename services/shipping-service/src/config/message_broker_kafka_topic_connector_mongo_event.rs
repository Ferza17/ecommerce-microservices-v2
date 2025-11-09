use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct MessageBrokerKafkaTopicConnectorMongoEvent {
    pub event_envelope: String,
    pub source_connector_event_envelopes: String,
}

impl Default for MessageBrokerKafkaTopicConnectorMongoEvent {
    fn default() -> Self {
        Self {
            event_envelope: "".to_string(),
            source_connector_event_envelopes: "".to_string(),
        }
    }
}

impl MessageBrokerKafkaTopicConnectorMongoEvent {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.event_envelope =
            crate::config::config::get_kv(client, format!("{}/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENT/EVENT_ENVELOPES", env)).await;
        self.source_connector_event_envelopes =
            crate::config::config::get_kv(client, format!("{}/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/EVENT/EVENT_ENVELOPES", env)).await;
        Ok(self.clone())
    }
}