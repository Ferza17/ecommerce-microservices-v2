use std::task::{Context, Poll};
use tracing::instrument;

#[derive(Clone, Debug)]
pub struct AuthLayer;
impl<S> tower::Layer<S> for AuthLayer {
    type Service = AuthLayerService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        AuthLayerService { inner }
    }
}

#[derive(Clone, Debug)]
pub struct AuthLayerService<S> {
    pub inner: S,
}

// HTTP
impl<S> tower::Service<axum::http::Request<axum::body::Body>> for AuthLayerService<S>
where
    S: tower::Service<
            axum::http::Request<axum::body::Body>,
            Error = std::convert::Infallible,
            Response = axum::http::Response<axum::body::Body>,
        > + Clone
        + Send
        + std::fmt::Debug
        + 'static,
    S::Response: Send + 'static,
    S::Future: Send + 'static,
{
    type Response = axum::http::Response<axum::body::Body>;
    type Error = std::convert::Infallible;
    type Future = futures::future::Either<
        S::Future,
        futures::future::Ready<
            Result<axum::http::Response<axum::body::Body>, std::convert::Infallible>,
        >,
    >;

    #[instrument("AuthLayer.poll_ready")]
    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    #[instrument("AuthLayer.call")]

    fn call(&mut self, mut req: axum::http::Request<axum::body::Body>) -> Self::Future {
        fn unauthorize_response(
            status: tonic::Code,
            message: &str,
        ) -> Result<axum::http::Response<axum::body::Body>, std::convert::Infallible> {
            let json_response = axum::response::Json(crate::model::rpc::response::Response {
                status: "unauthorized".to_string(),
                message: message.to_string(),
                data: None,
            });

            // Convert Json to Response<Body>
            let response = axum::http::Response::builder()
                .status(crate::util::convert_status::tonic_to_http_status(status))
                .header("content-type", "application/json")
                .body(axum::body::Body::from(
                    serde_json::to_vec(&json_response.0).unwrap_or_default(),
                ))
                .unwrap_or_else(|_| axum::http::Response::new(axum::body::Body::empty()));

            Ok(response)
        }

        let token = match req
            .headers()
            .get(crate::package::context::auth::AUTHORIZATION_HEADER)
        {
            Some(val) => {
                let token_str = match val.to_str() {
                    Ok(s) => s,
                    Err(_) => {
                        return futures::future::Either::Right(futures::future::ready(
                            unauthorize_response(
                                tonic::Code::Unauthenticated,
                                "invalid authorization header",
                            ),
                        ));
                    }
                };

                if !token_str.starts_with("Bearer ")
                    || token_str.trim_start_matches("Bearer ").trim().is_empty()
                {
                    return futures::future::Either::Right(futures::future::ready(
                        unauthorize_response(
                            tonic::Code::Unauthenticated,
                            "invalid bearer token format",
                        ),
                    ));
                }

                token_str.trim_start_matches("Bearer ").trim().to_string()
            }
            None => {
                return futures::future::Either::Right(futures::future::ready(
                    unauthorize_response(
                        tonic::Code::Unauthenticated,
                        "missing authorization header",
                    ),
                ));
            }
        };

        if token.is_empty() {
            return futures::future::Either::Right(futures::future::ready(unauthorize_response(
                tonic::Code::Unauthenticated,
                "invalid bearer token format",
            )));
        }

        req.headers_mut().insert(
            crate::package::context::auth::AUTHORIZATION_HEADER,
            token.parse().unwrap(),
        );
        futures::future::Either::Left(self.inner.call(req))
    }
}

// GRPC
impl<S> tower::Service<hyper::Request<tonic::body::BoxBody>> for AuthLayerService<S>
where
    S: tower::Service<
            hyper::Request<tonic::body::BoxBody>,
            Response = hyper::Response<tonic::body::BoxBody>,
        > + std::fmt::Debug,
    S::Future: Send + 'static,
    S::Error: Into<Box<dyn std::error::Error + Send + Sync>>,
{
    type Response = hyper::Response<tonic::body::BoxBody>;
    type Error = S::Error;
    type Future = futures::future::Either<
        S::Future,
        futures::future::Ready<Result<Self::Response, Self::Error>>,
    >;

    #[instrument("AuthLayer.poll_ready")]
    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    #[instrument("AuthLayer.call")]
    fn call(&mut self, mut req: hyper::Request<tonic::body::BoxBody>) -> Self::Future {
        // VERIFY TOKEN
        let token_from_header = match req
            .headers()
            .get(crate::package::context::auth::AUTHORIZATION_HEADER)
        {
            Some(val) => val.to_str().unwrap_or_default().to_string(),
            None => {
                return futures::future::Either::Right(futures::future::ready(Ok(
                    tonic::Status::unauthenticated("No authorization header").into_http(),
                )));
            }
        };

        if !token_from_header.starts_with("Bearer ") {
            return futures::future::Either::Right(futures::future::ready(Ok(
                tonic::Status::unauthenticated("Invalid authorization header").into_http(),
            )));
        }

        let token = token_from_header.trim_start_matches("Bearer ").trim();
        if token.is_empty() {
            return futures::future::Either::Right(futures::future::ready(Ok(
                tonic::Status::unauthenticated("missing token").into_http(),
            )));
        }

        req.headers_mut().insert(
            crate::package::context::auth::AUTHORIZATION_HEADER,
            token.parse().unwrap(),
        );
        futures::future::Either::Left(self.inner.call(req))
    }
}
