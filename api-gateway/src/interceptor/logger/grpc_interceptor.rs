use crate::interceptor::logger::LoggerService;
use crate::package::context::url_path::URL_PATH;
use axum::http::Request;
use std::task::{Context, Poll};
use tonic::body::BoxBody;
impl<S> tower::Service<Request<BoxBody>> for LoggerService<S>
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
        let method_path = req.uri().path().to_string();
        println!("gRPC method: {}", method_path);

        req.headers_mut()
            .insert(URL_PATH, method_path.parse().unwrap());
        req.extensions_mut().insert(method_path.clone());
        self.inner.call(req)
    }
}
