use crate::module::user::usecase::UserUseCase;

pub mod grpc_interceptor;
pub mod http_interceptor;

#[derive(Clone)]
pub struct AuthLayer {
    pub user_use_case: std::sync::Arc<UserUseCase>,
}

impl AuthLayer {
    pub fn new(user_use_case: UserUseCase) -> Self {
        Self {
            user_use_case: std::sync::Arc::new(user_use_case),
        }
    }
}

#[derive(Clone, Debug)]
pub struct AuthService<S> {
    pub inner: S,
    pub user_use_case: UserUseCase,
}

impl<S> tower::Layer<S> for AuthLayer {
    type Service = AuthService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        AuthService {
            inner,
            user_use_case: (*self.user_use_case).clone(),
        }
    }
}
