use crate::model::rpc::user::{
    AuthUserFindUserByTokenRequest, AuthUserFindUserByTokenResponse,
    AuthUserLoginByEmailAndPasswordRequest, AuthUserRegisterRequest, AuthUserRegisterResponse,
    AuthUserVerifyOtpRequest, AuthUserVerifyOtpResponse,
};
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct UseCase {
    auth_transport_grpc: crate::module::auth::transport_grpc::Transport,
}

impl UseCase {
    pub fn new(auth_transport_grpc: crate::module::auth::transport_grpc::Transport) -> Self {
        Self {
            auth_transport_grpc,
        }
    }

    #[instrument("auth.usecase.auth_register")]
    pub async fn auth_register(
        &mut self,
        request_id: String,
        request: tonic::Request<AuthUserRegisterRequest>,
    ) -> Result<AuthUserRegisterResponse, tonic::Status> {
        match self
            .auth_transport_grpc
            .auth_register(request_id, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }

    #[instrument("auth.usecase.auth_user_login_by_email_and_password")]
    pub async fn auth_user_login_by_email_and_password(
        &mut self,
        request_id: String,
        request: tonic::Request<AuthUserLoginByEmailAndPasswordRequest>,
    ) -> Result<(), tonic::Status> {
        match self
            .auth_transport_grpc
            .auth_user_login_by_email_and_password(request_id, request)
            .await
        {
            Err(e) => Err(e.into()),
            Ok(_) => Ok(()),
        }
    }

    #[instrument("auth.usecase.auth_user_verify_otp")]
    pub async fn auth_user_verify_otp(
        &mut self,
        request_id: String,
        request: tonic::Request<AuthUserVerifyOtpRequest>,
    ) -> Result<AuthUserVerifyOtpResponse, tonic::Status> {
        match self
            .auth_transport_grpc
            .auth_user_verify_otp(request_id, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }

    #[instrument("auth.usecase.auth_user_find_user_by_token")]
    pub async fn auth_user_find_user_by_token(
        &mut self,
        request_id: String,
        token: String,
        request: tonic::Request<AuthUserFindUserByTokenRequest>,
    ) -> Result<AuthUserFindUserByTokenResponse, tonic::Status> {
        match self
            .auth_transport_grpc
            .auth_user_find_user_by_token(request_id, token, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }
}
