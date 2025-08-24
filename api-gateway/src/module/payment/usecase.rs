use crate::model::{
    rpc::payment::{
        CreatePaymentRequest, CreatePaymentResponse, FindPaymentProviderByIdRequest,
        FindPaymentProvidersRequest, FindPaymentProvidersResponse,
    },
    rpc::shipping::GetShippingProviderByIdRequest,
    rpc::user::FindUserByIdRequest,
};

use tonic::Status;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct UseCase {
    payment_transport_grpc: crate::module::payment::transport_grpc::Transport,
    payment_provider_transport_grpc: crate::module::payment_providers::transport_grpc::Transport,
    shipping_provider_transport_grpc: crate::module::shipping_provider::transport_grpc::Transport,
    user_transport_grpc: crate::module::user::transport_grpc::Transport,
    products_transport_grpc: crate::module::product::transport_grpc::Transport,
}
impl UseCase {
    pub fn new(
        payment_transport_grpc: crate::module::payment::transport_grpc::Transport,
        payment_provider_transport_grpc: crate::module::payment_providers::transport_grpc::Transport,
        shipping_provider_transport_grpc: crate::module::shipping_provider::transport_grpc::Transport,
        user_transport_grpc: crate::module::user::transport_grpc::Transport,
        products_transport_grpc: crate::module::product::transport_grpc::Transport,
    ) -> Self {
        Self {
            payment_transport_grpc,
            payment_provider_transport_grpc,
            shipping_provider_transport_grpc,
            user_transport_grpc,
            products_transport_grpc,
        }
    }

    #[instrument("payment.usecase.create_payment")]
    pub async fn create_payment(
        &mut self,
        request_id: String,
        token: String,
        request: tonic::Request<CreatePaymentRequest>,
    ) -> Result<CreatePaymentResponse, tonic::Status> {
        // validate user
        match self
            .user_transport_grpc
            .find_user_by_id(
                request_id.clone(),
                token.clone(),
                tonic::Request::new(FindUserByIdRequest {
                    id: request.get_ref().user_id.clone(),
                }),
            )
            .await
        {
            Ok(_) => {}
            Err(e) => return Err(e.into()),
        }

        // validate payment provider
        match self
            .payment_provider_transport_grpc
            .find_payment_provider_by_id(
                request_id.clone(),
                token.clone(),
                tonic::Request::new(FindPaymentProviderByIdRequest {
                    id: request.get_ref().provider_id.clone(),
                }),
            )
            .await
        {
            Ok(_) => {}
            Err(e) => return Err(e.into()),
        }

        // validate shipping Provider
        match self
            .shipping_provider_transport_grpc
            .get_shipping_provider_by_id(
                request_id.clone(),
                token.clone(),
                tonic::Request::new(GetShippingProviderByIdRequest {
                    id: request.get_ref().shipping_provider_id.clone(),
                }),
            )
            .await
        {
            Ok(_) => {}
            Err(e) => return Err(e.into()),
        }

        // Validate items
        for item in request.get_ref().items.iter() {
            match self
                .products_transport_grpc
                .clone()
                .find_product_by_id(
                    request_id.clone(),
                    token.clone(),
                    tonic::Request::new(crate::model::rpc::product::FindProductByIdRequest {
                        id: item.product_id.clone(),
                    }),
                )
                .await
            {
                Ok(product) => {
                    if item.qty > product.stock as i32 {
                        return Err(Status::invalid_argument("stock is not enough"));
                    }
                }
                Err(e) => return Err(e.into()),
            }
        }

        match self
            .payment_transport_grpc
            .create_payment(request_id, token, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }
}
