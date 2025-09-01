use axum::http::{HeaderMap, HeaderValue};
pub const X_REQUEST_ID_HEADER: &str = "x-request-id";
pub fn get_request_id_from_header(header: &HeaderMap<HeaderValue>) -> String {
    header
        .get(X_REQUEST_ID_HEADER)
        .unwrap()
        .to_str()
        .expect("Invalid header value")
        .to_string()
}
