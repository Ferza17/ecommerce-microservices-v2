use crate::interceptor::auth::AuthService;
use crate::model::rpc::{
    response::Response as CommonResponse, user::AuthServiceVerifyIsExcludedRequest,
};
use crate::package::context::{auth::AUTHORIZATION_HEADER, request_id::get_request_id_from_header};
use crate::util;
use axum::{
    body::Body,
    http::{Request, Response},
    response::Json,
};
use futures::future::{Either, Ready, ready};
use std::{
    convert::Infallible,
    task::{Context, Poll},
};
use tokio::runtime::Handle;

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
        // Check bearer token
        let token = match req.headers().get(AUTHORIZATION_HEADER) {
            Some(val) => {
                let token_str = match val.to_str() {
                    Ok(s) => s,
                    Err(_) => {
                        return Either::Right(ready(unauthorize_response(
                            tonic::Code::Unauthenticated,
                            "invalid authorization header",
                        )));
                    }
                };

                if !token_str.starts_with("Bearer ")
                    || token_str.trim_start_matches("Bearer ").trim().is_empty()
                {
                    return Either::Right(ready(unauthorize_response(
                        tonic::Code::Unauthenticated,
                        "invalid bearer token format",
                    )));
                }

                token_str.trim_start_matches("Bearer ").trim().to_string()
            }
            None => {
                return Either::Right(ready(unauthorize_response(
                    tonic::Code::Unauthenticated,
                    "missing authorization header",
                )));
            }
        };

        req.headers_mut()
            .insert(AUTHORIZATION_HEADER, token.parse().unwrap());
        req.extensions_mut().insert(token);
        Either::Left(self.inner.call(req))
    }
}

fn unauthorize_response(status: tonic::Code, message: &str) -> Result<Response<Body>, Infallible> {
    let json_response = Json(CommonResponse {
        status: "".to_string(),
        message: message.to_string(),
        data: None,
    });

    // Convert Json to Response<Body>
    let response = Response::builder()
        .status(util::convert_status::tonic_to_http_status(status))
        .header("content-type", "application/json")
        .body(Body::from(
            serde_json::to_vec(&json_response.0).unwrap_or_default(),
        ))
        .unwrap_or_else(|_| Response::new(Body::empty()));

    Ok(response)
}
