use crate::config::config::AppConfig;
use crate::infrastructure::database::async_postgres::get_connection;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::interceptor::auth::AuthLayer;
use crate::interceptor::logger::LoggerLayer;
use crate::interceptor::request_id::RequestIdLayer;
use crate::module::shipping::presenter_http::ShippingHttpPresenter;
use crate::module::shipping_provider::presenter_http::ShippingProviderHttpPresenter;
use crate::module::shipping_provider::repository_postgres::ShippingProviderPostgresRepositoryImpl;
use crate::module::shipping_provider::usecase::ShippingProviderUseCaseImpl;
use axum::http::StatusCode;
use axum::routing::get;
use axum::{Router, response::Json};
use tower::ServiceBuilder;

pub struct HttpTransport {
    config: AppConfig,
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
        let user_service = UserServiceGrpcClient::new(self.config.clone()).await;

        // Repository Layer
        let shipping_provider_repository =
            ShippingProviderPostgresRepositoryImpl::new(postgres_pool);

        let app = Router::new()
            .nest(
                "/v1/shipping/shipping_providers",
                ShippingProviderHttpPresenter::new(
                    ShippingProviderUseCaseImpl::new(shipping_provider_repository),
                    user_service.clone(),
                )
                .router(),
            )
            .nest(
                "/v1/shipping/shippings",
                ShippingHttpPresenter::new(user_service.clone()).router(),
            )
            .route("/v1/shipping/checks", get(health_check_handler))
            .layer(
                ServiceBuilder::new()
                    .layer(AuthLayer::new(user_service.clone()))
                    .layer(LoggerLayer)
                    .layer(RequestIdLayer),
            );

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
