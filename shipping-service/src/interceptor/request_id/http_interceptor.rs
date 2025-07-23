use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use axum::http::{HeaderName, HeaderValue};
use axum::{extract::Request, middleware::Next, response::Response};
use tracing::info;
use uuid::Uuid;

pub async fn x_request_id_middleware(mut request: Request, next: Next) -> Response {
    let header_name = HeaderName::from_static(X_REQUEST_ID_HEADER);

    // Check if a header exists
    let request_id = match request.headers().get(&header_name) {
        Some(value) => value.to_str().unwrap_or("").to_string(),
        None => {
            let new_id = Uuid::new_v4().to_string();

            // Insert it into headers so that it's carried to downstream
            request
                .headers_mut()
                .insert(header_name.clone(), HeaderValue::from_str(&new_id).unwrap());
            new_id
        }
    };

    // Insert into request extensions so handler can use it
    request.extensions_mut().insert(request_id.clone());

    info!("{} is: {}", X_REQUEST_ID_HEADER, request_id);

    let mut response = next.run(request).await;

    // Add to response headers (optional but useful for tracing)
    response
        .headers_mut()
        .insert(header_name, HeaderValue::from_str(&request_id).unwrap());

    response
}
