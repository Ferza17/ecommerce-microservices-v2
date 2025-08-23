use crate::model::rpc::shipping::{ListShippingProvidersRequest, ListShippingProvidersResponse};
use crate::module::shipping::transport_grpc::ShippingTransportGrpc;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct ShippingUseCase {
    shipping_transport_grpc: ShippingTransportGrpc,
}

impl ShippingUseCase {
    pub fn new(shipping_transport_grpc: ShippingTransportGrpc) -> Self {
        Self {
            shipping_transport_grpc,
        }
    }

    #[instrument("ShippingUseCase.list_shipping_providers")]
    pub async fn list_shipping_providers(
        &mut self,
        request_id: String,
        token: String,
        request: tonic::Request<ListShippingProvidersRequest>,
    ) -> Result<ListShippingProvidersResponse, tonic::Status> {
        match self
            .shipping_transport_grpc
            .list_shipping_providers(request_id, token, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }
}
