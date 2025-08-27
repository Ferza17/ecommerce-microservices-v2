use crate::config::config::AppConfig;
use crate::model::rpc::shipping::shipping_provider_service_client::ShippingProviderServiceClient;
use crate::model::rpc::shipping::shipping_service_client::ShippingServiceClient;
#[derive(Debug, Clone)]
pub struct Transport {
    shipping_provider_service_client: ShippingProviderServiceClient<tonic::transport::Channel>,
    shipping_service_client: ShippingServiceClient<tonic::transport::Channel>,
}

impl Transport {
    pub async fn new(config: AppConfig) -> Result<Self, anyhow::Error> {
        let channel = tonic::transport::Channel::from_shared(
            format!(
                "http://{}:{}",
                config.service_shipping.rpc_host, config.service_shipping.rpc_port
            )
            .to_string(),
        )
        .expect("Failed to connect to shipping service")
        .connect()
        .await
        .map_err(|e| panic!("shipping service not connected : {}", e))
        .unwrap();
        Ok(Self {
            shipping_provider_service_client: ShippingProviderServiceClient::new(channel.clone()),
            shipping_service_client: ShippingServiceClient::new(channel),
        })
    }
}
