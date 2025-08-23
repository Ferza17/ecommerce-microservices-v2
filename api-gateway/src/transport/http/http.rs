use crate::config::config::AppConfig;
use crate::interceptor::request_id::RequestIdLayer;
use crate::module::payment::http_presenter::PaymentPresenterHttp;
use crate::module::payment::usecase::PaymentUseCase;
use crate::module::product::usecase::ProductUseCase;
use crate::module::shipping::http_presenter::ShippingPresenterHttp;
use crate::module::shipping::transport_grpc::ShippingTransportGrpc;
use crate::module::shipping::usecase::ShippingUseCase;
use crate::module::user::{
    http_presenter::UserPresenterHttp, transport_grpc::UserTransportGrpc,
    transport_rabbitmq::UserTransportRabbitMQ, usecase::UserUseCase,
};
use crate::module::{
    payment::transport_grpc::PaymentTransportGrpc, product::http_presenter::ProductPresenterHttp,
    product::transport_grpc::ProductTransportGrpc,
};
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::package::context::traceparent::TRACEPARENT_HEADER;
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
use opentelemetry::global;
use tower_http::{
    cors::{Any, CorsLayer},
    trace::TraceLayer,
};
use tracing::info_span;
use tracing_opentelemetry::OpenTelemetrySpanExt;
use utoipa::OpenApi;
use utoipa_swagger_ui::SwaggerUi;
use crate::util::metadata::HeaderExtractor;

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
        let shipping_transport_grpc = ShippingTransportGrpc::new(self.config.clone()).await?;

        // Use case layer
        let user_use_case = UserUseCase::new(user_transport_grpc.clone(), user_transport_rabbitmq);
        let product_use_case = ProductUseCase::new(product_transport_grpc.clone());
        let shipping_use_case = ShippingUseCase::new(shipping_transport_grpc.clone());
        let payment_use_case = PaymentUseCase::new(
            payment_transport_grpc,
            shipping_transport_grpc,
            user_transport_grpc,
            product_transport_grpc,
        );

        // Presenter layer
        let user_presenter = UserPresenterHttp::new(user_use_case.clone());
        let product_presenter = ProductPresenterHttp::new(product_use_case, user_use_case.clone());
        let payment_presenter = PaymentPresenterHttp::new(payment_use_case, user_use_case.clone());
        let shipping_presenter = ShippingPresenterHttp::new(shipping_use_case, user_use_case);

        let app = Router::new()
            .nest("/api/v1/auth", user_presenter.clone().auth_router())
            .nest("/api/v1/users", user_presenter.user_router())
            .nest("/api/v1/products", product_presenter.router())
            .nest(
                "/api/v1/payments",
                payment_presenter.clone().payment_router(),
            )
            .nest(
                "/api/v1/payment-providers",
                payment_presenter.payment_provider_router(),
            )
            .nest(
                "/api/v1/shippings",
                shipping_presenter.clone().shipping_route(),
            )
            .nest(
                "/api/v1/shipping-providers",
                shipping_presenter.shipping_provider_route(),
            )
            .route("/api/v1/checks", get(health_check_handler))
            .merge(SwaggerUi::new("/api/v1/docs").url("/api-docs/openapi.json", ApiDocs::openapi()))
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
