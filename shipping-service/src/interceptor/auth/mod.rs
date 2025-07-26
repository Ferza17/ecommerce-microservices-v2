use crate::infrastructure::services::user::UserServiceGrpcClient;
use std::sync::Arc;
use tower::Layer;

pub mod grpc_interceptor;
pub mod http_interceptor;

#[derive(Clone)]
pub struct AuthLayer {
    pub user_service: Arc<UserServiceGrpcClient>,
}

impl AuthLayer {
    pub fn new(user_service: UserServiceGrpcClient) -> Self {
        Self {
            user_service: Arc::new(user_service),
        }
    }
}

#[derive(Clone, Debug)]
pub struct AuthService<S> {
    pub inner: S,
    pub user_service: UserServiceGrpcClient,
}

impl<S> Layer<S> for AuthLayer {
    type Service = AuthService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        AuthService {
            inner,
            user_service: (*self.user_service).clone(),
        }
    }
}
