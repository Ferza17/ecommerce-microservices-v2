use crate::model::rpc::shipping::{
    CreateShippingProviderRequest, CreateShippingProviderResponse, DeleteShippingProviderRequest,
    DeleteShippingProviderResponse, GetShippingProviderByIdRequest,
    GetShippingProviderByIdResponse, ListShippingProvidersRequest, ListShippingProvidersResponse,
    UpdateShippingProviderRequest, UpdateShippingProviderResponse,
};
use crate::module::shipping_provider::validate::{
    validate_create_shipping_provider, validate_delete_shipping_provider,
    validate_get_shipping_provider_by_id, validate_list_shipping_providers,
    validate_update_shipping_provider,
};
use crate::package::context::request_id::get_request_id_from_header;
use crate::transport::http::http::AppState;
use axum::extract::State;
use axum::http::HeaderMap;
use axum::{
    extract::{Path, Query},
    http::StatusCode,
    response::Json,
};
use std::convert::Infallible;

use std::sync::Arc;
use tonic::Request;
use tracing::{error, instrument};

#[instrument(skip(state))]
pub async fn get_shipping_provider_by_id(
    State(state): State<Arc<AppState>>,
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

#[instrument(skip(state))]
pub async fn create_shipping_provider(
    State(state): State<Arc<AppState>>,
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

#[instrument(skip(state))]
pub async fn update_shipping_provider(
    State(state): State<Arc<AppState>>,
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

#[instrument(skip(state))]
pub async fn delete_shipping_provider(
    State(state): State<Arc<AppState>>,
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

#[instrument(skip(state))]
pub async fn list_shipping_providers(
    State(state): State<Arc<AppState>>,
    headers: HeaderMap,
    Query(query): Query<ListShippingProvidersRequest>,
) -> Result<(StatusCode, Json<ListShippingProvidersResponse>), Infallible> {
    let request = Request::new(query);

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
                })
            ))
        }
    }
}
