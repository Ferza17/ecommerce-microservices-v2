use crate::config::config::AppConfig;
use crate::model::rpc::payment::{
    FindPaymentProviderByIdRequest, FindPaymentProviderByIdResponse, FindPaymentProvidersRequest,
    FindPaymentProvidersResponse, payment_provider_service_client::PaymentProviderServiceClient,
};
use crate::package::context::{auth::AUTHORIZATION_HEADER, request_id::X_REQUEST_ID_HEADER};
use crate::util::metadata::inject_trace_context;
use opentelemetry::trace::FutureExt;
use tracing::{Level, Span, event, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Debug, Clone)]
pub struct Transport {
    client: PaymentProviderServiceClient<tonic::transport::Channel>,
}

impl Transport {
    pub async fn new(config: AppConfig) -> Result<Self, anyhow::Error> {
        let addr = format!(
            "http://{}:{}",
            config.payment_service_service_rpc_host, config.payment_service_service_rpc_port
        );
        let channel = tonic::transport::Channel::from_shared(addr.to_string())
            .expect("Failed to connect to payment service")
            .connect()
            .await
            .map_err(|e| panic!("payment service not connected : {}", e))
            .unwrap();
        Ok(Self {
            client: PaymentProviderServiceClient::new(channel),
        })
    }

    #[instrument("payment_providers.transport_grpc.find_payment_providers")]
    pub async fn find_payment_providers(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<FindPaymentProvidersRequest>,
    ) -> Result<FindPaymentProvidersResponse, tonic::Status> {
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );

        match self
            .client
            .find_payment_providers(inject_trace_context(request, Span::current().context()))
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
                    "Failed to get find_payment_providers"
                );
                Err(err.into())
            }
        }
    }

    #[instrument("payment_providers.transport_grpc.find_payment_provider_by_id")]
    pub async fn find_payment_provider_by_id(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<FindPaymentProviderByIdRequest>,
    ) -> Result<FindPaymentProviderByIdResponse, tonic::Status> {
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );

        match self
            .client
            .find_payment_provider_by_id(inject_trace_context(request, Span::current().context()))
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
                    "Failed to find_payment_provider_by_id"
                );
                Err(err.into())
            }
        }
    }
}
