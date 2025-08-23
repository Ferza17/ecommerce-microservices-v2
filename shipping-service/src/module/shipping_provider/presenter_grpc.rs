use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::model::rpc::shipping::shipping_provider_service_server::ShippingProviderService;
use crate::model::rpc::shipping::{
    GetShippingProviderByIdRequest, GetShippingProviderByIdResponse, ListShippingProvidersRequest,
    ListShippingProvidersResponse,
};
use crate::module::shipping_provider::usecase::{
    ShippingProviderUseCase, ShippingProviderUseCaseImpl,
};
use crate::package::context::request_id::get_request_id_from_metadata;
use prost_validate::NoopValidator;
use tonic::{Code, Request, Response, Status};
use tracing::instrument;

#[derive(Debug)]
pub struct ShippingProviderGrpcPresenter {
    shipping_provider_use_case: ShippingProviderUseCaseImpl,
    user_service: UserServiceGrpcClient,
}

impl ShippingProviderGrpcPresenter {
    pub fn new(
        shipping_provider_use_case: ShippingProviderUseCaseImpl,
        user_service: UserServiceGrpcClient,
    ) -> Self {
        Self {
            shipping_provider_use_case,
            user_service,
        }
    }
}

#[tonic::async_trait]
impl ShippingProviderService for ShippingProviderGrpcPresenter {
    #[instrument]
    async fn get_shipping_provider_by_id(
        &self,
        request: Request<GetShippingProviderByIdRequest>,
    ) -> Result<Response<GetShippingProviderByIdResponse>, Status> {
        request
            .validate()
            .map_err(|e| Status::new(Code::InvalidArgument, e.field.to_string()))?;
        self.shipping_provider_use_case
            .get_shipping_provider_by_id(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }

    #[instrument]
    async fn list_shipping_providers(
        &self,
        request: Request<ListShippingProvidersRequest>,
    ) -> Result<Response<ListShippingProvidersResponse>, Status> {
        request
            .validate()
            .map_err(|e| Status::new(Code::InvalidArgument, e.field.to_string()))?;

        // Execute a use case
        self.shipping_provider_use_case
            .list_shipping_providers(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::new(e.code(), e.to_string()))
    }
}
