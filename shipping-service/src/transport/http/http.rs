use crate::config::config::AppConfig;
use crate::infrastructure::database::postgres::create_postgres_pool;
use crate::interceptor::request_id::http_interceptor::x_request_id_middleware;
use crate::module::shipping_provider::presenter_http::{
    create_shipping_provider, delete_shipping_provider, get_shipping_provider_by_id,
    list_shipping_providers, update_shipping_provider,
};
use crate::module::shipping_provider::repository_postgres::ShippingProviderPostgresRepository;
use crate::module::shipping_provider::usecase::ShippingProviderUseCase;
use axum::http::StatusCode;
use axum::routing::{delete, get, post, put};
use axum::{
    Router,
    middleware::{self, Next},
    response::Json,
};
use axum_tracing_opentelemetry::middleware::OtelAxumLayer;
use std::sync::Arc;

pub struct HttpTransport {
    config: AppConfig,
}

pub struct AppState {
    pub shipping_provider_use_case: ShippingProviderUseCase,
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
        let postgres_pool = create_postgres_pool(&self.config.clone())
            .await
            .expect("Failed to create postgres pool");

        // Repository Layer
        let shipping_provider_repository = ShippingProviderPostgresRepository::new(postgres_pool);

        let app_state = Arc::new(AppState {
            shipping_provider_use_case: ShippingProviderUseCase::new(shipping_provider_repository),
        });

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
            .layer(middleware::from_fn(x_request_id_middleware))
            .layer(OtelAxumLayer::default())
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
