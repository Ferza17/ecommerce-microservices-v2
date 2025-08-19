use crate::config::config::AppConfig;
use crate::model::rpc::payment::{
    FindPaymentProvidersRequest, FindPaymentProvidersResponse,
    payment_provider_service_client::PaymentProviderServiceClient,
    payment_service_client::PaymentServiceClient,
};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use tracing::{Level, event};

#[derive(Debug, Clone)]
pub struct PaymentTransportGrpc {
    payment_service_client: PaymentServiceClient<tonic::transport::Channel>,
    payment_provider_service_client: PaymentProviderServiceClient<tonic::transport::Channel>,
}

impl PaymentTransportGrpc {
    pub async fn new(config: AppConfig) -> Result<Self, anyhow::Error> {
        let channel = tonic::transport::Channel::from_shared(
            format!(
                "http://{}:{}",
                config.payment_service_service_rpc_host, config.payment_service_service_rpc_port
            )
            .to_string(),
        )
        .expect("Failed to connect to user service")
        .connect()
        .await
        .map_err(|e| panic!("product service not connected : {}", e))
        .unwrap();
        Ok(Self {
            payment_service_client: PaymentServiceClient::new(channel.clone()),
            payment_provider_service_client: PaymentProviderServiceClient::new(channel),
        })
    }

    pub async fn find_payment_providers(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<FindPaymentProvidersRequest>,
    ) -> Result<FindPaymentProvidersResponse, tonic::Status> {
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );

        match self
            .payment_provider_service_client
            .find_payment_providers(request)
            .await
        {
            Ok(response) => {
                event!(
                    Level::INFO,
                    request_id = request_id,
                    data=?response
                );
                Ok(response.into_inner())
            }
            Err(err) => {
                event!(
                    Level::ERROR,
                    request_id = request_id,
                    error = %err,
                    "Failed to get find_payment_providers"
                );
                Err(err.into())
            }
        }
    }
}
