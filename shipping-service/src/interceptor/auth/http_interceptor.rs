use crate::interceptor::auth::AuthService;
use crate::model::rpc::response::Response as CommonResponse;
use crate::model::rpc::user::AuthServiceVerifyIsExcludedRequest;
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::get_request_id_from_header;
use axum::body::Body;
use axum::http::{Request, Response};
use axum::response::Json;
use futures::future::ready;
use futures::future::{Either, Ready};
use std::convert::Infallible;
use std::task::{Context, Poll};
use tokio::runtime::Handle;
use tonic::Code;

impl<S> tower::Service<Request<Body>> for AuthService<S>
where
    S: tower::Service<Request<Body>, Error = Infallible, Response = Response<Body>>
        + Clone
        + Send
        + 'static,
    S::Response: Send + 'static,
    S::Future: Send + 'static,
{
    type Response = Response<Body>;
    type Error = Infallible;
    type Future = Either<S::Future, Ready<Result<Response<Body>, Infallible>>>;

    fn poll_ready(&mut self, cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    fn call(&mut self, mut req: Request<Body>) -> Self::Future {
        let request_id = get_request_id_from_header(req.headers());
        let path = req.uri().path().to_string();
        let mut cloned_user_service = self.user_service.clone();

        let verify_result = tokio::task::block_in_place(|| {
            Handle::current().block_on(async move {
                cloned_user_service
                    .auth_service_verify_is_excluded(
                        request_id.clone(),
                        AuthServiceVerifyIsExcludedRequest {
                            full_method_name: Some(path.clone()),
                            http_url: None,
                            http_method: None,
                        },
                    )
                    .await
            })
        });

        match verify_result {
            Ok(res) => {
                let Some(data) = res.data else {
                    return Either::Right(ready(unauthorize_response("no data in response")));
                };
                if data.is_excluded {
                    return Either::Left(self.inner.call(req));
                }
            }
            Err(_err) => {
                return Either::Right(ready(unauthorize_response("failed to verify exclusion")));
            }
        }

        // Check bearer token
        let token = match req.headers().get(AUTHORIZATION_HEADER) {
            Some(val) => {
                let token_str = match val.to_str() {
                    Ok(s) => s,
                    Err(_) => {
                        return Either::Right(ready(unauthorize_response(
                            "invalid authorization header",
                        )));
                    }
                };

                if !token_str.starts_with("Bearer ")
                    || token_str.trim_start_matches("Bearer ").trim().is_empty()
                {
                    return Either::Right(ready(unauthorize_response(
                        "invalid bearer token format",
                    )));
                }

                token_str.trim_start_matches("Bearer ").trim().to_string()
            }
            None => {
                return Either::Right(ready(unauthorize_response("missing authorization header")));
            }
        };

        req.headers_mut()
            .insert(AUTHORIZATION_HEADER, token.parse().unwrap());
        req.extensions_mut().insert(token);
        Either::Left(self.inner.call(req))
    }
}

fn unauthorize_response(message: &str) -> Result<Response<Body>, Infallible> {
    let json_response = Json(CommonResponse {
        error: "unauthorized".to_string(),
        message: message.to_string(),
        code: Code::Unauthenticated as i32,
        data: None,
    });

    // Convert Json to Response<Body>
    let response = Response::builder()
        .status(401)
        .header("content-type", "application/json")
        .body(Body::from(
            serde_json::to_vec(&json_response.0).unwrap_or_default(),
        ))
        .unwrap_or_else(|_| Response::new(Body::empty()));

    Ok(response)
}
