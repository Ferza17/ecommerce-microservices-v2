use crate::config::config::AppConfig;
use crate::infrastructure::database::async_postgres::get_connection;
use crate::infrastructure::services::payment::PaymentServiceGrpcClient;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::module::shipping::presenter_http::PresenterHttp as ShippingHttpPresenter;
use crate::module::shipping::repository_postgres::ShippingPostgresRepositoryImpl;
use crate::module::shipping::usecase::ShippingUseCaseImpl;
use crate::module::shipping_provider::{
    presenter_http::PresenterHttp as ShippingProviderHttpPresenter,
    repository_postgres::ShippingProviderPostgresRepositoryImpl,
    usecase::ShippingProviderUseCaseImpl,
};
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::package::context::traceparent::TRACEPARENT_HEADER;
use crate::transport::http::api_docs::ApiDocs;
use crate::util::metadata::HeaderExtractor;
use axum::http::HeaderValue;
use axum::http::header::{AUTHORIZATION, CONTENT_TYPE};
use axum::{Router, http::StatusCode, response::Json, routing::get};
use opentelemetry::global;
use tower_http::cors::{Any, CorsLayer};
use tower_http::trace::TraceLayer;
use tracing::info_span;
use tracing_opentelemetry::OpenTelemetrySpanExt;
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
            self.config.service_shipping.http_host, self.config.service_shipping.http_port,
        )
        .to_string();

        // Infrastructure Layer
        let postgres_pool = get_connection(&self.config.clone()).await;
        let user_service = UserServiceGrpcClient::new(self.config.clone()).await;
        let payment_service = PaymentServiceGrpcClient::new(self.config.clone()).await;

        // Repository Layer
        let shipping_provider_postgres_repository =
            ShippingProviderPostgresRepositoryImpl::new(postgres_pool.clone());
        let shipping_postgres_repository =
            ShippingPostgresRepositoryImpl::new(postgres_pool.clone());

        // UseCase Layer
        let shipping_provider_use_case =
            ShippingProviderUseCaseImpl::new(shipping_provider_postgres_repository.clone());
        let shipping_use_case = ShippingUseCaseImpl::new(
            shipping_postgres_repository.clone(),
            shipping_provider_postgres_repository.clone(),
            user_service.clone(),
            payment_service.clone(),
        );

        // Presenter Layer
        let shipping_provider_presenter =
            ShippingProviderHttpPresenter::new(shipping_provider_use_case, user_service.clone());
        let shipping_presenter =
            ShippingHttpPresenter::new(shipping_use_case.clone(), user_service.clone());

        let app = Router::new()
            .merge(
                SwaggerUi::new("/v1/shipping/docs")
                    .url("/api-docs/openapi.json", ApiDocs::openapi()),
            )
            .nest(
                "/v1/shipping/shipping_providers",
                shipping_provider_presenter.router(),
            )
            .nest("/v1/shipping/shippings", shipping_presenter.router())
            .route("/v1/shipping/check", get(health_check_handler))
            .layer(
                CorsLayer::new()
                    .allow_methods(Any)
                    .allow_origin("*".parse::<HeaderValue>()?)
                    .allow_headers([
                        AUTHORIZATION,
                        CONTENT_TYPE,
                        X_REQUEST_ID_HEADER.parse()?,
                        TRACEPARENT_HEADER.parse()?,
                    ]),
            )
            .layer(
                TraceLayer::new_for_http()
                    .make_span_with(|request: &hyper::Request<_>| {
                        let span = info_span!(
                            "HTTP REQUEST",
                            method = ?request.method(),
                            path = %request.uri().path(),
                        );
                        span.set_parent(global::get_text_map_propagator(|prop| {
                            prop.extract(&HeaderExtractor(request.headers()))
                        }));
                        span
                    })
                    .on_request(|request: &hyper::Request<_>, _span: &tracing::Span| {
                        tracing::info!("started {} {}", request.method(), request.uri().path());
                    })
                    .on_response(
                        |response: &hyper::Response<_>,
                         latency: std::time::Duration,
                         _span: &tracing::Span| {
                            tracing::info!("response {} in {:?}", response.status(), latency);
                        },
                    ),
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
