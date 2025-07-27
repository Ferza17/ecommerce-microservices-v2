use crate::interceptor::logger::LoggerService;
use axum::body::Body;
use axum::http::Request;
use std::convert::Infallible;
use std::task::{Context, Poll};
impl<S> tower::Service<Request<Body>> for LoggerService<S>
where
    S: tower::Service<Request<Body>, Error = Infallible>,
{
    type Response = S::Response;
    type Error = S::Error;
    type Future = S::Future;

    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    fn call(&mut self, req: Request<Body>) -> Self::Future {
        println!(
            "HTTP method: {}  path: {}",
            req.method().to_string(),
            req.uri().path()
        );

        self.inner.call(req)
    }
}
