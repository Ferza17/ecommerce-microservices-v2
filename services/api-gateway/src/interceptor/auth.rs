use crate::model::rpc::user::{AuthUserFindUserByTokenRequest, User};
use crate::package::context::request_id::get_request_id_from_header;
use axum::extract::OriginalUri;
use std::task::{Context, Poll};

#[derive(Clone, Debug)]
pub struct AuthLayer {
    auth_use_case: crate::module::auth::usecase::UseCase,
}

impl AuthLayer {
    pub fn new(auth_use_case: crate::module::auth::usecase::UseCase) -> Self {
        Self { auth_use_case }
    }
}

impl<S> tower::Layer<S> for AuthLayer {
    type Service = AuthLayerService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        AuthLayerService {
            inner,
            auth_use_case: self.auth_use_case.clone(),
        }
    }
}

#[derive(Clone, Debug)]
pub struct AuthLayerService<S> {
    pub inner: S,
    pub auth_use_case: crate::module::auth::usecase::UseCase,
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

    #[tracing::instrument("AuthLayer.poll_ready")]
    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    #[tracing::instrument("AuthLayer.call")]
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
            token.clone().parse().unwrap(),
        );

        // Validate Expiration Token ON USER SERVICE
        let request_id = get_request_id_from_header(req.headers());

        let inner_user = self.auth_use_case.clone();
        let user = match tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async move {
                let cloned_token = token.clone();
                inner_user
                    .clone()
                    .auth_user_find_user_by_token(
                        request_id.clone(),
                        cloned_token.clone(),
                        tonic::Request::new(AuthUserFindUserByTokenRequest {
                            token: cloned_token.clone(),
                        }),
                    )
                    .await
            })
        }) {
            Ok(val) => {
                if val.data.is_none() {
                    return futures::future::Either::Right(futures::future::ready(
                        unauthorize_response(tonic::Code::Unauthenticated, "invalid bearer token"),
                    ));
                }

                val.data.unwrap().user.unwrap()
            }
            Err(err) => {
                return futures::future::Either::Right(futures::future::ready(
                    unauthorize_response(err.code(), err.message()),
                ));
            }
        };

        // VALIDATE ACCESS ON OPA
        let inner_user_validate = self.auth_use_case.clone();
        let path = match req.extensions().get::<OriginalUri>() {
            None => "".to_string(),
            Some(original_uri) => original_uri.path().to_string(),
        };

        let method = req.method().to_string().to_uppercase();
        match tokio::task::block_in_place(|| {
            tokio::runtime::Handle::current().block_on(async move {
                inner_user_validate
                    .clone()
                    .auth_validate_access(method.clone(), path, User::from(user.clone()))
                    .await
            })
        }) {
            Ok(_) => {}
            Err(err) => {
                return futures::future::Either::Right(futures::future::ready(
                    unauthorize_response(err.code(), err.message()),
                ));
            }
        };

        futures::future::Either::Left(self.inner.call(req))
    }
}
