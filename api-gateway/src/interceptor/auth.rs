use crate::model::rpc::user::AuthUserFindUserByTokenResponse;
use crate::module::user::usecase::UserUseCase;
use crate::package::context::request_id::get_request_id_from_header;
use std::task::{Context, Poll};
use tokio::runtime::Runtime;
use tonic::Status;
use tracing::instrument;

#[derive(Clone, Debug)]
pub struct AuthLayer {
    user_use_case: UserUseCase,
}

impl AuthLayer {
    pub fn new(user_use_case: UserUseCase) -> Self {
        Self { user_use_case }
    }
}

impl<S> tower::Layer<S> for AuthLayer {
    type Service = AuthLayerService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        AuthLayerService {
            inner,
            user_use_case: self.user_use_case.clone(),
        }
    }
}

#[derive(Clone, Debug)]
pub struct AuthLayerService<S> {
    pub inner: S,
    pub user_use_case: UserUseCase,
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

        // Validate Expiration Token ON USER SERVICE
       

        req.headers_mut().insert(
            crate::package::context::auth::AUTHORIZATION_HEADER,
            token.parse().unwrap(),
        );
        futures::future::Either::Left(self.inner.call(req))
    }
}
