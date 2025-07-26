use crate::interceptor::auth::AuthService;
use axum::body::Body;
use axum::http::Request;
use std::convert::Infallible;
use std::task::{Context, Poll};
use tonic::codegen::Service;

impl<S> Service<Request<Body>> for AuthService<S>
where
    S: Service<Request<Body>, Error = Infallible>,
{
    type Response = S::Response;
    type Error = S::Error;
    type Future = S::Future;

    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    fn call(&mut self, req: Request<Body>) -> Self::Future {
        self.inner.call(req)
    }
}
