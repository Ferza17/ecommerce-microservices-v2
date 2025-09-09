use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct TelemetryJaeger {
    pub host: String,
    pub rpc_port: String,
}

impl Default for TelemetryJaeger {
    fn default() -> Self {
        Self {
            host: "".to_string(),
            rpc_port: "".to_string(),
        }
    }
}

impl TelemetryJaeger {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.host = crate::config::config::get_kv(
            client,
            format!("{}/telemetry/jaeger/JAEGER_TELEMETRY_HOST", env),
        )
            .await;

        self.rpc_port = crate::config::config::get_kv(
            client,
            format!("{}/telemetry/jaeger/JAEGER_TELEMETRY_RPC_PORT", env),
        ).await;
        Ok(self.clone())
    }
}