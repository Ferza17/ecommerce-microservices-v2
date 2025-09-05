use crate::config::config::AppConfig;
use crate::interceptor::request_id::RequestIdLayer;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::package::context::traceparent::TRACEPARENT_HEADER;
use crate::transport::http::api_docs::ApiDocs;
use crate::util::metadata::HeaderExtractor;
use axum::{
    Router,
    http::{
        HeaderValue, StatusCode,
        header::{AUTHORIZATION, CONTENT_TYPE},
    },
    response::Json,
    routing::get,
};
use opentelemetry::global;
use tower_http::{
    cors::{Any, CorsLayer},
    trace::TraceLayer,
};
use tracing::{info_span, span};
use tracing_opentelemetry::OpenTelemetrySpanExt;
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
            self.config.service_api_gateway.http_host, self.config.service_api_gateway.http_port,
        )
        .to_string();

        // Infrastructure Layer
        let opa = crate::infrastructure::opa::opa::OPA::new(self.config.clone());
        let rabbitmq =
            crate::infrastructure::message_broker::rabbitmq::RabbitMQInfrastructure::new(
                self.config.clone(),
            )
            .await;

        // Transport Layer
        let user_transport_grpc =
            crate::module::user::transport_grpc::Transport::new(self.config.clone()).await?;
        let auth_transport_grpc =
            crate::module::auth::transport_grpc::Transport::new(self.config.clone()).await?;
        let user_transport_rabbitmq =
            crate::module::user::transport_rabbitmq::Transport::new(self.config.clone(), rabbitmq);

        let product_transport_grpc =
            crate::module::product::transport_grpc::Transport::new(self.config.clone()).await?;

        let payment_provider_transport_grpc =
            crate::module::payment_providers::transport_grpc::Transport::new(self.config.clone())
                .await?;
        let payment_transport_grpc =
            crate::module::payment::transport_grpc::Transport::new(self.config.clone()).await?;

        let shipping_provider_transport_grpc =
            crate::module::shipping_provider::transport_grpc::Transport::new(self.config.clone())
                .await?;
        let shipping_transport_grpc =
            crate::module::shipping::transport_grpc::Transport::new(self.config.clone()).await?;
        let event_transport_grpc =
            crate::module::event::transport_grpc::Transport::new(self.config.clone()).await?;

        // Use case layer
        let auth_use_case = crate::module::auth::usecase::UseCase::new(
            auth_transport_grpc,
            user_transport_grpc.clone(),
            event_transport_grpc,
            user_transport_rabbitmq.clone(),
            opa,
        );
        let user_use_case = crate::module::user::usecase::UseCase::new(
            user_transport_grpc.clone(),
            user_transport_rabbitmq,
        );
        let product_use_case =
            crate::module::product::usecase::UseCase::new(product_transport_grpc.clone());
        let shipping_provider_use_case = crate::module::shipping_provider::usecase::UseCase::new(
            shipping_provider_transport_grpc.clone(),
        );
        let shipping_use_case =
            crate::module::shipping::usecase::UseCase::new(shipping_transport_grpc.clone());
        let payment_use_case = crate::module::payment::usecase::UseCase::new(
            payment_transport_grpc.clone(),
            payment_provider_transport_grpc.clone(),
            shipping_provider_transport_grpc,
            user_transport_grpc,
            product_transport_grpc,
        );
        let payment_provider_use_case = crate::module::payment_providers::usecase::UseCase::new(
            payment_provider_transport_grpc,
        );
        let notification_use_case = crate::module::notification::usecase::UseCase::new();

        // Presenter layer
        let auth_presenter =
            crate::module::auth::http_presenter::Presenter::new(auth_use_case.clone());
        let user_presenter = crate::module::user::http_presenter::Presenter::new(
            user_use_case.clone(),
            auth_use_case.clone(),
        );
        let product_presenter = crate::module::product::http_presenter::Presenter::new(
            product_use_case,
            auth_use_case.clone(),
        );
        let payment_provider_presenter =
            crate::module::payment_providers::http_presenter::Presenter::new(
                payment_provider_use_case,
                auth_use_case.clone(),
            );
        let payment_presenter = crate::module::payment::http_presenter::Presenter::new(
            payment_use_case,
            auth_use_case.clone(),
        );
        let shipping_provider_presenter =
            crate::module::shipping_provider::http_presenter::Presenter::new(
                shipping_provider_use_case,
                auth_use_case.clone(),
            );
        let shipping_presenter = crate::module::shipping::http_presenter::Presenter::new(
            shipping_use_case,
            auth_use_case,
        );
        let notification_presenter =
            crate::module::notification::http_presenter::Presenter::new(notification_use_case);

        let app = Router::new()
            .nest(
                crate::module::auth::http_presenter::ROUTE_PREFIX,
                auth_presenter.router(),
            )
            .nest(
                crate::module::user::http_presenter::ROUTE_PREFIX,
                user_presenter.router(),
            )
            .nest(
                crate::module::product::http_presenter::ROUTE_PREFIX,
                product_presenter.router(),
            )
            .nest(
                crate::module::payment::http_presenter::ROUTE_PREFIX,
                payment_presenter.clone().router(),
            )
            .nest(
                crate::module::payment_providers::http_presenter::ROUTE_PREFIX,
                payment_provider_presenter.router(),
            )
            .nest(
                crate::module::shipping_provider::http_presenter::ROUTE_PREFIX,
                shipping_provider_presenter.router(),
            )
            .nest(
                crate::module::shipping::http_presenter::ROUTE_PREFIX,
                shipping_presenter.router(),
            )
            .nest(
                crate::module::notification::http_presenter::ROUTE_PREFIX,
                notification_presenter.router(),
            )
            .route("/api/v1/checks", get(health_check_handler))
            .merge(
                SwaggerUi::new(crate::transport::http::api_docs::ROUTE_PREFIX)
                    .url("/api-docs/openapi.json", ApiDocs::openapi()),
            )
            .layer(
                TraceLayer::new_for_http()
                    .make_span_with(|request: &hyper::Request<_>| {
                        let span = info_span!(
                            "HTTP REQUEST", // <-- convert String -> &str
                            method = %request.method(),
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
            )
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
            .layer(RequestIdLayer);

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
