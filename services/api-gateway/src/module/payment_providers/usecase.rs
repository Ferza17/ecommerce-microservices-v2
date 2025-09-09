use tracing::instrument;
use crate::model::rpc::payment::{FindPaymentProvidersRequest, FindPaymentProvidersResponse};

#[derive(Debug, Clone)]
pub struct UseCase {
    payment_provider_transport_grpc: crate::module::payment_providers::transport_grpc::Transport,
}

impl UseCase {
    pub fn new(
        payment_provider_transport_grpc: crate::module::payment_providers::transport_grpc::Transport,
    ) -> Self {
        Self {
            payment_provider_transport_grpc,
        }
    }

    #[instrument("payment_providers.usecase.find_payment_providers")]
    pub async fn find_payment_providers(
        &mut self,
        request_id: String,
        token: String,
        request: tonic::Request<FindPaymentProvidersRequest>,
    ) -> Result<FindPaymentProvidersResponse, tonic::Status> {
        match self
            .payment_provider_transport_grpc
            .find_payment_providers(request_id, token, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }
}
