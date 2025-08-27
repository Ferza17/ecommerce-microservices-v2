use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct ServiceShipping {
    pub name: String,
    pub rpc_host: String,
    pub rpc_port: String,
}

impl Default for ServiceShipping {
    fn default() -> Self {
        Self {
            name: "".to_string(),
            rpc_host: "".to_string(),
            rpc_port: "".to_string(),
        }
    }
}

impl ServiceShipping {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.name = crate::config::config::get_kv(
            client,
            format!("{}/services/shipping/SERVICE_NAME", env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.rpc_host =
            crate::config::config::get_kv(client, format!("{}/services/shipping/RPC_HOST", env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.rpc_port =
            crate::config::config::get_kv(client, format!("{}/services/shipping/RPC_PORT", env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        Ok(self.clone())
    }
}
