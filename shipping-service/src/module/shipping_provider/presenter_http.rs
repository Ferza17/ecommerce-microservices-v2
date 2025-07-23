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
use std::sync::Arc;
use tonic::Request;

pub async fn get_shipping_provider_by_id(
    State(state): State<Arc<AppState>>,
    headers: HeaderMap,
    Path(id): Path<String>,
) -> Result<Json<GetShippingProviderByIdResponse>, StatusCode> {
    let request = Request::new(GetShippingProviderByIdRequest { id });
    if let Some(_status) = validate_get_shipping_provider_by_id(&request) {
        return Err(StatusCode::BAD_REQUEST);
    }

    let result = state
        .shipping_provider_use_case
        .get_shipping_provider_by_id(get_request_id_from_header(&headers), request)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(Json(result.into_inner()))
}

pub async fn create_shipping_provider(
    State(state): State<Arc<AppState>>,
    headers: HeaderMap,
    Json(payload): Json<CreateShippingProviderRequest>,
) -> Result<Json<CreateShippingProviderResponse>, StatusCode> {
    let request = Request::new(payload);
    if let Some(_status) = validate_create_shipping_provider(&request) {
        return Err(StatusCode::BAD_REQUEST);
    }

    let result = state
        .shipping_provider_use_case
        .create_shipping_provider(get_request_id_from_header(&headers), request)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(Json(result.into_inner()))
}

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
        return Err(StatusCode::BAD_REQUEST);
    }

    let result = state
        .shipping_provider_use_case
        .update_shipping_provider(&get_request_id_from_header(&headers), request)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(Json(result.into_inner()))
}

pub async fn delete_shipping_provider(
    State(state): State<Arc<AppState>>,
    headers: HeaderMap,
    Path(id): Path<String>,
) -> Result<Json<DeleteShippingProviderResponse>, StatusCode> {
    let request = Request::new(DeleteShippingProviderRequest { id });

    if let Some(_status) = validate_delete_shipping_provider(&request) {
        return Err(StatusCode::BAD_REQUEST);
    }

    let result = state
        .shipping_provider_use_case
        .delete_shipping_provider(get_request_id_from_header(&headers), request)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(Json(result.into_inner()))
}

pub async fn list_shipping_providers(
    State(state): State<Arc<AppState>>,
    headers: HeaderMap,
    Query(query): Query<ListShippingProvidersRequest>,
) -> Result<Json<ListShippingProvidersResponse>, StatusCode> {
    let request = Request::new(query);

    if let Some(_status) = validate_list_shipping_providers(&request) {
        return Err(StatusCode::BAD_REQUEST);
    }

    let result = state
        .shipping_provider_use_case
        .list_shipping_providers(get_request_id_from_header(&headers), request)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(Json(result.into_inner()))
}
