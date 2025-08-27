use crate::infrastructure::opa::opa::OpaInput;
use crate::model::rpc::user::{
    AuthUserFindUserByTokenRequest, AuthUserFindUserByTokenResponse,
    AuthUserLoginByEmailAndPasswordRequest, AuthUserRegisterRequest, AuthUserRegisterResponse,
    AuthUserVerifyOtpRequest, AuthUserVerifyOtpResponse, EnumRole, FindUserByEmailRequest, User,
};
use tracing::instrument;

#[derive(Debug, Clone)]
pub struct UseCase {
    auth_transport_grpc: crate::module::auth::transport_grpc::Transport,
    user_transport_grpc: crate::module::user::transport_grpc::Transport,
    rabbitmq_infrastructure: crate::module::user::transport_rabbitmq::Transport,
    opa_infrastructure: crate::infrastructure::opa::opa::OPA,
}

impl UseCase {
    pub fn new(
        auth_transport_grpc: crate::module::auth::transport_grpc::Transport,
        user_transport_grpc: crate::module::user::transport_grpc::Transport,
        rabbitmq_infrastructure: crate::module::user::transport_rabbitmq::Transport,
        opa_infrastructure: crate::infrastructure::opa::opa::OPA,
    ) -> Self {
        Self {
            auth_transport_grpc,
            user_transport_grpc,
            rabbitmq_infrastructure,
            opa_infrastructure,
        }
    }

    #[instrument("auth.usecase.auth_register")]
    pub async fn auth_register(
        &mut self,
        request_id: String,
        request: tonic::Request<AuthUserRegisterRequest>,
    ) -> Result<AuthUserRegisterResponse, tonic::Status> {
        // VALIDATE ROLE
        match EnumRole::from_i32(request.get_ref().role) {
            None => return Err(tonic::Status::invalid_argument("Role is not valid")),
            Some(_) => {}
        }

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
            .rabbitmq_infrastructure
            .publish_user_login(request_id, request)
            .await
        {
            Ok(_) => Ok(()),
            Err(e) => Err(e.into()),
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

    pub async fn auth_validate_access(
        &mut self,
        method: String,
        path: String,
        request: User,
    ) -> Result<(), tonic::Status> {
        match self
            .opa_infrastructure
            .validate_http_access(OpaInput {
                method: method.to_uppercase(),
                path: path.to_lowercase(),
                user_id: request.id,
                user_role: EnumRole::from_i32(request.role.unwrap().role)
                    .unwrap()
                    .as_str_name()
                    .to_string(),
            })
            .await
        {
            Ok(is_valid) => {
                if !is_valid {
                    Err(tonic::Status::permission_denied("Access denied"))
                } else {
                    Ok(())
                }
            }
            Err(e) => Err(tonic::Status::permission_denied(e.to_string()))?,
        }
    }
}
