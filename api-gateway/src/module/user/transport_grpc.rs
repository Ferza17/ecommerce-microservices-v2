use crate::config::config::AppConfig;
use crate::model::rpc::user::auth_service_client::AuthServiceClient;
use crate::model::rpc::user::user_service_client::UserServiceClient;
use crate::model::rpc::user::{
    AuthServiceVerifyIsExcludedRequest, AuthServiceVerifyIsExcludedResponse,
    AuthUserLoginByEmailAndPasswordRequest, AuthUserRegisterRequest, AuthUserRegisterResponse,
    AuthUserVerifyAccessControlRequest, AuthUserVerifyOtpRequest, AuthUserVerifyOtpResponse,
    FindUserByIdRequest, FindUserByIdResponse,
};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use opentelemetry::trace::FutureExt;
use tonic::{Response, Status};
use tracing::{Level, Span, event, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Debug, Clone)]
pub struct UserTransportGrpc {
    pub(crate) auth_service_client: AuthServiceClient<tonic::transport::Channel>,
    user_service_client: UserServiceClient<tonic::transport::Channel>,
}

impl UserTransportGrpc {
    #[instrument]
    pub async fn new(config: AppConfig) -> Result<Self, anyhow::Error> {
        let addr = format!(
            "http://{}:{}",
            config.user_service_service_rpc_host, config.user_service_service_rpc_port
        );
        let channel = tonic::transport::Channel::from_shared(addr.to_string())
            .expect("Failed to connect to user service")
            .connect()
            .await
            .map_err(|e| panic!("user service not connected : {}", e))
            .unwrap();
        Ok(Self {
            auth_service_client: AuthServiceClient::new(channel.clone()),
            user_service_client: UserServiceClient::new(channel),
        })
    }

    #[instrument]
    pub async fn auth_service_verify_is_excluded(
        &mut self,
        request_id: String,
        mut request: tonic::Request<AuthServiceVerifyIsExcludedRequest>,
    ) -> Result<AuthServiceVerifyIsExcludedResponse, tonic::Status> {
        // REQUEST ID TO HEADER
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());

        match self
            .auth_service_client
            .auth_service_verify_is_excluded(request)
            .with_context(Span::current().context())
            .await
        {
            Ok(response) => {
                event!(
                    Level::INFO,
                    request_id = request_id,
                    data=?response
                );
                Ok(response.into_inner())
            }
            Err(err) => {
                event!(
                    Level::ERROR,
                    request_id = request_id,
                    error = %err,
                    "Failed to get auth_service_verify_is_excluded"
                );
                Err(err)
            }
        }
    }

    pub async fn auth_user_verify_access_control(
        &mut self,
        request_id: String,
        mut request: tonic::Request<AuthUserVerifyAccessControlRequest>,
    ) {
        // TOKEN TO HEADER
        let token = request.get_ref().token.clone();
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );
        // REQUEST ID TO HEADER
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
    }

    #[instrument]
    pub async fn auth_register(
        &mut self,
        request_id: String,
        mut request: tonic::Request<AuthUserRegisterRequest>,
    ) -> Result<AuthUserRegisterResponse, tonic::Status> {
        // REQUEST ID TO HEADER
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());

        match self
            .auth_service_client
            .auth_user_register(request)
            .with_context(Span::current().context())
            .await
        {
            Ok(response) => {
                event!(
                    Level::INFO,
                    request_id = request_id,
                    data=?response
                );
                Ok(response.into_inner())
            }
            Err(err) => {
                event!(
                    Level::ERROR,
                    request_id = request_id,
                    error = %err,
                    "Failed to get auth_user_register"
                );
                Err(err)
            }
        }
    }

    #[instrument]
    pub async fn auth_user_login_by_email_and_password(
        &mut self,
        request_id: String,
        mut request: tonic::Request<AuthUserLoginByEmailAndPasswordRequest>,
    ) -> Result<(), tonic::Status> {
        // REQUEST ID TO HEADER
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());

        match self
            .auth_service_client
            .auth_user_login_by_email_and_password(request)
            .with_context(Span::current().context())
            .await
        {
            Ok(response) => {
                event!(
                    Level::INFO,
                    request_id = request_id,
                    data=?response
                );
                Ok(response.into_inner())
            }
            Err(err) => {
                event!(
                    Level::ERROR,
                    request_id = request_id,
                    error = %err,
                    "Failed to get auth_user_login_by_email_and_password"
                );
                Err(err)
            }
        }
    }

    #[instrument]
    pub async fn auth_user_verify_otp(
        &mut self,
        request_id: String,
        mut request: tonic::Request<AuthUserVerifyOtpRequest>,
    ) -> Result<AuthUserVerifyOtpResponse, tonic::Status> {
        // REQUEST ID TO HEADER
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());

        match self
            .auth_service_client
            .auth_user_verify_otp(request)
            .with_context(Span::current().context())
            .await
        {
            Ok(response) => {
                event!(
                    Level::INFO,
                    request_id = request_id,
                    data=?response
                );
                Ok(response.into_inner())
            }
            Err(err) => {
                event!(
                    Level::ERROR,
                    request_id = request_id,
                    error = %err,
                    "Failed to get auth_user_verify_otp"
                );
                Err(err)
            }
        }
    }

    #[instrument]
    pub async fn find_user_by_id(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<FindUserByIdRequest>,
    ) -> Result<FindUserByIdResponse, tonic::Status> {
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );

        match self
            .user_service_client
            .find_user_by_id(request)
            .with_context(Span::current().context())
            .await
        {
            Ok(response) => {
                event!(
                    Level::INFO,
                    request_id = request_id,
                    data=?response
                );
                Ok(response.into_inner())
            }
            Err(err) => {
                event!(
                    Level::ERROR,
                    request_id = request_id,
                    error = %err,
                    "Failed to get find_user_by_id"
                );
                Err(err.into())
            }
        }
    }
}
