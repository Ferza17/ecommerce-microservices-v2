use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::model::rpc::shipping::shipping_provider_service_server::ShippingProviderService;
use crate::model::rpc::shipping::{
    GetShippingProviderByIdRequest, GetShippingProviderByIdResponse, ListShippingProvidersRequest,
    ListShippingProvidersResponse,
};
use crate::model::rpc::user::AuthUserVerifyAccessControlRequest;
use crate::module::shipping_provider::usecase::{
    ShippingProviderUseCase, ShippingProviderUseCaseImpl,
};
use crate::package::context::auth::get_request_authorization_token_from_metadata;
use crate::package::context::request_id::get_request_id_from_metadata;
use crate::package::context::url_path::get_url_path_from_metadata;
use prost_validate::NoopValidator;
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
    #[instrument]
    async fn get_shipping_provider_by_id(
        &self,
        request: Request<GetShippingProviderByIdRequest>,
    ) -> Result<Response<GetShippingProviderByIdResponse>, Status> {
        match request.validate() {
            Ok(_) => {}
            Err(e) => {
                return Err(Status::new(
                    tonic::Code::InvalidArgument,
                    format!("Invalid request: {}", e.field),
                ));
            }
        }

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
        // Validate request
        match request.validate() {
            Ok(_) => {}
            Err(e) => {
                return Err(Status::new(
                    tonic::Code::InvalidArgument,
                    format!("Invalid request: {}", e.field),
                ));
            }
        }

        // Validate access control
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

        // Execute a use case
        self.shipping_provider_use_case
            .list_shipping_providers(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::new(e.code(), e.to_string()))
    }
}
