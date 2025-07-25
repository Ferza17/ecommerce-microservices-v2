use crate::config::config::AppConfig;
use crate::infrastructure::database::async_postgres::get_connection;
use crate::interceptor::logger::LoggerLayer;
use crate::interceptor::request_id::RequestIdLayer;
use crate::module::shipping_provider::presenter_http::{
    create_shipping_provider, delete_shipping_provider, get_shipping_provider_by_id,
    list_shipping_providers, update_shipping_provider,
};
use crate::module::shipping_provider::repository_postgres::ShippingProviderPostgresRepositoryImpl;
use crate::module::shipping_provider::usecase::ShippingProviderUseCaseImpl;
use axum::http::StatusCode;
use axum::routing::{delete, get, post, put};
use axum::{Router, response::Json};
use std::sync::Arc;
use tower::ServiceBuilder;

pub struct HttpTransport {
    config: AppConfig,
}

pub struct AppState {
    pub shipping_provider_use_case: ShippingProviderUseCaseImpl,
}

impl HttpTransport {
    pub fn new(config: AppConfig) -> Self {
        HttpTransport { config }
    }

    pub async fn serve(&self) -> Result<(), Box<dyn std::error::Error>> {
        let addr = format!(
            "{}:{}",
            self.config.shipping_service_service_http_host,
            self.config.shipping_service_service_http_port
        )
        .to_string();

        // Infrastructure Layer
        let postgres_pool = get_connection(&self.config.clone()).await;

        // Repository Layer
        let shipping_provider_repository =
            ShippingProviderPostgresRepositoryImpl::new(postgres_pool);

        let app_state = Arc::new(AppState {
            shipping_provider_use_case: ShippingProviderUseCaseImpl::new(
                shipping_provider_repository,
            ),
        });

        let middleware_stack = ServiceBuilder::new()
            .layer(LoggerLayer)
            .layer(RequestIdLayer);

        let app = Router::new()
            .route("/v1/shipping/checks", get(health_check_handler))
            // SHIPPING PROVIDERS ROUTES
            .route(
                "/v1/shipping/shipping_providers",
                post(create_shipping_provider),
            )
            .route(
                "/v1/shipping/shipping_providers",
                get(list_shipping_providers),
            )
            .route(
                "/v1/shipping/shipping_providers/{id}",
                put(update_shipping_provider),
            )
            .route(
                "/v1/shipping/shipping_providers/{id}",
                delete(delete_shipping_provider),
            )
            .route(
                "/v1/shipping/shipping_providers/{id}",
                get(get_shipping_provider_by_id),
            )
            .layer(middleware_stack)
            .with_state(app_state);

        let listener = tokio::net::TcpListener::bind(addr.as_str()).await?;
        eprintln!("Starting HTTP server on {}", addr.as_str());

        axum::serve(listener, app).await?;
        Ok(())
    }
}

async fn health_check_handler() -> (StatusCode, Json<serde_json::Value>) {
    eprintln!("Health check requested: Liveness probe.");
    (
        StatusCode::OK,
        Json(serde_json::json!({
            "status": "Ok",
            "service": "shipping-service",
        })),
    )
}
