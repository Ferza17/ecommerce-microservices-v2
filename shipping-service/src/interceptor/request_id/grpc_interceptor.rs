use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use std::str::FromStr;
use tonic::metadata::MetadataKey;
use tonic::{Request, Status, metadata::MetadataValue, service::Interceptor};
use tracing::{error, info};
use uuid::Uuid;

#[derive(Clone)]
pub struct GrpcRequestIdInterceptor;
impl Interceptor for GrpcRequestIdInterceptor {
    fn call(&mut self, mut req: Request<()>) -> Result<Request<()>, Status> {
        let key: MetadataKey<_> = MetadataKey::from_static(X_REQUEST_ID_HEADER);
        let request_id = if let Some(existing) = req.metadata().get(X_REQUEST_ID_HEADER) {
            existing.to_str().unwrap_or_default().to_string()
        } else {
            Uuid::new_v4().to_string()
        };

        let val = MetadataValue::from_str(&request_id).map_err(|e| {
            error!("Invalid {} metadata: {}", X_REQUEST_ID_HEADER, e);
            Status::internal(format!("Invalid {} metadata: {}", X_REQUEST_ID_HEADER, e))
        })?;

        req.metadata_mut().insert(key, val);
        info!("{} : {}", X_REQUEST_ID_HEADER, request_id);
        Ok(req)
    }
}
