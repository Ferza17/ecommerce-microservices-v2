use crate::config::config::AppConfig;
use crate::model::rpc::product::{
    FindProductByIdRequest, FindProductsWithPaginationRequest, FindProductsWithPaginationResponse,
    Product, product_service_client::ProductServiceClient,
};
use crate::package::context::{auth::AUTHORIZATION_HEADER, request_id::X_REQUEST_ID_HEADER};
use crate::util::metadata::inject_trace_context;
use opentelemetry::trace::FutureExt;
use tracing::{Level, Span, error, event, info, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Debug, Clone)]
pub struct Transport {
    product_service_client: ProductServiceClient<tonic::transport::Channel>,
}

impl Transport {
    pub async fn new(config: AppConfig) -> Result<Self, anyhow::Error> {
        let addr = format!(
            "http://{}:{}",
            config.service_product.rpc_host, config.service_product.rpc_port
        );
        let channel = tonic::transport::Channel::from_shared(addr.to_string())
            .expect("Failed to connect to user service")
            .connect()
            .await
            .map_err(|e| panic!("product service not connected : {}", e))
            .unwrap();

        Ok(Self {
            product_service_client: ProductServiceClient::new(channel),
        })
    }

    #[instrument("product.transport_grpc.find_products_with_pagination")]
    pub async fn find_products_with_pagination(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<FindProductsWithPaginationRequest>,
    ) -> Result<FindProductsWithPaginationResponse, tonic::Status> {
        // REQUEST ID TO HEADER
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());

        // TOKEN TO HEADER
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );
        info!(request_id = request_id, request = ?request, "Request to find products with pagination");
        match self
            .product_service_client
            .find_products_with_pagination(inject_trace_context(request, Span::current().context()))
            .with_current_context()
            .await
        {
            Ok(response) => {
                info!(request_id = request_id, data=?response,"Response to find products with pagination");
                Ok(response.into_inner())
            }
            Err(err) => {
                error!(
                    request_id = request_id,
                    error = %err,
                    "Failed to get find_products_with_pagination"
                );
                Err(err)
            }
        }
    }

    #[instrument("product.transport_grpc.find_product_by_id")]
    pub async fn find_product_by_id(
        &mut self,
        request_id: String,
        token: String,
        mut request: tonic::Request<FindProductByIdRequest>,
    ) -> Result<Product, tonic::Status> {
        request
            .metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        request.metadata_mut().insert(
            AUTHORIZATION_HEADER,
            format!("Bearer {}", token).parse().unwrap(),
        );

        info!(request_id = request_id, request = ?request, "Find products with pagination");
        match self
            .product_service_client
            .find_product_by_id(inject_trace_context(request, Span::current().context()))
            .with_current_context()
            .await
        {
            Ok(response) => {
                info!(request_id = request_id, data=?response,"Response to find_product_by_id");
                Ok(response.into_inner())
            }
            Err(err) => {
                error!(
                    request_id = request_id,
                    error = %err,
                    "Failed to find_product_by_id"
                );
                Err(err)
            }
        }
    }
}
