use tower::Layer;

pub mod grpc_interceptor;
pub mod http_interceptor;


#[derive(Clone)]
pub struct RequestIdLayer;
impl<S> Layer<S> for RequestIdLayer {
    type Service = RequestIdService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        RequestIdService { inner }
    }
}

#[derive(Clone)]
pub struct RequestIdService<S> {
    pub inner: S,
}