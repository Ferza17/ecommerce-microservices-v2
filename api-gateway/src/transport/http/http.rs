use crate::config::config::AppConfig;
use crate::module::payment::http_presenter::PaymentPresenterHttp;
use crate::module::payment::usecase::PaymentUseCase;
use crate::module::product::usecase::ProductUseCase;
use crate::module::user::{
    http_presenter::UserPresenterHttp, transport_grpc::UserTransportGrpc,
    transport_rabbitmq::UserTransportRabbitMQ, usecase::UserUseCase,
};
use crate::module::{
    payment::transport_grpc::PaymentTransportGrpc, product::http_presenter::ProductPresenterHttp,
    product::transport_grpc::ProductTransportGrpc,
};
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::transport::http::api_docs::ApiDocs;
use axum::{
    Router,
    http::{
        HeaderValue, StatusCode,
        header::{AUTHORIZATION, CONTENT_TYPE},
    },
    response::Json,
    routing::get,
};
use tower_http::cors::{Any, CorsLayer};
use utoipa::OpenApi;
use utoipa_swagger_ui::SwaggerUi;

pub struct HttpTransport {
    config: AppConfig,
}

impl HttpTransport {
    pub fn new(config: AppConfig) -> Self {
        Self { config }
    }

    pub async fn serve(self) -> Result<(), anyhow::Error> {
        let addr = format!(
            "{}:{}",
            self.config.api_gateway_service_service_http_host,
            self.config.api_gateway_service_service_http_port
        )
        .to_string();

        // Transport Layer
        let user_transport_grpc = UserTransportGrpc::new(self.config.clone()).await?;
        let user_transport_rabbitmq = UserTransportRabbitMQ::new(self.config.clone());
        let product_transport_grpc = ProductTransportGrpc::new(self.config.clone()).await?;
        let payment_transport_grpc = PaymentTransportGrpc::new(self.config.clone()).await?;

        // Use case layer
        let user_use_case = UserUseCase::new(user_transport_grpc, user_transport_rabbitmq);
        let product_use_case = ProductUseCase::new(product_transport_grpc);
        let payment_use_case = PaymentUseCase::new(payment_transport_grpc);

        // Presenter layer
        let user_presenter = UserPresenterHttp::new(user_use_case.clone());
        let product_presenter = ProductPresenterHttp::new(product_use_case, user_use_case.clone());
        let payment_presenter = PaymentPresenterHttp::new(payment_use_case, user_use_case);

        let app = Router::new()
            // USER ROUTE
            .nest("/api/v1/auth", user_presenter.clone().auth_router())
            .nest("/api/v1/users", user_presenter.user_router())
            // PRODUCT ROUTE
            .nest("/api/v1/products", product_presenter.router())
            // PAYMENT ROUTE
            .nest(
                "/api/v1/payments",
                payment_presenter.clone().payment_router(),
            )
            .nest(
                "/api/v1/payment-providers",
                payment_presenter.payment_provider_router(),
            )
            .route("/api/v1/checks", get(health_check_handler))
            .merge(SwaggerUi::new("/api/v1/docs").url("/api-docs/openapi.json", ApiDocs::openapi()))
            .layer(
                CorsLayer::new()
                    .allow_methods(Any)
                    .allow_origin("*".parse::<HeaderValue>()?)
                    .allow_headers([AUTHORIZATION, CONTENT_TYPE, X_REQUEST_ID_HEADER.parse()?]),
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
            "service": "api-gateway-service",
        })),
    )
}
