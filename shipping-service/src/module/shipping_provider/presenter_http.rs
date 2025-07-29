use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::interceptor::auth::AuthLayer;
use crate::interceptor::logger::LoggerLayer;
use crate::interceptor::request_id::RequestIdLayer;
use crate::model::rpc::shipping::{
    GetShippingProviderByIdRequest, GetShippingProviderByIdResponse, ListShippingProvidersRequest,
    ListShippingProvidersResponse,
};
use crate::model::rpc::user::AuthUserVerifyAccessControlRequest;
use crate::module::shipping_provider::usecase::{
    ShippingProviderUseCase, ShippingProviderUseCaseImpl,
};
use crate::package::context::auth::get_request_authorization_token_from_header;
use crate::package::context::request_id::get_request_id_from_header;
use crate::util;
use axum::extract::State;
use axum::http::HeaderMap;
use axum::routing::get;
use axum::{
    extract::{Path, Query},
    http::StatusCode,
    response::Json,
};
use prost_validate::NoopValidator;
use std::convert::Infallible;
use std::sync::Arc;
use tonic::{Code, Request};
use tower::ServiceBuilder;
use tracing::{error, instrument};

#[derive(Debug, Clone)]
pub struct ShippingProviderHttpPresenter {
    shipping_provider_use_case: ShippingProviderUseCaseImpl,
    user_service: UserServiceGrpcClient,
}

impl ShippingProviderHttpPresenter {
    pub fn new(
        shipping_provider_use_case: ShippingProviderUseCaseImpl,
        user_service: UserServiceGrpcClient,
    ) -> Self {
        Self {
            shipping_provider_use_case,
            user_service,
        }
    }

    pub fn router(&self) -> axum::Router {
        axum::Router::new()
            .route("/", get(list_shipping_providers))
            .route("/{id}", get(get_shipping_provider_by_id))
            .layer(
                ServiceBuilder::new()
                    .layer(RequestIdLayer)
                    .layer(LoggerLayer)
                    .layer(AuthLayer::new(self.user_service.clone())),
            )
            .with_state(Arc::from(self.clone()))
    }
}

#[utoipa::path(
    get,
    tag = "Shipping Provider",
    params(
        ("page" = u32, Query, description = "shipping provider page $gt 0"),
        ("limit" = u32, Query, description = "shipping provider limit $gt 0")
    ),
    security(
       ("authorization" = [])
    ),
    path = "/v1/shipping/shipping_providers",
    responses(
        (status = OK, body = ListShippingProvidersResponse, content_type = "application/json")
    )
)]
#[instrument(skip(state))]
pub async fn list_shipping_providers(
    State(state): State<Arc<ShippingProviderHttpPresenter>>,
    headers: HeaderMap,
    Query(query): Query<ListShippingProvidersRequest>,
) -> Result<(StatusCode, Json<ListShippingProvidersResponse>), StatusCode> {
    // VALIDATE REQUEST
    let request = Request::new(query);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(Code::InvalidArgument),
                Json(ListShippingProvidersResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }
    // Validate ACL
    match state
        .user_service
        .clone()
        .auth_user_verify_access_control(
            get_request_id_from_header(&headers),
            AuthUserVerifyAccessControlRequest {
                token: get_request_authorization_token_from_header(&headers),
                full_method_name: Some(
                    "/shipping.ShippingProviderService/ListShippingProviders".to_string(),
                ),
                http_url: None,
                http_method: None,
            },
        )
        .await
    {
        Ok(response) => {
            if !response.data.unwrap().is_valid {
                return Ok((
                    util::convert_status::tonic_to_http_status(Code::PermissionDenied),
                    Json(ListShippingProvidersResponse {
                        message: "forbidden".to_string(),
                        status: "error".to_string(),
                        data: None,
                    }),
                ));
            }
        }
        Err(err) => {
            error!("AuthUserVerifyAccessControl failed: {}", err.message());
            return Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                Json(ListShippingProvidersResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    let result = state
        .shipping_provider_use_case
        .list_shipping_providers(get_request_id_from_header(&headers), request)
        .await;

    match result {
        Ok(response) => Ok((
            util::convert_status::tonic_to_http_status(Code::Ok),
            Json(response.into_inner()),
        )),
        Err(err) => {
            error!("ListShippingProviders failed: {}", err.message());
            Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                Json(ListShippingProvidersResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ))
        }
    }
}

#[utoipa::path(
    get,
    path = "/v1/shipping/shipping_providers/{id}",
    tag = "Shipping Provider",
    params(
        ("id" = String, Path, description = "shipping provider id"),
    ),
    responses(
        (status = OK, body = GetShippingProviderByIdResponse, content_type = "application/json" )
    )
)]
#[instrument(skip(state))]
pub async fn get_shipping_provider_by_id(
    State(state): State<Arc<ShippingProviderHttpPresenter>>,
    headers: HeaderMap,
    Path(id): Path<String>,
) -> Result<(StatusCode, Json<GetShippingProviderByIdResponse>), Infallible> {
    // VALIDATE REQUEST
    let request = Request::new(GetShippingProviderByIdRequest { id });
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(Code::InvalidArgument),
                Json(GetShippingProviderByIdResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    // Validate ACL
    match state
        .user_service
        .clone()
        .auth_user_verify_access_control(
            get_request_id_from_header(&headers),
            AuthUserVerifyAccessControlRequest {
                token: get_request_authorization_token_from_header(&headers),
                full_method_name: Some(
                    "/shipping.ShippingProviderService/GetShippingProviderById".to_string(),
                ),
                http_url: None,
                http_method: None,
            },
        )
        .await
    {
        Ok(response) => {
            if !response.data.unwrap().is_valid {
                return Ok((
                    util::convert_status::tonic_to_http_status(Code::PermissionDenied),
                    Json(GetShippingProviderByIdResponse {
                        message: "forbidden".to_string(),
                        status: "error".to_string(),
                        data: None,
                    }),
                ));
            }
        }
        Err(err) => {
            error!("AuthUserVerifyAccessControl failed: {}", err.message());
            return Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                Json(GetShippingProviderByIdResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    let result = state
        .shipping_provider_use_case
        .get_shipping_provider_by_id(get_request_id_from_header(&headers), request)
        .await;

    match result {
        Ok(response) => Ok((
            util::convert_status::tonic_to_http_status(Code::Ok),
            Json(response.into_inner()),
        )),
        Err(err) => {
            error!("GetShippingProviderById failed: {}", err.message());
            Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                Json(GetShippingProviderByIdResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ))
        }
    }
}
