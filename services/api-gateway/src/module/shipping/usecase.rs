#[derive(Debug, Clone)]
pub struct UseCase {
    shipping_transport_grpc: crate::module::shipping::transport_grpc::Transport,
}

impl UseCase {
    pub fn new(
        shipping_transport_grpc: crate::module::shipping::transport_grpc::Transport,
    ) -> Self {
        Self {
            shipping_transport_grpc,
        }
    }
}
