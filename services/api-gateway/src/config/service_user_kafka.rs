use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct ServiceUserKafka {
    pub topic_snapshot_users_user_login: String,
    pub topic_snapshot_users_user_created: String,
}

impl Default for ServiceUserKafka {
    fn default() -> Self {
        Self {
            topic_snapshot_users_user_login: "".to_string(),
            topic_snapshot_users_user_created: "".to_string(),       
        }
    }
}

impl ServiceUserKafka {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.topic_snapshot_users_user_login = crate::config::config::get_kv(
            client,
            format!("{}/broker/kafka/TOPICS/USER/USER_LOGIN", env),
        )
        .await;

        self.topic_snapshot_users_user_created = crate::config::config::get_kv(
            client,
            format!("{}/broker/kafka/TOPICS/USER/USER_CREATED", env),
        )
            .await;

        Ok(self.clone())
    }
}
