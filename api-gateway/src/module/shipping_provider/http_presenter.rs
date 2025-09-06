use crate::interceptor::auth::AuthLayer;
use crate::model::rpc::shipping::{ListShippingProvidersRequest, ListShippingProvidersResponse};
use crate::package::context::{
    auth::get_request_authorization_token_from_header, request_id::get_request_id_from_header,
};
use crate::util;
use axum::extract::{Query, State};
use axum::http::HeaderMap;
use axum::routing::get;
use prost_validate::NoopValidator;
use tower::ServiceBuilder;
use tracing::{error, info, instrument};

#[derive(Debug, Clone)]
pub struct Presenter {
    shipping_provider_use_case: crate::module::shipping_provider::usecase::UseCase,
    auth_use_case: crate::module::auth::usecase::UseCase,
}

pub const ROUTE_PREFIX: &str = "/api/v1/shipping-providers";
pub const TAG: &str = "ShippingProviders";

impl Presenter {
    pub fn new(
        shipping_provider_use_case: crate::module::shipping_provider::usecase::UseCase,
        auth_use_case: crate::module::auth::usecase::UseCase,
    ) -> Self {
        Self {
            shipping_provider_use_case,
            auth_use_case,
        }
    }

    pub fn router(&self) -> axum::Router {
        axum::Router::new()
            .route("/", get(list_shipping_providers))
            .layer(ServiceBuilder::new().layer(AuthLayer::new(self.auth_use_case.clone())))
            .with_state(self.clone())
    }
}

#[utoipa::path(
    get,
    path = ROUTE_PREFIX.to_string(),
    params(
        ("page" = u32, Query, description = "required page $gt 0"),
        ("limit" = u32, Query, description = "required limit $gt 0")
    ),
    security(
       ("x-request-id" = []),
       ("authorization" = [])
    ),
    tag = TAG,
    responses(
        (status = OK, body = ListShippingProvidersResponse, content_type = "application/json" ))
)]
#[instrument("shipping_provider.http_presenter.list_shipping_providers")]
pub async fn list_shipping_providers(
    State(state): State<Presenter>,
    headers: HeaderMap,
    Query(query): Query<ListShippingProvidersRequest>,
) -> Result<
    (
        axum::http::StatusCode,
        axum::Json<ListShippingProvidersResponse>,
    ),
    axum::http::StatusCode,
> {
    // TODO: Validate RBAC

    let request = tonic::Request::new(query);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            error!("ShippingProviders.list_shipping_providers: {:?}", e);
            return Ok((
                util::convert_status::tonic_to_http_status(tonic::Code::InvalidArgument),
                axum::Json(ListShippingProvidersResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }
    match state
        .shipping_provider_use_case
        .clone()
        .list_shipping_providers(
            get_request_id_from_header(&headers),
            get_request_authorization_token_from_header(&headers),
            request,
        )
        .await
    {
        Ok(response) => {
            info!(
                "ShippingProviders.list_shipping_providers: {:?}",
                response.data.as_ref().unwrap()
            );
            Ok((axum::http::StatusCode::OK.into(), axum::Json(response)))
        }
        Err(err) => {
            error!("ShippingProviders.list_shipping_providers: {:?}", err);
            Ok((
                util::convert_status::tonic_to_http_status(err.code()),
                axum::Json(ListShippingProvidersResponse {
                    message: err.message().to_string(),
                    status: "error".to_string(),
                    data: None,
                }),
            ))
        }
    }
}
