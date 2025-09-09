use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct ServiceApiGateway {
    pub name: String,
    pub http_host: String,
    pub http_port: String,
    pub metric_http_port: String,
}

impl Default for ServiceApiGateway {
    fn default() -> Self {
        Self {
            name: "".to_string(),
            http_host: "".to_string(),
            http_port: "".to_string(),
            metric_http_port: "".to_string(),
        }
    }
}

impl ServiceApiGateway {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.name = crate::config::config::get_kv(
            client,
            format!("{}/services/api-gateway/SERVICE_NAME", env),
        )
        .await
        .parse()
        .map_err(|e| anyhow::anyhow!("Error Consul :  {:?}", e))?;

        self.http_host = crate::config::config::get_kv(
            client,
            format!("{}/services/api-gateway/HTTP_HOST", env),
        )
        .await
        .parse()
        .map_err(|e| anyhow::anyhow!("Error Consul :  {:?}", e))?;

        self.http_port = crate::config::config::get_kv(
            client,
            format!("{}/services/api-gateway/HTTP_PORT", env),
        )
        .await
        .parse()
        .map_err(|e| anyhow::anyhow!("Error Consul :  {:?}", e))?;

        self.metric_http_port = crate::config::config::get_kv(
            client,
            format!("{}/services/api-gateway/METRIC_HTTP_PORT", env),
        )
        .await
        .parse()
        .map_err(|e| anyhow::anyhow!("Error Consul :  {:?}", e))?;

        Ok(self.clone())
    }
}
