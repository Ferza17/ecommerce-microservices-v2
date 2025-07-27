use crate::interceptor::request_id::RequestIdService;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use axum::{
    body::Body,
    http::{HeaderName, HeaderValue, Request},
};
use std::convert::Infallible;
use std::task::{Context, Poll};
use uuid::Uuid;

impl<S> tower::Service<Request<Body>> for RequestIdService<S>
where
    S: tower::Service<Request<Body>, Error = Infallible>,
{
    type Response = S::Response;
    type Error = S::Error;
    type Future = S::Future;

    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    fn call(&mut self, mut req: Request<Body>) -> Self::Future {
        // Extract or generate request ID
        let request_id = req
            .headers()
            .get(HeaderName::from_static(X_REQUEST_ID_HEADER))
            .and_then(|val| val.to_str().ok())
            .map(|s| s.to_string())
            .unwrap_or_else(|| Uuid::new_v4().to_string());
        // Store into extensions for handler access
        
        
        eprintln!("Request ID: {}", request_id);
        
        req.headers_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.clone().parse().unwrap());
        req.extensions_mut().insert(request_id.clone());
        self.inner.call(req)
    }
}
