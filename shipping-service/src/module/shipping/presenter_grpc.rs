use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::model::rpc::shipping::shipping_service_server::ShippingService;
use crate::model::rpc::shipping::{
    CreateShippingRequest, CreateShippingResponse, DeleteShippingRequest, DeleteShippingResponse,
    GetShippingByIdRequest, GetShippingByIdResponse, ListShippingRequest, ListShippingResponse,
    UpdateShippingRequest, UpdateShippingResponse,
};
use crate::model::rpc::user::AuthUserVerifyAccessControlRequest;
use crate::module::shipping::usecase::{ShippingUseCase, ShippingUseCaseImpl};
use crate::package::context::auth::get_request_authorization_token_from_metadata;
use crate::package::context::request_id::get_request_id_from_metadata;
use crate::package::context::url_path::get_url_path_from_metadata;
use prost_validate::NoopValidator;
use tonic::{Request, Response, Status};
use tracing::instrument;

#[derive(Debug)]
pub struct ShippingGrpcPresenter {
    shipping_use_case: ShippingUseCaseImpl,
    user_service: UserServiceGrpcClient,
}

impl ShippingGrpcPresenter {
    pub fn new(
        shipping_use_case: ShippingUseCaseImpl,
        user_service: UserServiceGrpcClient,
    ) -> Self {
        Self {
            shipping_use_case,
            user_service,
        }
    }
}

#[tonic::async_trait]
impl ShippingService for ShippingGrpcPresenter {
    #[instrument]
    async fn create_shipping(
        &self,
        request: Request<CreateShippingRequest>,
    ) -> Result<Response<CreateShippingResponse>, Status> {
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

        self.shipping_use_case
            .create_shipping(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::new(e.code(), e.message()))
    }

    #[instrument]
    async fn get_shipping_by_id(
        &self,
        request: Request<GetShippingByIdRequest>,
    ) -> Result<Response<GetShippingByIdResponse>, Status> {
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

        self.shipping_use_case
            .get_shipping_by_id(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::new(e.code(), e.message()))
    }

    #[instrument]
    async fn list_shipping(
        &self,
        request: Request<ListShippingRequest>,
    ) -> Result<Response<ListShippingResponse>, Status> {
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

        self.shipping_use_case
            .list_shipping(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::new(e.code(), e.message()))
    }

    #[instrument]
    async fn update_shipping(
        &self,
        request: Request<UpdateShippingRequest>,
    ) -> Result<Response<UpdateShippingResponse>, Status> {
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

        self.shipping_use_case
            .update_shipping(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::new(e.code(), e.message()))
    }

    #[instrument]
    async fn delete_shipping(
        &self,
        request: Request<DeleteShippingRequest>,
    ) -> Result<Response<DeleteShippingResponse>, Status> {
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

        self.shipping_use_case
            .delete_shipping(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::new(e.code(), e.message()))
    }
}
