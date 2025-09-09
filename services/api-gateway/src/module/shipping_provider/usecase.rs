use tracing::instrument;
use crate::model::rpc::shipping::{ListShippingProvidersRequest, ListShippingProvidersResponse};

#[derive(Debug, Clone)]
pub struct UseCase {
    shipping_provider_transport_grpc: crate::module::shipping_provider::transport_grpc::Transport,
}

impl UseCase {
    pub fn new(
        shipping_provider_transport_grpc: crate::module::shipping_provider::transport_grpc::Transport,
    ) -> Self {
        Self {
            shipping_provider_transport_grpc,
        }
    }

    #[instrument("shipping_provider.usecase.list_shipping_providers")]
    pub async fn list_shipping_providers(
        &mut self,
        request_id: String,
        token: String,
        request: tonic::Request<ListShippingProvidersRequest>,
    ) -> Result<ListShippingProvidersResponse, tonic::Status> {
        match self
            .shipping_provider_transport_grpc
            .list_shipping_providers(request_id, token, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }
}
