use crate::config::config::AppConfig;
use crate::model::rpc::payment::{
    CreatePaymentRequest, CreatePaymentResponse, FindPaymentProviderByIdRequest,
    FindPaymentProviderByIdResponse, FindPaymentProvidersRequest, FindPaymentProvidersResponse,
    payment_provider_service_client::PaymentProviderServiceClient,
    payment_service_client::PaymentServiceClient,
};
use crate::package::context::{auth::AUTHORIZATION_HEADER, request_id::X_REQUEST_ID_HEADER};
use crate::util::metadata::inject_trace_context;
use tracing::{Level, Span, event, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Debug, Clone)]
pub struct Transport {
    payment_service_client: PaymentServiceClient<tonic::transport::Channel>,
    payment_provider_service_client: PaymentProviderServiceClient<tonic::transport::Channel>,
}

impl Transport {
    pub async fn new(config: AppConfig) -> Result<Self, anyhow::Error> {
        let channel = tonic::transport::Channel::from_shared(
            format!(
                "http://{}:{}",
                config.service_payment.rpc_host, config.service_payment.rpc_port
            )
            .to_string(),
        )
        .expect("Failed to connect to payment service")
        .connect()
        .await
        .map_err(|e| panic!("payment service not connected : {}", e))
        .unwrap();
        Ok(Self {
            payment_service_client: PaymentServiceClient::new(channel.clone()),
            payment_provider_service_client: PaymentProviderServiceClient::new(channel),
        })
    }

    #[instrument("payment.transport_grpc.create_payment")]
    pub async fn create_payment(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<CreatePaymentRequest>,
    ) -> Result<CreatePaymentResponse, tonic::Status> {
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );
        match self
            .payment_service_client
            .create_payment(inject_trace_context(request, Span::current().context()))
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
                    "Failed to create_payment"
                );
                Err(err.into())
            }
        }
    }
}
