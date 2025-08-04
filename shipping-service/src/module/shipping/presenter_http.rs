use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::interceptor::auth::AuthLayer;
use crate::interceptor::logger::LoggerLayer;
use crate::interceptor::request_id::RequestIdLayer;
use crate::model::rpc::shipping::{
    CreateShippingRequest, CreateShippingResponse, DeleteShippingRequest, DeleteShippingResponse,
    GetShippingByIdRequest, GetShippingByIdResponse, ListShippingRequest, ListShippingResponse,
    UpdateShippingRequest, UpdateShippingResponse,
};
use crate::model::rpc::user::AuthUserVerifyAccessControlRequest;
use crate::module::shipping::usecase::{ShippingUseCase, ShippingUseCaseImpl};
use crate::package::context::auth::get_request_authorization_token_from_header;
use crate::package::context::request_id::get_request_id_from_header;
use crate::util;
use axum::Json;
use axum::extract::{Path, Query, State};
use axum::http::{HeaderMap, StatusCode};
use axum::routing::{delete, get, post, put};
use prost_validate::NoopValidator;
use std::sync::Arc;
use tonic::Code;
use tower::ServiceBuilder;
use tracing::{error, instrument};

#[derive(Debug, Clone)]
pub struct ShippingHttpPresenter {
    shipping_use_case: ShippingUseCaseImpl,
    user_service: UserServiceGrpcClient,
}

impl ShippingHttpPresenter {
    pub fn new(
        shipping_use_case: ShippingUseCaseImpl,
        user_service: UserServiceGrpcClient,
    ) -> Self {
        Self {
            user_service,
            shipping_use_case,
        }
    }

    pub fn router(&self) -> axum::Router {
        axum::Router::new()
            .route("/", post(create_shipping))
            .route("/", get(list_shipping_providers))
            .route("/{id}", get(get_shipping_provider_by_id))
            .route("/{id}", put(update_shipping))
            .route("/{id}", delete(delete_shipping))
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
    post,
    path = "/v1/shipping/shippings",
    tag = "Shipping",
    request_body = CreateShippingRequest,
    responses(
        (status = OK, body = CreateShippingResponse, content_type = "application/json" )
    )
)]
#[instrument(skip(state))]
pub async fn create_shipping(
    State(state): State<Arc<ShippingHttpPresenter>>,
    headers: HeaderMap,
    Json(payload): Json<CreateShippingRequest>,
) -> Result<(StatusCode, Json<CreateShippingResponse>), StatusCode> {
    let request = tonic::Request::new(payload);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(Code::InvalidArgument),
                Json(CreateShippingResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }
    let validate_acl = state
        .user_service
        .clone()
        .auth_user_verify_access_control(
            get_request_id_from_header(&headers),
            tonic::Request::new(AuthUserVerifyAccessControlRequest {
                token: get_request_authorization_token_from_header(&headers),
                full_method_name: Some(
                    http::uri::PathAndQuery::from_static(
                        "/shipping.ShippingService/CreateShipping",
                    )
                    .to_string(),
                ),
                http_url: None,
                http_method: None,
            }),
        )
        .await;
    match validate_acl {
        Ok(response) => {
            if !response.data.unwrap().is_valid {
                return Ok((
                    util::convert_status::tonic_to_http_status(Code::PermissionDenied),
                    Json(CreateShippingResponse {
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
                Json(CreateShippingResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    let result = state
        .shipping_use_case
        .create_shipping(
            get_request_id_from_header(&headers),
            get_request_authorization_token_from_header(&headers),
            request,
        )
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
                Json(CreateShippingResponse {
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
    path = "/v1/shipping/shippings/{id}",
    tag = "Shipping",
    params(
        ("id" = String, Path, description = "shipping id"),
    ),
    responses(
        (status = OK, body = GetShippingByIdResponse, content_type = "application/json" )
    )
)]
#[instrument(skip(state))]
pub async fn get_shipping_provider_by_id(
    State(state): State<Arc<ShippingHttpPresenter>>,
    headers: HeaderMap,
    Path(id): Path<String>,
) -> Result<(StatusCode, Json<GetShippingByIdResponse>), StatusCode> {
    // VALIDATE REQUEST
    let request = tonic::Request::new(GetShippingByIdRequest { id: id.clone() });
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(Code::InvalidArgument),
                Json(GetShippingByIdResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    // VALIDATE ACL
    match state
        .user_service
        .clone()
        .auth_user_verify_access_control(
            get_request_id_from_header(&headers),
            tonic::Request::new(AuthUserVerifyAccessControlRequest {
                token: get_request_authorization_token_from_header(&headers),
                full_method_name: Some(
                    http::uri::PathAndQuery::from_static(
                        "/shipping.ShippingService/GetShippingById",
                    )
                    .to_string(),
                ),
                http_url: None,
                http_method: None,
            }),
        )
        .await
    {
        Ok(response) => {
            if !response.data.unwrap().is_valid {
                return Ok((
                    util::convert_status::tonic_to_http_status(Code::PermissionDenied),
                    Json(GetShippingByIdResponse {
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
                Json(GetShippingByIdResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    let result = state
        .shipping_use_case
        .get_shipping_by_id(
            get_request_id_from_header(&headers),
            get_request_authorization_token_from_header(&headers),
            request,
        )
        .await;

    match result {
        Ok(response) => Ok((
            util::convert_status::tonic_to_http_status(Code::Ok),
            Json(response.into_inner()),
        )),
        Err(err) => {
            error!("get_shipping_by_id failed: {}", err.message());
            Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                Json(GetShippingByIdResponse {
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
    tag = "Shipping",
    params(
        ("page" = u32, Query, description = "shipping provider page $gt 0"),
        ("limit" = u32, Query, description = "shipping provider limit $gt 0")
    ),
    security(
       ("authorization" = [])
    ),
    path = "/v1/shipping/shippings",
    responses(
        (status = OK, body = ListShippingResponse, content_type = "application/json")
    )
)]
#[instrument(skip(state))]
pub async fn list_shipping_providers(
    State(state): State<Arc<ShippingHttpPresenter>>,
    headers: HeaderMap,
    Query(query): Query<ListShippingRequest>,
) -> Result<(StatusCode, Json<ListShippingResponse>), StatusCode> {
    let request = tonic::Request::new(query);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(Code::InvalidArgument),
                Json(ListShippingResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: vec![],
                }),
            ));
        }
    }
    // VALIDATE ACL
    match state
        .user_service
        .clone()
        .auth_user_verify_access_control(
            get_request_id_from_header(&headers),
            tonic::Request::new(AuthUserVerifyAccessControlRequest {
                token: get_request_authorization_token_from_header(&headers),
                full_method_name: Some(
                    http::uri::PathAndQuery::from_static("/shipping.ShippingService/ListShipping")
                        .to_string(),
                ),
                http_url: None,
                http_method: None,
            }),
        )
        .await
    {
        Ok(response) => {
            if !response.data.unwrap().is_valid {
                return Ok((
                    util::convert_status::tonic_to_http_status(Code::PermissionDenied),
                    Json(ListShippingResponse {
                        message: "forbidden".to_string(),
                        status: "error".to_string(),
                        data: vec![],
                    }),
                ));
            }
        }
        Err(err) => {
            error!("AuthUserVerifyAccessControl failed: {}", err.message());
            return Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                Json(ListShippingResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: vec![],
                }),
            ));
        }
    }

    let result = state
        .shipping_use_case
        .list_shipping(get_request_id_from_header(&headers), request)
        .await;

    match result {
        Ok(response) => Ok((
            util::convert_status::tonic_to_http_status(Code::Ok),
            Json(response.into_inner()),
        )),
        Err(err) => {
            error!("list_shipping failed: {}", err.message());
            Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                Json(ListShippingResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: vec![],
                }),
            ))
        }
    }
}

#[utoipa::path(
    put,
    path = "/v1/shipping/shippings/{id}",
    tag = "Shipping",
    request_body = UpdateShippingRequest,
    params(
        ("id" = String, Path, description = "shipping id"),
    ),
    responses(
        (status = OK, body = UpdateShippingResponse, content_type = "application/json" )
    )
)]
#[instrument(skip(state))]
pub async fn update_shipping(
    State(state): State<Arc<ShippingHttpPresenter>>,
    headers: HeaderMap,
    Path(id): Path<String>,
    Json(payload): Json<UpdateShippingRequest>,
) -> Result<(StatusCode, Json<UpdateShippingResponse>), StatusCode> {
    let request = tonic::Request::new(UpdateShippingRequest {
        id,
        user_id: payload.user_id,
        payment_id: payload.payment_id,
        shipping_provider_id: payload.shipping_provider_id,
    });
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(Code::InvalidArgument),
                Json(UpdateShippingResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }
    // VALIDATE ACL
    match state
        .user_service
        .clone()
        .auth_user_verify_access_control(
            get_request_id_from_header(&headers),
            tonic::Request::new(AuthUserVerifyAccessControlRequest {
                token: get_request_authorization_token_from_header(&headers),
                full_method_name: Some(
                    http::uri::PathAndQuery::from_static(
                        "/shipping.ShippingService/UpdateShipping",
                    )
                    .to_string(),
                ),
                http_url: None,
                http_method: None,
            }),
        )
        .await
    {
        Ok(response) => {
            if !response.data.unwrap().is_valid {
                return Ok((
                    util::convert_status::tonic_to_http_status(Code::PermissionDenied),
                    Json(UpdateShippingResponse {
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
                Json(UpdateShippingResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }
    let result = state
        .shipping_use_case
        .update_shipping(
            get_request_id_from_header(&headers),
            get_request_authorization_token_from_header(&headers),
            request,
        )
        .await;

    match result {
        Ok(response) => Ok((
            util::convert_status::tonic_to_http_status(Code::Ok),
            Json(response.into_inner()),
        )),
        Err(err) => {
            error!("update_shipping failed: {}", err.message());
            Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                Json(UpdateShippingResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ))
        }
    }
}

#[utoipa::path(
    delete,
    path = "/v1/shipping/shippings/{id}",
    tag = "Shipping",
    params(
        ("id" = String, Path, description = "shipping id"),
    ),
    responses(
        (status = OK, body = DeleteShippingResponse, content_type = "application/json" )
    )
)]
#[instrument(skip(state))]
pub async fn delete_shipping(
    State(state): State<Arc<ShippingHttpPresenter>>,
    headers: HeaderMap,
    Path(id): Path<String>,
) -> Result<(StatusCode, Json<DeleteShippingResponse>), StatusCode> {
    let request = tonic::Request::new(DeleteShippingRequest { id });
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(Code::InvalidArgument),
                Json(DeleteShippingResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    match state
        .user_service
        .clone()
        .auth_user_verify_access_control(
            get_request_id_from_header(&headers),
            tonic::Request::new(AuthUserVerifyAccessControlRequest {
                token: get_request_authorization_token_from_header(&headers),
                full_method_name: Some(
                    http::uri::PathAndQuery::from_static(
                        "/shipping.ShippingService/DeleteShipping",
                    )
                    .to_string(),
                ),
                http_url: None,
                http_method: None,
            }),
        )
        .await
    {
        Ok(response) => {
            if !response.data.unwrap().is_valid {
                return Ok((
                    util::convert_status::tonic_to_http_status(Code::PermissionDenied),
                    Json(DeleteShippingResponse {
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
                Json(DeleteShippingResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }
    let result = state
        .shipping_use_case
        .delete_shipping(get_request_id_from_header(&headers), request)
        .await;

    match result {
        Ok(response) => Ok((
            util::convert_status::tonic_to_http_status(Code::Ok),
            Json(response.into_inner()),
        )),
        Err(err) => {
            error!("delete_shipping failed: {}", err.message());
            Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                Json(DeleteShippingResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ))
        }
    }
}
