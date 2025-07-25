use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use axum::{
    body::Body,
    http::{HeaderName, HeaderValue, Request},
};
use std::convert::Infallible;
use std::task::{Context, Poll};
use tonic::codegen::Service;
use uuid::Uuid;
use crate::interceptor::request_id::RequestIdService;

impl<S> Service<Request<Body>> for RequestIdService<S>
where
    S: Service<Request<Body>, Error = Infallible>,
{
    type Response = S::Response;
    type Error = S::Error;
    type Future = S::Future;

    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    fn call(&mut self, mut req: Request<Body>) -> Self::Future {
        let header_name = HeaderName::from_static(X_REQUEST_ID_HEADER);

        // Extract or generate request ID
        let request_id = match req.headers().get(&header_name) {
            Some(val) => val.to_str().unwrap_or("").to_string(),
            None => {
                let new_id = Uuid::new_v4().to_string();
                if let Ok(hv) = HeaderValue::from_str(&new_id) {
                    req.headers_mut().insert(header_name.clone(), hv);
                }
                new_id
            }
        };

        // Store into extensions for handler access
        req.extensions_mut().insert(request_id.clone());

        eprintln!("{} is: {}", X_REQUEST_ID_HEADER, request_id);

        self.inner.call(req)
    }
}
