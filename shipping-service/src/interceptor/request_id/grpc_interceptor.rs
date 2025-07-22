use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::grpc::extract_metadata_value;
use tonic::{Request, Status, service::Interceptor};
use uuid::Uuid;

#[derive(Clone)]
pub struct GrpcRequestIdInterceptor;
impl Interceptor for GrpcRequestIdInterceptor {
    fn call(&mut self, mut req: Request<()>) -> Result<Request<()>, Status> {
        let request_id =
            if let Some(id) = extract_metadata_value(req.metadata(), X_REQUEST_ID_HEADER) {
                id.to_string()
            } else {
                Uuid::new_v4().to_string()
            };

        req.metadata_mut()
            .insert(X_REQUEST_ID_HEADER, request_id.parse().unwrap());
        Ok(req)
    }
}
