use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::interceptor::auth::AuthLayer;
use crate::interceptor::logger::LoggerLayer;
use crate::interceptor::request_id::RequestIdLayer;
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
    validate_create_shipping_provider, validate_delete_shipping_provider,
    validate_get_shipping_provider_by_id, validate_list_shipping_providers,
    validate_update_shipping_provider,
};
use crate::package::context::auth::get_request_authorization_token_from_header;
use crate::package::context::request_id::get_request_id_from_header;
use axum::extract::State;
use axum::http::HeaderMap;
use axum::routing::{delete, get, post, put};
use axum::{
    extract::{Path, Query},
    http::StatusCode,
    response::Json,
};
use std::convert::Infallible;
use std::sync::Arc;
use tonic::Request;
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
            .route("/", post(create_shipping_provider))
            .route("/", get(list_shipping_providers))
            .route("/{id}", get(get_shipping_provider_by_id))
            .route("/{id}", put(update_shipping_provider))
            .route("/{id}", delete(delete_shipping_provider))
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
    tag="Shipping Provider",
    path = "/v1/shipping/shipping_providers",
    request_body = CreateShippingProviderRequest,
    responses(
        (status = OK, body = CreateShippingProviderResponse, content_type = "application/json")
    )
)]
#[instrument(skip(state))]
pub async fn create_shipping_provider(
    State(state): State<Arc<ShippingProviderHttpPresenter>>,
    headers: HeaderMap,
    Json(payload): Json<CreateShippingProviderRequest>,
) -> Result<Json<CreateShippingProviderResponse>, StatusCode> {
    let request = Request::new(payload);
    if let Some(_status) = validate_create_shipping_provider(&request) {
        error!("Invalid request parameters");
        return Err(StatusCode::BAD_REQUEST);
    }
    let result = state
        .shipping_provider_use_case
        .create_shipping_provider(get_request_id_from_header(&headers), request)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;
    Ok(Json(result.into_inner()))
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
    let request = Request::new(query);

    // Validate access control
    let validate_acl = state
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
        .await;

    match validate_acl {
        Ok(response) => {
            if !response.data.unwrap().is_valid {
                return Ok((
                    StatusCode::FORBIDDEN,
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
                StatusCode::INTERNAL_SERVER_ERROR,
                Json(ListShippingProvidersResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    if let Some(_status) = validate_list_shipping_providers(&request) {
        error!("Invalid request parameters");
        return Ok((
            StatusCode::BAD_REQUEST,
            Json(ListShippingProvidersResponse {
                message: "bad request".to_string(),
                status: "error".to_string(),
                data: None,
            }),
        ));
    }

    let result = state
        .shipping_provider_use_case
        .list_shipping_providers(get_request_id_from_header(&headers), request)
        .await;

    match result {
        Ok(response) => Ok((StatusCode::OK, Json(response.into_inner()))),
        Err(err) => {
            error!("ListShippingProviders failed: {}", err.message());
            Ok((
                StatusCode::INTERNAL_SERVER_ERROR,
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
    let request = Request::new(GetShippingProviderByIdRequest { id });
    if let Some(_status) = validate_get_shipping_provider_by_id(&request) {
        error!("Invalid request parameters");
        return Ok((
            StatusCode::BAD_REQUEST,
            Json(GetShippingProviderByIdResponse {
                message: "bad request".to_string(),
                status: "error".to_string(),
                data: None,
            }),
        ));
    }

    let result = state
        .shipping_provider_use_case
        .get_shipping_provider_by_id(get_request_id_from_header(&headers), request)
        .await;

    match result {
        Ok(response) => Ok((StatusCode::OK, Json(response.into_inner()))),
        Err(err) => {
            error!("GetShippingProviderById failed: {}", err.message());
            Ok((
                StatusCode::INTERNAL_SERVER_ERROR,
                Json(GetShippingProviderByIdResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ))
        }
    }
}

#[utoipa::path(
    put,
    tag = "Shipping Provider",
    path = "/v1/shipping/shipping_providers/{id}",
    responses(
        (status = OK, body = UpdateShippingProviderResponse)
    )
)]
#[instrument(skip(state))]
pub async fn update_shipping_provider(
    State(state): State<Arc<ShippingProviderHttpPresenter>>,
    headers: HeaderMap,
    Path(id): Path<String>,
    Json(payload): Json<UpdateShippingProviderRequest>,
) -> Result<Json<UpdateShippingProviderResponse>, StatusCode> {
    let request = Request::new(UpdateShippingProviderRequest {
        id,
        name: payload.name,
    });
    if let Some(_status) = validate_update_shipping_provider(&request) {
        error!("Invalid request parameters");
        return Err(StatusCode::BAD_REQUEST);
    }

    let result = state
        .shipping_provider_use_case
        .update_shipping_provider(&get_request_id_from_header(&headers), request)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(Json(result.into_inner()))
}

#[utoipa::path(
    delete,
    tag = "Shipping Provider",
    path = "/v1/shipping/shipping_providers/{id}",
    responses(
        (status = OK, body = DeleteShippingProviderResponse, content_type = "application/json")
    )
)]
#[instrument(skip(state))]
pub async fn delete_shipping_provider(
    State(state): State<Arc<ShippingProviderHttpPresenter>>,
    headers: HeaderMap,
    Path(id): Path<String>,
) -> Result<Json<DeleteShippingProviderResponse>, StatusCode> {
    let request = Request::new(DeleteShippingProviderRequest { id });

    if let Some(_status) = validate_delete_shipping_provider(&request) {
        error!("Invalid request parameters");
        return Err(StatusCode::BAD_REQUEST);
    }

    let result = state
        .shipping_provider_use_case
        .delete_shipping_provider(get_request_id_from_header(&headers), request)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(Json(result.into_inner()))
}
