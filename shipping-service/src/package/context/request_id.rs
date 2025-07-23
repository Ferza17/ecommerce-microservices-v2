use axum::http::{HeaderMap, HeaderValue};
use tonic::metadata::MetadataMap;

pub const X_REQUEST_ID_HEADER: &str = "x-request-id";

pub fn get_request_id_from_metadata(metadata: &MetadataMap) -> String {
    metadata
        .get(X_REQUEST_ID_HEADER)
        .unwrap()
        .to_str()
        .unwrap()
        .to_string()
}

pub fn get_request_id_from_header(header: &HeaderMap<HeaderValue>) -> String {
    header
        .get(X_REQUEST_ID_HEADER)
        .unwrap()
        .to_str()
        .unwrap()
        .to_string()
}
