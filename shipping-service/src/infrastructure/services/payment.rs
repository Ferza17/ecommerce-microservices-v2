use crate::config::config::AppConfig;
use crate::model::rpc::payment::payment_service_client::PaymentServiceClient;
use crate::model::rpc::payment::{FindPaymentByIdRequest, FindPaymentByIdResponse, Payment};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::MetadataInjector;
use opentelemetry::global;
use tonic::transport::Channel;
use tonic::{Request, Status};
use tracing::{Level, Span, event, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Clone, Debug)]
pub struct PaymentServiceGrpcClient {
    payment_service_client: PaymentServiceClient<Channel>,
}

impl PaymentServiceGrpcClient {
    pub async fn new(config: AppConfig) -> Self {
        let channel = Channel::from_shared(
            format!(
                "http://{}:{}",
                config.payment_service_service_rpc_host, config.payment_service_service_rpc_port
            )
            .to_string(),
        )
        .expect("Failed to connect to user service")
        .connect()
        .await
        .map_err(|e| panic!("payment service not connected : {}", e))
        .unwrap();

        Self {
            payment_service_client: PaymentServiceClient::new(channel),
        }
    }

    #[instrument]
    pub async fn find_payment_by_id(
        &mut self,
        request_id: String,
        token: String,
        request: FindPaymentByIdRequest,
    ) -> Result<FindPaymentByIdResponse, Status> {
        let mut req = Request::new(request.clone());
        // TOKEN TO HEADER
        req.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );
        // REQUEST ID TO HEADER
        req.metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        // SEND TRACER PARENT
        global::get_text_map_propagator(|propagator| {
            propagator.inject_context(
                &Span::current().context(),
                &mut MetadataInjector(req.metadata_mut()),
            )
        });
        match self.payment_service_client.find_payment_by_id(req).await {
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
}
