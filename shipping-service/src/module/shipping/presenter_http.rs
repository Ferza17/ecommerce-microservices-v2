use crate::infrastructure::services::user::UserServiceGrpcClient;
#[derive(Debug, Clone)]
pub struct ShippingHttpPresenter {
    user_service: UserServiceGrpcClient,
}

impl ShippingHttpPresenter {
    pub fn new(user_service: UserServiceGrpcClient) -> Self {
        Self { user_service }
    }

    pub fn router(&self) -> axum::Router {
        axum::Router::new()
    }
}
