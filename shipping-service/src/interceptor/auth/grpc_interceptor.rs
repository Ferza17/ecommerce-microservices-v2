use crate::interceptor;
use crate::model::rpc::user::AuthServiceVerifyIsExcludedRequest;
use crate::package::context::{auth::AUTHORIZATION_HEADER, request_id::get_request_id_from_header};
use axum::http::{Request, Response};
use futures::future::{Ready, ready};
use std::task::{Context, Poll};
use tokio::runtime::Handle;
use tonic::{Status, body::BoxBody};

impl<S> tower::Service<Request<BoxBody>> for interceptor::auth::AuthService<S>
where
    S: tower::Service<Request<BoxBody>, Response = Response<BoxBody>> + Send + 'static,
    S::Future: Send + 'static,
    S::Error: Into<Box<dyn std::error::Error + Send + Sync>>,
{
    type Response = Response<BoxBody>;
    type Error = S::Error;
    type Future = futures::future::Either<S::Future, Ready<Result<Self::Response, Self::Error>>>;

    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    fn call(&mut self, mut req: Request<BoxBody>) -> Self::Future {
        // VERIFY TOKEN
        let token_from_header = match req.headers().get(AUTHORIZATION_HEADER) {
            Some(val) => val.to_str().unwrap_or_default().to_string(),
            None => {
                return futures::future::Either::Right(ready(Ok(Status::unauthenticated(
                    "No authorization header",
                )
                .into_http())));
            }
        };

        if !token_from_header.starts_with("Bearer ") {
            return futures::future::Either::Right(ready(Ok(Status::unauthenticated(
                "Invalid authorization header",
            )
            .into_http())));
        }

        let token = token_from_header.trim_start_matches("Bearer ").trim();

        if token.is_empty() {
            return futures::future::Either::Right(ready(Ok(Status::unauthenticated(
                "missing token",
            )
            .into_http())));
        }

        req.headers_mut()
            .insert(AUTHORIZATION_HEADER, token.parse().unwrap());
        req.extensions_mut().insert(token.to_string());
        // Proceed to inner service if passed
        futures::future::Either::Left(self.inner.call(req))
    }
}
