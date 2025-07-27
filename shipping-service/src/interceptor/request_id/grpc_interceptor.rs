use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use hyper::{Request, header::HeaderValue};
use std::task::{Context, Poll};
use tonic::body::BoxBody;
use tracing::error;
use uuid::Uuid;
use crate::interceptor::request_id::RequestIdService;

impl<S> tower::Service<Request<BoxBody>> for RequestIdService<S>
where
    S: tower::Service<Request<BoxBody>>,
{
    type Response = S::Response;
    type Error = S::Error;
    type Future = S::Future;

    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    fn call(&mut self, mut req: Request<BoxBody>) -> Self::Future {
        let request_id = match req.headers().get(X_REQUEST_ID_HEADER) {
            Some(val) => val.to_str().unwrap_or_default().to_string(),
            None => Uuid::new_v4().to_string(),
        };

        if req.headers().get(X_REQUEST_ID_HEADER).is_none() {
            if let Ok(value) = HeaderValue::from_str(&request_id) {
                req.headers_mut().insert(X_REQUEST_ID_HEADER, value);
            } else {
                error!("Failed to insert generated request ID into header.");
            }
        }

        // Optionally inject into extensions
        req.extensions_mut().insert(request_id.clone());
        eprintln!("{}: {}", X_REQUEST_ID_HEADER, request_id);
        self.inner.call(req)
    }
}
