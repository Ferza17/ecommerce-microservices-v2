use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::interceptor::auth::AuthLayer;
use crate::interceptor::logger::LoggerLayer;
use crate::interceptor::request_id::RequestIdLayer;
use tower::ServiceBuilder;

#[derive(Debug, Clone)]
pub struct ShippingHttpPresenter {
    user_service: UserServiceGrpcClient,
}

impl ShippingHttpPresenter {
    pub fn new(user_service: UserServiceGrpcClient) -> Self {
        Self { user_service }
    }

    pub fn router(&self) -> axum::Router {
        axum::Router::new().layer(
            ServiceBuilder::new()
                .layer(RequestIdLayer)
                .layer(LoggerLayer)
                .layer(AuthLayer::new(self.user_service.clone())),
        )
    }
}
