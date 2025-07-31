use crate::config::config::AppConfig;
use crate::model::rpc::user::auth_service_client::AuthServiceClient;
use crate::model::rpc::user::user_service_client::UserServiceClient;
use crate::model::rpc::user::{
    AuthServiceVerifyIsExcludedRequest, AuthServiceVerifyIsExcludedResponse,
    AuthUserFindUserByTokenRequest, AuthUserFindUserByTokenResponse,
    AuthUserVerifyAccessControlRequest, AuthUserVerifyAccessControlResponse, FindUserByIdRequest,
    FindUserByIdResponse,
};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::MetadataInjector;
use opentelemetry::global;
use tonic::transport::Channel;
use tonic::{Request, Status};
use tracing::{Level, Span, event, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Clone, Debug)]
pub struct UserServiceGrpcClient {
    auth_service_client: AuthServiceClient<Channel>,
    user_service_client: UserServiceClient<Channel>,
}

impl UserServiceGrpcClient {
    #[instrument]
    pub async fn new(config: AppConfig) -> Self {
        let addr = format!(
            "http://{}:{}",
            config.user_service_service_rpc_host, config.user_service_service_rpc_port
        );
        let channel = Channel::from_shared(addr.to_string())
            .expect("Failed to connect to user service")
            .connect()
            .await
            .map_err(|e| panic!("user service not connected : {}", e))
            .unwrap();
        Self {
            auth_service_client: AuthServiceClient::new(channel.clone()),
            user_service_client: UserServiceClient::new(channel),
        }
    }

    #[instrument]
    pub async fn auth_user_verify_access_control(
        &mut self,
        request_id: String,
        mut request: Request<AuthUserVerifyAccessControlRequest>,
    ) -> Result<AuthUserVerifyAccessControlResponse, Status> {
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
        // SEND TRACER PARENT
        global::get_text_map_propagator(|propagator| {
            propagator.inject_context(
                &Span::current().context(),
                &mut MetadataInjector(request.metadata_mut()),
            )
        });

        match self
            .auth_service_client
            .auth_user_verify_access_control(request)
            .await
        {
            Ok(response) => {
                event!(Level::INFO,
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
                    "Failed to get auth_user_verify_access_control"
                );
                Err(err)
            }
        }
    }

    #[instrument]
    pub async fn auth_service_verify_is_excluded(
        &mut self,
        request_id: String,
        mut request: Request<AuthServiceVerifyIsExcludedRequest>,
    ) -> Result<AuthServiceVerifyIsExcludedResponse, Status> {
        // REQUEST ID TO HEADER
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());

        match self
            .auth_service_client
            .auth_service_verify_is_excluded(request)
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

    #[instrument]
    pub async fn auth_user_find_user_by_token(
        &mut self,
        request_id: String,
        mut request: Request<AuthUserFindUserByTokenRequest>,
    ) -> Result<AuthUserFindUserByTokenResponse, Status> {
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

        // SEND TRACER PARENT
        global::get_text_map_propagator(|propagator| {
            propagator.inject_context(
                &Span::current().context(),
                &mut MetadataInjector(request.metadata_mut()),
            )
        });

        match self
            .auth_service_client
            .auth_user_find_user_by_token(request)
            .await
        {
            Ok(response) => {
                event!(Level::INFO,
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
                Err(err)
            }
        }
    }

    #[instrument]
    pub async fn find_user_by_id(
        &mut self,
        request_id: String,
        token: String,
        mut request: Request<FindUserByIdRequest>,
    ) -> Result<FindUserByIdResponse, Status> {
        // TOKEN TO HEADER
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );
        // REQUEST ID TO HEADER
        request.metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        // SEND TRACER PARENT
        global::get_text_map_propagator(|propagator| {
            propagator.inject_context(
                &Span::current().context(),
                &mut MetadataInjector(request.metadata_mut()),
            )
        });
        match self.user_service_client.clone().find_user_by_id(request).await {
            Ok(response) => {
                event!(Level::INFO,
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
                Err(err)
            }
        }
    }
}
