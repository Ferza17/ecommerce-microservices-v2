use opentelemetry::trace::FutureExt;
use tracing::{event, instrument, Level, Span};
use tracing_opentelemetry::OpenTelemetrySpanExt;
use crate::config::config::AppConfig;
use crate::model::rpc::shipping::{GetShippingProviderByIdRequest, GetShippingProviderByIdResponse, ListShippingProvidersRequest, ListShippingProvidersResponse};
use crate::model::rpc::shipping::shipping_provider_service_client::ShippingProviderServiceClient;
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::inject_trace_context;

#[derive(Debug, Clone)]
pub struct Transport {
    client: ShippingProviderServiceClient<tonic::transport::Channel>,
}

impl Transport {
    pub async fn new(config: AppConfig) -> Result<Self, anyhow::Error> {
        let channel = tonic::transport::Channel::from_shared(
            format!(
                "http://{}:{}",
                config.shipping_service_service_rpc_host, config.shipping_service_service_rpc_port
            )
            .to_string(),
        )
        .expect("Failed to connect to shipping service")
        .connect()
        .await
        .map_err(|e| panic!("shipping service not connected : {}", e))
        .unwrap();
        Ok(Self {
            client: ShippingProviderServiceClient::new(channel),
        })
    }

    #[instrument("shipping_provider.transport_grpc.list_shipping_providers")]
    pub async fn list_shipping_providers(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<ListShippingProvidersRequest>,
    ) -> Result<ListShippingProvidersResponse, tonic::Status> {
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );

        match self
            .client
            .list_shipping_providers(inject_trace_context(request, Span::current().context()))
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
                    "Failed to get list_shipping_providers"
                );
                Err(err.into())
            }
        }
    }

    #[instrument("shipping_provider.transport_grpc.get_shipping_provider_by_id")]
    pub async fn get_shipping_provider_by_id(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<GetShippingProviderByIdRequest>,
    ) -> Result<GetShippingProviderByIdResponse, tonic::Status> {
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );
        match self
            .client
            .get_shipping_provider_by_id(inject_trace_context(request, Span::current().context()))
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
                    "Failed to get get_shipping_provider_by_id"
                );
                Err(err.into())
            }
        }
    }
}
