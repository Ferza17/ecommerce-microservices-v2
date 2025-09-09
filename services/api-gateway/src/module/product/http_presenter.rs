use crate::interceptor::auth::AuthLayer;
use crate::model::rpc::product::{
    FindProductsWithPaginationRequest, FindProductsWithPaginationResponse,
};
use crate::package::context::auth::get_request_authorization_token_from_header;
use crate::package::context::request_id::get_request_id_from_header;
use crate::util;
use axum::extract::Query;
use axum::{Router, extract::State, http::HeaderMap, routing::get};
use prost_validate::NoopValidator;
use tower::ServiceBuilder;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct Presenter {
    product_use_case: crate::module::product::usecase::UseCase,
    auth_use_case: crate::module::auth::usecase::UseCase,
}

pub const ROUTE_PREFIX: &str = "/api/v1/products";
pub const TAG: &str = "Product";
impl Presenter {
    pub fn new(
        product_use_case: crate::module::product::usecase::UseCase,
        auth_use_case: crate::module::auth::usecase::UseCase,
    ) -> Self {
        Self {
            product_use_case,
            auth_use_case,
        }
    }

    #[instrument]
    pub fn router(&self) -> Router {
        Router::new()
            .route("/", get(find_products_with_pagination))
            .layer(ServiceBuilder::new().layer(AuthLayer::new(self.auth_use_case.clone())))
            .with_state(self.clone())
    }
}

#[utoipa::path(
    get,
    path = ROUTE_PREFIX.to_string(),
    params(
        ("ids" = Option<Vec<String>>, Query, description = "Optional product ids (?ids=1&ids=2)"),
        ("names" = Option<Vec<String>>, Query, description = "Optional product names (?ids=abc&ids=def)"),
        ("page" = u32, Query, description = "required product page $gt 0"),
        ("limit" = u32, Query, description = "required product limit $gt 0")
    ),
    security(
       ("x-request-id" = []),
       ("authorization" = [])
    ),
    tag = TAG,
    responses(
        (status = OK, body = FindProductsWithPaginationResponse, content_type = "application/json" ))
)]
#[instrument("product.http_presenter.find_products_with_pagination")]
pub async fn find_products_with_pagination(
    State(state): State<Presenter>,
    headers: HeaderMap,
    Query(query): Query<FindProductsWithPaginationRequest>,
) -> Result<
    (
        axum::http::StatusCode,
        axum::Json<FindProductsWithPaginationResponse>,
    ),
    axum::http::StatusCode,
> {
    // TODO: Validate RBAC

    let request = tonic::Request::new(query);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(tonic::Code::InvalidArgument),
                axum::Json(FindProductsWithPaginationResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    match state
        .product_use_case
        .clone()
        .find_products_with_pagination(
            get_request_id_from_header(&headers),
            get_request_authorization_token_from_header(&headers),
            request,
        )
        .await
    {
        Ok(response) => Ok((axum::http::StatusCode::OK.into(), axum::Json(response))),
        Err(err) => Ok((
            util::convert_status::tonic_to_http_status(err.code()),
            axum::Json(FindProductsWithPaginationResponse {
                message: err.message().to_string(),
                status: "error".to_string(),
                data: None,
            }),
        )),
    }
}
