use crate::interceptor::{auth::AuthLayer, logger::LoggerLayer, request_id::RequestIdLayer};
use crate::model::rpc::payment::{FindPaymentProvidersRequest, FindPaymentProvidersResponse};
use crate::module::{payment::usecase::PaymentUseCase, user::usecase::UserUseCase};
use crate::package::context::{
    auth::get_request_authorization_token_from_header, request_id::get_request_id_from_header,
};
use crate::util;
use axum::{
    Router,
    extract::{Query, State},
    http::HeaderMap,
    routing::get,
};
use prost_validate::NoopValidator;
use tower::ServiceBuilder;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct PaymentPresenterHttp {
    payment_use_case: PaymentUseCase,
    user_use_case: UserUseCase,
}

impl PaymentPresenterHttp {
    pub fn new(payment_use_case: PaymentUseCase, user_use_case: UserUseCase) -> Self {
        Self {
            payment_use_case,
            user_use_case,
        }
    }

    pub fn payment_provider_router(&self) -> Router {
        Router::new()
            .route("/", get(find_payment_providers))
            .layer(
                ServiceBuilder::new()
                    .layer(RequestIdLayer)
                    .layer(LoggerLayer)
                    .layer(AuthLayer::new(self.user_use_case.clone())),
            )
            .with_state(self.clone())
    }

    pub fn payment_router(self) -> Router {
        Router::new()
            .layer(
                ServiceBuilder::new()
                    .layer(RequestIdLayer)
                    .layer(LoggerLayer)
                    .layer(AuthLayer::new(self.user_use_case.clone())),
            )
            .with_state(self.clone())
    }
}

#[utoipa::path(
    get,
    path = "/api/v1/payment-providers",
    params(
        ("name" = String, Query, description = "Optional providers names (?name=abc)"),
    ),
    security(
       ("authorization" = [])
    ),
    tag = "PaymentProviders",
    responses(
        (status = OK, body = FindPaymentProvidersResponse, content_type = "application/json" ))
)]
#[instrument(skip(state))]
pub async fn find_payment_providers(
    State(mut state): State<PaymentPresenterHttp>,
    headers: HeaderMap,
    Query(query): Query<FindPaymentProvidersRequest>,
) -> Result<
    (
        axum::http::StatusCode,
        axum::Json<FindPaymentProvidersResponse>,
    ),
    axum::http::StatusCode,
> {
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
        .payment_use_case
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
