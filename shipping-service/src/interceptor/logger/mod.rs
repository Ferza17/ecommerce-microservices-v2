use tower::Layer;

pub mod grpc_interceptor;
pub mod http_interceptor;

#[derive(Clone)]
pub struct LoggerLayer;
impl<S> Layer<S> for LoggerLayer {
    type Service = LoggerService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        LoggerService { inner }
    }
}

#[derive(Clone)]
pub struct LoggerService<S> {
    pub inner: S,
}
