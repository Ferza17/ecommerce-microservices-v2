use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::model::rpc::shipping::shipping_provider_service_server::ShippingProviderService;
use crate::model::rpc::shipping::{
    CreateShippingProviderRequest, CreateShippingProviderResponse, DeleteShippingProviderRequest,
    DeleteShippingProviderResponse, GetShippingProviderByIdRequest,
    GetShippingProviderByIdResponse, ListShippingProvidersRequest, ListShippingProvidersResponse,
    UpdateShippingProviderRequest, UpdateShippingProviderResponse,
};
use crate::model::rpc::user::AuthUserVerifyAccessControlRequest;
use crate::module::shipping_provider::usecase::{
    ShippingProviderUseCase, ShippingProviderUseCaseImpl,
};
use crate::module::shipping_provider::validate::{
    validate_get_shipping_provider_by_id, validate_list_shipping_providers,
};
use crate::package::context::auth::get_request_authorization_token_from_metadata;
use crate::package::context::request_id::get_request_id_from_metadata;
use crate::package::context::url_path::get_url_path_from_metadata;
use tonic::{Request, Response, Status};
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
    #[allow(unused_variables)]
    #[instrument]
    async fn create_shipping_provider(
        &self,
        request: Request<CreateShippingProviderRequest>,
    ) -> Result<Response<CreateShippingProviderResponse>, Status> {
        self.shipping_provider_use_case
            .create_shipping_provider(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }

    #[instrument]
    async fn get_shipping_provider_by_id(
        &self,
        request: Request<GetShippingProviderByIdRequest>,
    ) -> Result<Response<GetShippingProviderByIdResponse>, Status> {
        if let Some(status) = validate_get_shipping_provider_by_id(&request) {
            return Err(status.into());
        }

        self.shipping_provider_use_case
            .get_shipping_provider_by_id(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }

    #[instrument]
    async fn update_shipping_provider(
        &self,
        request: Request<UpdateShippingProviderRequest>,
    ) -> Result<Response<UpdateShippingProviderResponse>, Status> {
        self.shipping_provider_use_case
            .update_shipping_provider(&get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }

    #[instrument]
    async fn delete_shipping_provider(
        &self,
        request: Request<DeleteShippingProviderRequest>,
    ) -> Result<Response<DeleteShippingProviderResponse>, Status> {
        self.shipping_provider_use_case
            .delete_shipping_provider(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }

    #[instrument]
    async fn list_shipping_providers(
        &self,
        request: Request<ListShippingProvidersRequest>,
    ) -> Result<Response<ListShippingProvidersResponse>, Status> {
        if let Some(status) = validate_list_shipping_providers(&request) {
            return Err(status);
        }

        self.user_service
            .clone()
            .auth_user_verify_access_control(
                get_request_id_from_metadata(request.metadata()),
                AuthUserVerifyAccessControlRequest {
                    token: get_request_authorization_token_from_metadata(request.metadata()),
                    full_method_name: Some(get_url_path_from_metadata(request.metadata())),
                    http_url: None,
                    http_method: None,
                },
            )
            .await
            .map_err(|e| Status::from_error(Box::new(e)))?;

        self.shipping_provider_use_case
            .list_shipping_providers(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }
}
