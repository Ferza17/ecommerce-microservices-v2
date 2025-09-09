use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct ServicePayment {
    pub name: String,
    pub rpc_host: String,
    pub rpc_port: String,
}

impl Default for ServicePayment {
    fn default() -> Self {
        Self {
            name: String::from(""),
            rpc_host: String::from(""),
            rpc_port: String::from(""),
        }
    }
}

impl ServicePayment {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.name =
            crate::config::config::get_kv(client, format!("{}/services/payment/SERVICE_NAME", env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.rpc_host =
            crate::config::config::get_kv(client, format!("{}/services/payment/RPC_HOST", env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        self.rpc_port =
            crate::config::config::get_kv(client, format!("{}/services/payment/RPC_PORT", env))
                .await
                .parse()
                .unwrap_or_else(|_| "".to_string());

        Ok(self.clone())
    }
}
