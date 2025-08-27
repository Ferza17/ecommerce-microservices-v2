use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct OpenPolicyAgent {
    pub path: String,
}

impl Default for OpenPolicyAgent {
    fn default() -> Self {
        Self {
            path: "".to_string(),
        }
    }
}

impl OpenPolicyAgent {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.path = crate::config::config::get_kv(client, format!("{}/policy/opa/PATH", env))
            .await
            .parse()
            .map_err(|e| anyhow::anyhow!("Error Consul :  {:?}", e))?;

        Ok(self.clone())
    }
}
