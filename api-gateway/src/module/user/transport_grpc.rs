use crate::config::config::AppConfig;
use crate::model::rpc::user::{
    FindUserByIdRequest, FindUserByIdResponse, auth_service_client::AuthServiceClient,
    user_service_client::UserServiceClient,
};
use crate::package::context::{auth::AUTHORIZATION_HEADER, request_id::X_REQUEST_ID_HEADER};
use crate::util::metadata::inject_trace_context;
use opentelemetry::trace::FutureExt;
use tracing::{Level, Span, event, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Debug, Clone)]
pub struct Transport {
    pub(crate) auth_service_client: AuthServiceClient<tonic::transport::Channel>,
    user_service_client: UserServiceClient<tonic::transport::Channel>,
}

impl Transport {
    #[instrument]
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
            auth_service_client: AuthServiceClient::new(channel.clone()),
            user_service_client: UserServiceClient::new(channel),
        })
    }

    #[instrument("user.transport_grpc.find_user_by_id")]
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
            .find_user_by_id(inject_trace_context(request, Span::current().context()))
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
                    "Failed to get find_user_by_id"
                );
                Err(err.into())
            }
        }
    }
}
