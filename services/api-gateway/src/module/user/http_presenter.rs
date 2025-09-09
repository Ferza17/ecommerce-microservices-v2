use crate::interceptor::auth::AuthLayer;
use tower::ServiceBuilder;

#[derive(Debug, Clone)]
pub struct Presenter {
    user_use_case: crate::module::user::usecase::UseCase,
    auth_use_case: crate::module::auth::usecase::UseCase,
}
pub const ROUTE_PREFIX: &str = "/api/v1/users";
pub const TAG: &str = "USER";


impl Presenter {
    pub fn new(
        user_use_case: crate::module::user::usecase::UseCase,
        auth_use_case: crate::module::auth::usecase::UseCase,
    ) -> Self {
        Self {
            user_use_case,
            auth_use_case,
        }
    }

    pub fn router(&self) -> axum::Router {
        axum::Router::new()
            .layer(ServiceBuilder::new().layer(AuthLayer::new(self.auth_use_case.clone())))
            .with_state(self.clone())
    }
}
