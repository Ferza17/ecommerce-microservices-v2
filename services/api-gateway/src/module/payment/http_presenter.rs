use crate::interceptor::auth::AuthLayer;
use crate::model::rpc::payment::{CreatePaymentRequest, CreatePaymentResponse};
use crate::package::context::{
    auth::get_request_authorization_token_from_header, request_id::get_request_id_from_header,
};
use crate::util;
use axum::routing::post;
use axum::{Router, extract::State, http::HeaderMap};
use prost_validate::NoopValidator;
use tower::ServiceBuilder;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct Presenter {
    payment_use_case: crate::module::payment::usecase::UseCase,
    auth_use_case: crate::module::auth::usecase::UseCase,
}

pub const ROUTE_PREFIX: &'static str = "/api/v1/payments";
pub const TAG: &'static str = "Payment";

impl Presenter {
    pub fn new(
        payment_use_case: crate::module::payment::usecase::UseCase,
        auth_use_case: crate::module::auth::usecase::UseCase,
    ) -> Self {
        Self {
            payment_use_case,
            auth_use_case,
        }
    }

    #[instrument]
    pub fn router(self) -> Router {
        Router::new()
            .route("/", post(create_payment))
            .layer(ServiceBuilder::new().layer(AuthLayer::new(self.auth_use_case.clone())))
            .with_state(self.clone())
    }
}

#[utoipa::path(
    post,
    path = ROUTE_PREFIX.to_string(),
    tag = TAG,
    request_body = CreatePaymentRequest,
    security(
       ("x-request-id" = []),
       ("authorization" = [])
    ),
    responses(
        (status = OK, body = CreatePaymentResponse, content_type = "application/json" )
    )
)]
#[instrument("payment.http_presenter.create_payment")]
pub async fn create_payment(
    State(mut state): State<Presenter>,
    headers: HeaderMap,
    axum::Json(payload): axum::Json<CreatePaymentRequest>,
) -> Result<(axum::http::StatusCode, axum::Json<CreatePaymentResponse>), axum::http::StatusCode> {
    // TODO: Validate RBAC

    let request = tonic::Request::new(payload);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(tonic::Code::InvalidArgument),
                axum::Json(CreatePaymentResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }
    match state
        .payment_use_case
        .create_payment(
            get_request_id_from_header(&headers),
            get_request_authorization_token_from_header(&headers),
            request,
        )
        .await
    {
        Ok(response) => Ok((axum::http::StatusCode::OK.into(), axum::Json(response))),
        Err(err) => Ok((
            util::convert_status::tonic_to_http_status(err.code()),
            axum::Json(CreatePaymentResponse {
                message: err.message().to_string(),
                status: "error".to_string(),
                data: None,
            }),
        )),
    }
}
