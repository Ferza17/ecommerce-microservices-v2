use tonic::metadata::MetadataMap;

pub const X_REQUEST_ID_HEADER: &str = "X-Request-Id";

pub fn get_request_id_from_metadata(metadata: &MetadataMap) -> String {
    metadata
        .get(X_REQUEST_ID_HEADER)
        .unwrap()
        .to_str()
        .unwrap()
        .to_string()
}
