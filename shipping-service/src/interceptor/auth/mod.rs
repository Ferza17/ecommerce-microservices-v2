use crate::infrastructure::services::user::UserServiceGrpcClient;

pub mod grpc_interceptor;
pub mod http_interceptor;

#[derive(Clone)]
pub struct AuthLayer {
    pub user_service: std::sync::Arc<UserServiceGrpcClient>,
}

impl AuthLayer {
    pub fn new(user_service: UserServiceGrpcClient) -> Self {
        Self {
            user_service: std::sync::Arc::new(user_service),
        }
    }
}

#[derive(Clone, Debug)]
pub struct AuthService<S> {
    pub inner: S,
    pub user_service: UserServiceGrpcClient,
}

impl<S> tower::Layer<S> for AuthLayer {
    type Service = AuthService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        AuthService {
            inner,
            user_service: (*self.user_service).clone(),
        }
    }
}
