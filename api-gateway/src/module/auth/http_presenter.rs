use crate::model::rpc::response::Response;
use crate::model::rpc::user::{
    AuthUserLoginByEmailAndPasswordRequest, AuthUserRegisterRequest, AuthUserRegisterResponse,
    AuthUserVerifyOtpRequest, AuthUserVerifyOtpResponse,
};
use crate::package::context::request_id::get_request_id_from_header;
use crate::util;
use axum::{extract::State, http::HeaderMap, routing::post};
use prost_validate::NoopValidator;
use tracing::instrument;
#[derive(Debug, Clone)]
pub struct Presenter {
    auth_use_case: crate::module::auth::usecase::UseCase,
}
pub const ROUTE_PREFIX: &str = "/api/v1/auth";
pub const TAG: &str = "Auth";

impl Presenter {
    pub fn new(auth_use_case: crate::module::auth::usecase::UseCase) -> Self {
        Self { auth_use_case }
    }

    pub fn router(&self) -> axum::Router {
        axum::Router::new()
            .route("/register", post(auth_register))
            .route("/login", post(auth_user_login_by_email_and_password))
            .route("/verify-otp", post(auth_user_verify_otp))
            .with_state(self.clone())
    }
}

#[utoipa::path(
    post,
    path =format!("{}/register", ROUTE_PREFIX.to_string()),
    tag = TAG,
    request_body = AuthUserRegisterRequest,
    security(
       ("x-request-id" = [])
    ),
    responses(
        (status = OK, body = AuthUserRegisterResponse, content_type = "application/json" )
    )
)]
#[instrument("auth.http_presenter.auth_register")]
pub async fn auth_register(
    State(state): State<Presenter>,
    headers: HeaderMap,
    axum::Json(payload): axum::Json<AuthUserRegisterRequest>,
) -> Result<(axum::http::StatusCode, axum::Json<AuthUserRegisterResponse>), axum::http::StatusCode>
{
    let request = tonic::Request::new(payload);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(tonic::Code::InvalidArgument),
                axum::Json(AuthUserRegisterResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    match state
        .auth_use_case
        .clone()
        .auth_register(get_request_id_from_header(&headers), request)
        .await
    {
        Ok(response) => Ok((axum::http::StatusCode::ACCEPTED.into(), axum::Json(response))),
        Err(err) => Ok((
            util::convert_status::tonic_to_http_status(err.code()),
            axum::Json(AuthUserRegisterResponse {
                message: err.message().to_string(),
                status: "error".to_string(),
                data: None,
            }),
        )),
    }
}

#[utoipa::path(
    post,
    path =format!("{}/login", ROUTE_PREFIX.to_string()),
    tag = TAG,
    security(
       ("x-request-id" = [])
    ),
    request_body = AuthUserLoginByEmailAndPasswordRequest,
    responses(
        (status = OK, body = Response, content_type = "application/json" )
    )
)]
#[instrument("auth.http_presenter.auth_user_login_by_email_and_password")]
pub async fn auth_user_login_by_email_and_password(
    State(mut state): State<Presenter>,
    headers: HeaderMap,
    axum::Json(payload): axum::Json<AuthUserLoginByEmailAndPasswordRequest>,
) -> Result<(axum::http::StatusCode, axum::Json<Response>), axum::http::StatusCode> {
    let request = tonic::Request::new(payload);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(tonic::Code::InvalidArgument),
                axum::Json(Response {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }
    match state
        .auth_use_case
        .auth_user_login_by_email_and_password(get_request_id_from_header(&headers), request)
        .await
    {
        Ok(_) => Ok((
            axum::http::StatusCode::ACCEPTED.into(),
            axum::Json(Response {
                status: "success".to_string(),
                message: "auth_user_login_by_email_and_password".to_string(),
                data: None,
            }),
        )),
        Err(err) => Ok((
            util::convert_status::tonic_to_http_status(err.code()),
            axum::Json(Response {
                message: err.message().to_string(),
                status: "error".to_string(),
                data: None,
            }),
        )),
    }
}

#[utoipa::path(
    post,
    path =format!("{}/verify-otp", ROUTE_PREFIX.to_string()),
    tag = TAG,
    request_body = AuthUserVerifyOtpRequest,
    security(
       ("x-request-id" = [])
    ),
    responses(
        (status = OK, body = AuthUserVerifyOtpResponse, content_type = "application/json" )
    )
)]
#[instrument("auth.http_presenter.auth_user_verify_otp")]
pub async fn auth_user_verify_otp(
    State(mut state): State<Presenter>,
    headers: HeaderMap,
    axum::Json(payload): axum::Json<AuthUserVerifyOtpRequest>,
) -> Result<
    (
        axum::http::StatusCode,
        axum::Json<AuthUserVerifyOtpResponse>,
    ),
    axum::http::StatusCode,
> {
    let request = tonic::Request::new(payload);
    match request.validate() {
        Ok(_) => {}
        Err(e) => {
            return Ok((
                util::convert_status::tonic_to_http_status(tonic::Code::InvalidArgument),
                axum::Json(AuthUserVerifyOtpResponse {
                    message: format!("Invalid argument: {}", e.field),
                    status: "error".to_string(),
                    data: None,
                }),
            ));
        }
    }

    match state
        .auth_use_case
        .auth_user_verify_otp(get_request_id_from_header(&headers), request)
        .await
    {
        Ok(response) => Ok((axum::http::StatusCode::OK.into(), axum::Json(response))),
        Err(err) => Ok((
            util::convert_status::tonic_to_http_status(err.code()),
            axum::Json(AuthUserVerifyOtpResponse {
                message: err.message().to_string(),
                status: "error".to_string(),
                data: None,
            }),
        )),
    }
}
