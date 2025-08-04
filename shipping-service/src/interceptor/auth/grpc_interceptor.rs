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
        let request_id = get_request_id_from_header(req.headers());
        // VERIFY ACL ON METHOD
        let verify_result = tokio::task::block_in_place(|| {
            Handle::current().block_on(async {
                self.user_service
                    .clone()
                    .auth_service_verify_is_excluded(
                        request_id.clone(),
                        tonic::Request::new(AuthServiceVerifyIsExcludedRequest {
                            full_method_name: Some(req.uri().path().to_string()),
                            http_url: None,
                            http_method: None,
                        }),
                    )
                    .await
            })
        });
        match verify_result {
            Ok(res) => {
                let Some(data) = res.data else {
                    return futures::future::Either::Right(ready(Ok(Status::unauthenticated(
                        "No data in response",
                    )
                    .into_http())));
                };

                if data.is_excluded {
                    return futures::future::Either::Left(self.inner.call(req));
                }
            }
            Err(err) => {
                tracing::error!(%request_id, "Auth check failed: {err}");
                return futures::future::Either::Right(ready(Ok(Status::unauthenticated(
                    "Failed to verify exclusion",
                )
                .into_http())));
            }
        }

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
