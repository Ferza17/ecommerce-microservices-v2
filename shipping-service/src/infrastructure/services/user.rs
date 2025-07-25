use crate::config::config::AppConfig;
use crate::model::rpc::user::auth_service_client::AuthServiceClient;
use crate::model::rpc::user::{
    AuthServiceVerifyIsExcludedRequest, AuthServiceVerifyIsExcludedResponse,
    AuthUserVerifyAccessControlRequest, AuthUserVerifyAccessControlResponse,
};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::{MetadataInjector, grpc_inject_trace_context};
use opentelemetry::global;
use opentelemetry::propagation::Injector;
use tonic::metadata::{MetadataKey, MetadataValue};
use tonic::transport::Channel;
use tonic::{Request, Status};
use tracing::{Level, Span, event, info, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Clone, Debug)]
pub struct UserServiceGrpcClient {
    config: AppConfig,
    auth_service_client: AuthServiceClient<Channel>,
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
            .await;

        let channel = match channel {
            Ok(chan) => {
                eprintln!(
                    "connected to {} addr {}:{} ",
                    config.user_service_service_name,
                    config.user_service_service_rpc_host,
                    config.user_service_service_rpc_port
                );
                chan
            }
            Err(e) => panic!("Failed to connect to user service: {}", e),
        };

        Self {
            config,
            auth_service_client: AuthServiceClient::new(channel),
        }
    }

    #[instrument]
    pub async fn auth_user_verify_access_control(
        &mut self,
        request_id: String,
        request: AuthUserVerifyAccessControlRequest,
    ) -> Result<AuthUserVerifyAccessControlResponse, Status> {
        let mut req = Request::new(request.clone());
        // TOKEN TO HEADER
        req.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", request.token).parse().unwrap(),
        );

        // REQUEST ID TO HEADER
        req.metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());

        // SEND TRACER PARENT
        req = grpc_inject_trace_context(req);


        req.metadata_mut().iter_mut().for_each(|m| {
            eprintln!("{:?}", m);
        });

        let response = self
            .auth_service_client
            .auth_user_verify_access_control(req)
            .await;

        match response {
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

    pub async fn auth_service_verify_is_excluded(
        &mut self,
        request_id: String,
        request: AuthServiceVerifyIsExcludedRequest,
    ) -> Result<AuthServiceVerifyIsExcludedResponse, Status> {
        let response = self
            .auth_service_client
            .auth_service_verify_is_excluded(request)
            .await;

        match response {
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
}
