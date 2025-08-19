use crate::model::rpc::payment::{FindPaymentProvidersRequest, FindPaymentProvidersResponse};
use crate::module::payment::transport_grpc::PaymentTransportGrpc;

#[derive(Debug, Clone)]
pub struct PaymentUseCase {
    payment_transport_grpc: PaymentTransportGrpc,
}
impl PaymentUseCase {
    pub fn new(payment_transport_grpc: PaymentTransportGrpc) -> Self {
        Self {
            payment_transport_grpc,
        }
    }

    pub async fn find_payment_providers(
        &mut self,
        request_id: String,
        token: String,
        request: tonic::Request<FindPaymentProvidersRequest>,
    ) -> Result<FindPaymentProvidersResponse, tonic::Status> {
        match self
            .payment_transport_grpc
            .find_payment_providers(request_id, token, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }
}
