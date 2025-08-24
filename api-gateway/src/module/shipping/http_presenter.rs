use crate::interceptor::auth::AuthLayer;
use tower::ServiceBuilder;
use tracing::instrument;
#[derive(Debug, Clone)]
pub struct Presenter {
    shipping_use_case: crate::module::shipping::usecase::UseCase,
    auth_use_case: crate::module::auth::usecase::UseCase,
}
pub const ROUTE_PREFIX: &'static str = "/api/v1/shipping-providers";
pub const TAG: &str = "Shipping";

impl Presenter {
    pub fn new(
        shipping_use_case: crate::module::shipping::usecase::UseCase,
        auth_use_case: crate::module::auth::usecase::UseCase,
    ) -> Self {
        Self {
            shipping_use_case,
            auth_use_case,
        }
    }

    #[instrument]
    pub fn router(&self) -> axum::Router {
        axum::Router::new()
            .layer(ServiceBuilder::new().layer(AuthLayer::new(self.auth_use_case.clone())))
            .with_state(self.clone())
    }
}
