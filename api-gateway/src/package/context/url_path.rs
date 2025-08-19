// use axum::http::{HeaderMap, HeaderValue};
use tonic::metadata::MetadataMap;

pub const URL_PATH: &str = "URL_PATH";

pub fn get_url_path_from_metadata(metadata: &MetadataMap) -> String {
    metadata
        .get(URL_PATH)
        .unwrap()
        .to_str()
        .unwrap()
        .to_string()
}