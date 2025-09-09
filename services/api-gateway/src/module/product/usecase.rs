use crate::model::rpc::product::{
    FindProductsWithPaginationRequest, FindProductsWithPaginationResponse,
};
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct UseCase {
    product_transport_grpc: crate::module::product::transport_grpc::Transport,
}

impl UseCase {
    pub fn new(product_transport_grpc: crate::module::product::transport_grpc::Transport) -> Self {
        Self {
            product_transport_grpc,
        }
    }

    #[instrument("product.usecase.find_products_with_pagination")]
    pub async fn find_products_with_pagination(
        &mut self,
        request_id: String,
        token: String,
        request: tonic::Request<FindProductsWithPaginationRequest>,
    ) -> Result<FindProductsWithPaginationResponse, tonic::Status> {
        match self
            .product_transport_grpc
            .find_products_with_pagination(request_id, token, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }
}
