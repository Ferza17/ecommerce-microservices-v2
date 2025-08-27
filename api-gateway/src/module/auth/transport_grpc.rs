use opentelemetry::trace::FutureExt;
use crate::config::config::AppConfig;
use crate::model::rpc::user::{AuthUserLoginByEmailAndPasswordRequest, AuthUserRegisterRequest, AuthUserRegisterResponse, AuthUserVerifyOtpRequest, AuthUserVerifyOtpResponse, auth_service_client::AuthServiceClient, AuthUserFindUserByTokenRequest, AuthUserFindUserByTokenResponse};
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::inject_trace_context;
use tracing::{Level, Span, event, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;
use crate::package::context::auth::AUTHORIZATION_HEADER;

#[derive(Debug, Clone)]
pub struct Transport {
    client: AuthServiceClient<tonic::transport::Channel>,
}

impl Transport {
    pub async fn new(config: AppConfig) -> Result<Self, anyhow::Error> {
        let addr = format!(
            "http://{}:{}",
            config.service_user.rpc_host, config.service_user.rpc_port
        );
        let channel = tonic::transport::Channel::from_shared(addr.to_string())
            .expect("Failed to connect to user service")
            .connect()
            .await
            .map_err(|e| panic!("user service not connected : {}", e))
            .unwrap();
        Ok(Self {
            client: AuthServiceClient::new(channel),
        })
    }

    #[instrument("auth.transport_grpc.auth_register")]
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
            .client
            .auth_user_register(inject_trace_context(request, Span::current().context()))
            .with_current_context()
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

    #[instrument("auth.transport_grpc.auth_user_login_by_email_and_password")]
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
            .client
            .auth_user_login_by_email_and_password(inject_trace_context(
                request,
                Span::current().context(),
            ))
            .with_current_context()
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

    #[instrument("auth.transport_grpc.auth_user_verify_otp")]
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
            .client
            .auth_user_verify_otp(inject_trace_context(request, Span::current().context()))
            .with_current_context()
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

    #[instrument("auth.transport_grpc.auth_user_find_user_by_token")]
    pub async fn auth_user_find_user_by_token(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<AuthUserFindUserByTokenRequest>,
    ) -> Result<AuthUserFindUserByTokenResponse, tonic::Status> {
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );

        match self
            .client
            .auth_user_find_user_by_token(inject_trace_context(request, Span::current().context()))
            .with_current_context()
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
                    "Failed to get auth_user_find_user_by_token"
                );
                Err(err.into())
            }
        }
    }
}
