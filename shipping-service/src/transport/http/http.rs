use crate::config::config::AppConfig;
use crate::infrastructure::database::async_postgres::get_connection;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::module::shipping::presenter_http::ShippingHttpPresenter;
use crate::module::shipping::repository_postgres::ShippingPostgresRepositoryImpl;
use crate::module::shipping::usecase::ShippingUseCaseImpl;
use crate::module::shipping_provider::{
    presenter_http::ShippingProviderHttpPresenter,
    repository_postgres::ShippingProviderPostgresRepositoryImpl,
    usecase::ShippingProviderUseCaseImpl,
};
use crate::transport::http::api_docs::ApiDocs;
use axum::{Router, http::StatusCode, response::Json, routing::get};
use utoipa::OpenApi;
use utoipa_swagger_ui::SwaggerUi;

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

        let app = Router::new()
            .merge(
                SwaggerUi::new("/v1/shipping/docs")
                    .url("/api-docs/openapi.json", ApiDocs::openapi()),
            )
            .nest(
                "/v1/shipping/shipping_providers",
                ShippingProviderHttpPresenter::new(
                    ShippingProviderUseCaseImpl::new(ShippingProviderPostgresRepositoryImpl::new(
                        postgres_pool.clone(),
                    )),
                    user_service.clone(),
                )
                .router(),
            )
            .nest(
                "/v1/shipping/shippings",
                ShippingHttpPresenter::new(
                    ShippingUseCaseImpl::new(
                        ShippingPostgresRepositoryImpl::new(
                        postgres_pool.clone(),
                    )),
                    user_service.clone(),
                )
                .router(),
            )
            .route("/v1/shipping/checks", get(health_check_handler));

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
