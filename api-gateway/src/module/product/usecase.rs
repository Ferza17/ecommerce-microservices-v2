use crate::model::rpc::product::{
    FindProductsWithPaginationRequest, FindProductsWithPaginationResponse,
};
use crate::module::product::transport_grpc::ProductTransportGrpc;

#[derive(Debug, Clone)]
pub struct ProductUseCase {
    product_transport_grpc: ProductTransportGrpc,
}

impl ProductUseCase {
    pub fn new(product_transport_grpc: ProductTransportGrpc) -> Self {
        Self {
            product_transport_grpc,
        }
    }

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
