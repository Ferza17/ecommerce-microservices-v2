#[derive(Debug, Clone)]
pub struct UseCase {
    user_service_grpc: crate::module::user::transport_grpc::Transport,
    user_transport_kafka: crate::module::user::transport_kafka::Transport,
}

impl UseCase {
    pub fn new(
        user_service_grpc: crate::module::user::transport_grpc::Transport,
        user_transport_kafka: crate::module::user::transport_kafka::Transport,
    ) -> Self {
        Self {
            user_service_grpc,
            user_transport_kafka,
        }
    }
}
