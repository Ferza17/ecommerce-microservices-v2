use crate::interceptor::auth::AuthLayerService;
use std::task::{Context, Poll};
use tracing::instrument;
use uuid::Uuid;

#[derive(Clone, Debug)]
pub struct RequestIdLayer;
impl<S> tower::Layer<S> for RequestIdLayer {
    type Service = RequestIdLayerService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        RequestIdLayerService { inner }
    }
}

#[derive(Clone, Debug)]
pub struct RequestIdLayerService<S> {
    pub inner: S,
}

impl<S> tower::Service<axum::http::Request<axum::body::Body>> for RequestIdLayerService<S>
where
    S: tower::Service<
            axum::http::Request<axum::body::Body>,
            Error = std::convert::Infallible,
            Response = axum::http::Response<axum::body::Body>,
        > + Clone
        + Send
        + std::fmt::Debug
        + Sync
        + 'static,
{
    type Response = axum::http::Response<axum::body::Body>;
    type Error = S::Error;
    type Future = S::Future;

    #[instrument("RequestIdLayer.poll_ready")]
    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    #[instrument("RequestIdLayer.call")]
    fn call(&mut self, mut req: axum::http::Request<axum::body::Body>) -> Self::Future {
        let request_id = req
            .headers()
            .get(crate::package::context::request_id::X_REQUEST_ID_HEADER)
            .and_then(|val| val.to_str().ok())
            .map(|s| s.to_string())
            .unwrap_or_else(|| Uuid::new_v4().to_string());
        req.headers_mut().insert(
            crate::package::context::request_id::X_REQUEST_ID_HEADER,
            request_id.clone().parse().unwrap(),
        );
        self.inner.call(req)
    }
}
