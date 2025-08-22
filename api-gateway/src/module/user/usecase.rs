use crate::model::rpc::user::{
    AuthServiceVerifyIsExcludedRequest, AuthServiceVerifyIsExcludedResponse,
    AuthUserLoginByEmailAndPasswordRequest, AuthUserRegisterRequest, AuthUserRegisterResponse,
    AuthUserVerifyOtpRequest, AuthUserVerifyOtpResponse,
};
use crate::module::user::{
    transport_grpc::UserTransportGrpc, transport_rabbitmq::UserTransportRabbitMQ,
};
use futures::future::{Either, ready};
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct UserUseCase {
    user_service_grpc: UserTransportGrpc,
    user_service_rabbitmq: UserTransportRabbitMQ,
}

impl UserUseCase {
    pub fn new(
        user_service_grpc: UserTransportGrpc,
        user_service_rabbitmq: UserTransportRabbitMQ,
    ) -> Self {
        Self {
            user_service_grpc,
            user_service_rabbitmq,
        }
    }

    #[instrument("UserUseCase.auth_register")]
    pub async fn auth_register(
        &mut self,
        request_id: String,
        request: tonic::Request<AuthUserRegisterRequest>,
    ) -> Result<AuthUserRegisterResponse, tonic::Status> {
        match self
            .user_service_grpc
            .auth_register(request_id, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }

    #[instrument("UserUseCase.auth_user_login_by_email_and_password")]
    pub async fn auth_user_login_by_email_and_password(
        &mut self,
        request_id: String,
        request: tonic::Request<AuthUserLoginByEmailAndPasswordRequest>,
    ) -> Result<(), tonic::Status> {
        match self
            .user_service_grpc
            .auth_user_login_by_email_and_password(request_id, request)
            .await
        {
            Err(e) => Err(e.into()),
            Ok(_) => Ok(()),
        }
    }

    #[instrument("UserUseCase.auth_user_verify_otp")]
    pub async fn auth_user_verify_otp(
        &mut self,
        request_id: String,
        request: tonic::Request<AuthUserVerifyOtpRequest>,
    ) -> Result<AuthUserVerifyOtpResponse, tonic::Status> {
        match self
            .user_service_grpc
            .auth_user_verify_otp(request_id, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => Ok(response?),
        }
    }

    #[instrument("UserUseCase.auth_service_verify_is_excluded")]
    pub async fn auth_service_verify_is_excluded(
        &mut self,
        request_id: String,
        request: tonic::Request<AuthServiceVerifyIsExcludedRequest>,
    ) -> Result<AuthServiceVerifyIsExcludedResponse, tonic::Status> {
        match self
            .user_service_grpc
            .auth_service_verify_is_excluded(request_id, request)
            .await
        {
            Err(e) => Err(e.into()),
            response => {
                let resp = response?;
                let Some(data) = resp.data else {
                    return Err(tonic::Status::unauthenticated("no data in response"));
                };

                if data.is_excluded {
                    return Err(tonic::Status::unauthenticated("user is excluded"));
                }

                Ok(resp)
            }
        }
    }
}
