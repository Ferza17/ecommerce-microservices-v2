use crate::interceptor::auth::AuthLayer;
use crate::model::rpc::{
    response::Response,
    user::{
        AuthUserLoginByEmailAndPasswordRequest, AuthUserRegisterRequest, AuthUserRegisterResponse,
        AuthUserVerifyOtpRequest, AuthUserVerifyOtpResponse,
    },
};
use crate::module::user::usecase::UserUseCase;
use crate::package::context::request_id::get_request_id_from_header;
use crate::util;
use axum::{extract::State, http::HeaderMap, routing::post};
use prost_validate::NoopValidator;
use tower::ServiceBuilder;
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct UserPresenterHttp {
    user_use_case: UserUseCase,
}

impl UserPresenterHttp {
    pub fn new(user_use_case: UserUseCase) -> Self {
        Self { user_use_case }
    }

    #[instrument]
    pub fn auth_router(&self) -> axum::Router {
        axum::Router::new()
            .route("/register", post(auth_register))
            .route("/login", post(auth_user_login_by_email_and_password))
            .route("/verify-otp", post(auth_user_verify_otp))
            .with_state(self.clone())
    }

    #[instrument]
    pub fn user_router(&self) -> axum::Router {
        axum::Router::new()
            .layer(ServiceBuilder::new().layer(AuthLayer::new(self.user_use_case.clone())))
            .with_state(self.clone())
    }
}

#[utoipa::path(
    post,
    path = "/api/v1/auth/register",
    tag = "Auth",
    request_body = AuthUserRegisterRequest,
    responses(
        (status = OK, body = AuthUserRegisterResponse, content_type = "application/json" )
    )
)]
#[instrument("UserPresenterHttp.auth_register")]
pub async fn auth_register(
    State(state): State<UserPresenterHttp>,
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
        .user_use_case
        .clone()
        .auth_register(get_request_id_from_header(&headers), request)
        .await
    {
        Ok(response) => Ok((axum::http::StatusCode::OK.into(), axum::Json(response))),
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
    path = "/api/v1/auth/login",
    tag = "Auth",
    request_body = AuthUserLoginByEmailAndPasswordRequest,
    responses(
        (status = OK, body = Response, content_type = "application/json" )
    )
)]
#[instrument("UserPresenterHttp.auth_user_login_by_email_and_password")]
pub async fn auth_user_login_by_email_and_password(
    State(mut state): State<UserPresenterHttp>,
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
        .user_use_case
        .auth_user_login_by_email_and_password(get_request_id_from_header(&headers), request)
        .await
    {
        Ok(_) => Ok((
            axum::http::StatusCode::OK.into(),
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
    path = "/api/v1/auth/verify-otp",
    tag = "Auth",
    request_body = AuthUserVerifyOtpRequest,
    responses(
        (status = OK, body = AuthUserVerifyOtpResponse, content_type = "application/json" )
    )
)]
#[instrument("UserPresenterHttp.auth_user_verify_otp")]
pub async fn auth_user_verify_otp(
    State(mut state): State<UserPresenterHttp>,
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
        .user_use_case
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
