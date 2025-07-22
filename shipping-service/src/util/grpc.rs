use tonic::metadata::MetadataMap;

pub fn extract_metadata_value<'a>(metadata: &'a MetadataMap, key: &str) -> Option<&'a str> {
    metadata.get(key).and_then(|val| val.to_str().ok())
}
