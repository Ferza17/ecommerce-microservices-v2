use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct MessageBrokerKafkaTopicShipping {
    pub snapshot_shippings_shipping_created: String,
    pub confirm_snapshot_shippings_shipping_created: String,
    pub compensate_snapshot_shippings_shipping_created: String,

    pub snapshot_shippings_shipping_updated: String,
    pub confirm_snapshot_shippings_shipping_updated: String,
    pub compensate_snapshot_shippings_shipping_updated: String,
}

impl Default for MessageBrokerKafkaTopicShipping {
    fn default() -> Self {
        Self {
            snapshot_shippings_shipping_created: "".to_string(),
            confirm_snapshot_shippings_shipping_created: "".to_string(),
            compensate_snapshot_shippings_shipping_created: "".to_string(),
            snapshot_shippings_shipping_updated: "".to_string(),
            confirm_snapshot_shippings_shipping_updated: "".to_string(),
            compensate_snapshot_shippings_shipping_updated: "".to_string(),
        }
    }
}

impl MessageBrokerKafkaTopicShipping {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.snapshot_shippings_shipping_created = crate::config::config::get_kv(
            client,
            format!("{}/broker/kafka/TOPICS/SHIPPING/SHIPPING_CREATED", env),
        )
        .await;
        self.confirm_snapshot_shippings_shipping_created = crate::config::config::get_kv(
            client,
            format!(
                "{}/broker/kafka/TOPICS/SHIPPING/CONFIRM/SHIPPING_CREATED",
                env
            ),
        )
        .await;
        self.compensate_snapshot_shippings_shipping_created = crate::config::config::get_kv(
            client,
            format!(
                "{}/broker/kafka/TOPICS/SHIPPING/COMPENSATE/SHIPPING_CREATED",
                env
            ),
        )
        .await;

        self.snapshot_shippings_shipping_updated = crate::config::config::get_kv(
            client,
            format!("{}/broker/kafka/TOPICS/SHIPPING/SHIPPING_UPDATED", env),
        )
        .await;
        self.confirm_snapshot_shippings_shipping_updated = crate::config::config::get_kv(
            client,
            format!(
                "{}/broker/kafka/TOPICS/SHIPPING/CONFIRM/SHIPPING_UPDATED",
                env
            ),
        )
        .await;
        self.compensate_snapshot_shippings_shipping_updated = crate::config::config::get_kv(
            client,
            format!(
                "{}/broker/kafka/TOPICS/SHIPPING/COMPENSATE/SHIPPING_UPDATED",
                env
            ),
        )
        .await;
        Ok(self.clone())
    }
}
