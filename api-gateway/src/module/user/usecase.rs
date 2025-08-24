#[derive(Debug, Clone)]
pub struct UseCase {
    user_service_grpc: crate::module::user::transport_grpc::Transport,
    user_service_rabbitmq: crate::module::user::transport_rabbitmq::Transport,
}

impl UseCase {
    pub fn new(
        user_service_grpc: crate::module::user::transport_grpc::Transport,
        user_service_rabbitmq: crate::module::user::transport_rabbitmq::Transport,
    ) -> Self {
        Self {
            user_service_grpc,
            user_service_rabbitmq,
        }
    }
}
