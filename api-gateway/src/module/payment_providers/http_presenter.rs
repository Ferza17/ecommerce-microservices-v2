use crate::interceptor::auth::AuthLayer;
use crate::model::rpc::payment::{FindPaymentProvidersRequest, FindPaymentProvidersResponse};
use crate::package::context::auth::get_request_authorization_token_from_header;
use crate::package::context::request_id::get_request_id_from_header;
use crate::util;
use axum::extract::{Query, State};
use axum::http::HeaderMap;
use axum::routing::get;
use prost_validate::NoopValidator;
use tower::ServiceBuilder;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct Presenter {
    payment_provider_use_case: crate::module::payment_providers::usecase::UseCase,
    auth_use_case: crate::module::auth::usecase::UseCase,
}

pub const ROUTE_PREFIX: &str = "/api/v1/payment-providers";
pub const TAG: &str = "PaymentProviders";

impl Presenter {
    pub fn new(
        payment_provider_use_case: crate::module::payment_providers::usecase::UseCase,
        auth_use_case: crate::module::auth::usecase::UseCase,
    ) -> Self {
        Self {
            payment_provider_use_case,
            auth_use_case,
        }
    }

    pub fn router(&self) -> axum::Router {
        axum::Router::new()
            .route("/", get(find_payment_providers))
            .layer(ServiceBuilder::new().layer(AuthLayer::new(self.auth_use_case.clone())))
            .with_state(self.clone())
    }
}

#[utoipa::path(
    get,
    path =ROUTE_PREFIX.to_string(),
    params(
        ("name" = String, Query, description = "Optional providers names (?name=abc)"),
    ),
    security(
       ("authorization" = [])
    ),
    tag = TAG,
    responses(
        (status = OK, body = FindPaymentProvidersResponse, content_type = "application/json" ))
)]
#[instrument("payment_providers.http_presenter.find_payment_providers")]
pub async fn find_payment_providers(
    State(mut state): State<Presenter>,
    headers: HeaderMap,
    Query(query): Query<FindPaymentProvidersRequest>,
) -> Result<
    (
        axum::http::StatusCode,
        axum::Json<FindPaymentProvidersResponse>,
    ),
    axum::http::StatusCode,
> {
    // TODO: Validate RBAC IN OPEN POLICY AGENT

    let request = tonic::Request::new(query);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(tonic::Code::InvalidArgument),
                axum::Json(FindPaymentProvidersResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    match state
        .payment_provider_use_case
        .find_payment_providers(
            get_request_id_from_header(&headers),
            get_request_authorization_token_from_header(&headers),
            request,
        )
        .await
    {
        Ok(response) => Ok((axum::http::StatusCode::OK.into(), axum::Json(response))),
        Err(err) => Ok((
            util::convert_status::tonic_to_http_status(err.code()),
            axum::Json(FindPaymentProvidersResponse {
                message: err.message().to_string(),
                status: "error".to_string(),
                data: None,
            }),
        )),
    }
}
