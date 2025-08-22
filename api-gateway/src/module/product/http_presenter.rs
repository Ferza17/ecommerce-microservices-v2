use crate::interceptor::{auth::AuthLayer, request_id::RequestIdLayer};
use crate::model::rpc::product::{
    FindProductsWithPaginationRequest, FindProductsWithPaginationResponse,
};
use crate::module::{product::usecase::ProductUseCase, user::usecase::UserUseCase};
use crate::package::context::auth::get_request_authorization_token_from_header;
use crate::package::context::request_id::get_request_id_from_header;
use crate::util;
use axum::extract::Query;
use axum::{Router, extract::State, http::HeaderMap, routing::get};
use prost_validate::NoopValidator;
use tower::ServiceBuilder;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct ProductPresenterHttp {
    product_use_case: ProductUseCase,
    user_use_case: UserUseCase,
}

impl ProductPresenterHttp {
    pub fn new(product_use_case: ProductUseCase, user_use_case: UserUseCase) -> Self {
        Self {
            product_use_case,
            user_use_case,
        }
    }

    pub fn router(&self) -> Router {
        Router::new()
            .route("/", get(find_products_with_pagination))
            .layer(ServiceBuilder::new().layer(RequestIdLayer).layer(AuthLayer))
            .with_state(self.clone())
    }
}

#[utoipa::path(
    get,
    path = "/api/v1/products",
    params(
        ("ids" = Option<Vec<String>>, Query, description = "Optional product ids (?ids=1&ids=2)"),
        ("names" = Option<Vec<String>>, Query, description = "Optional product names (?ids=abc&ids=def)"),
        ("page" = u32, Query, description = "required product page $gt 0"),
        ("limit" = u32, Query, description = "required product limit $gt 0")
    ),
    security(
       ("authorization" = [])
    ),
    tag = "Product",
    responses(
        (status = OK, body = FindProductsWithPaginationResponse, content_type = "application/json" ))
)]
#[instrument("ProductPresenterHttp.find_products_with_pagination")]
pub async fn find_products_with_pagination(
    State(state): State<ProductPresenterHttp>,
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
