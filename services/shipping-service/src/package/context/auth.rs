use axum::http::{HeaderMap, HeaderValue};
use tonic::metadata::MetadataMap;

pub const AUTHORIZATION_HEADER: &str = "authorization";

pub fn get_request_authorization_token_from_metadata(metadata: &MetadataMap) -> String {
    metadata
        .get(AUTHORIZATION_HEADER)
        .unwrap()
        .to_str()
        .unwrap()
        .to_string()
}

pub fn get_request_authorization_token_from_header(header: &HeaderMap<HeaderValue>) -> String {
    header
        .get(AUTHORIZATION_HEADER)
        .unwrap()
        .to_str()
        .expect("Invalid header value")
        .to_string()
}