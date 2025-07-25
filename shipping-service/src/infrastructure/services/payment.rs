use crate::config::config::AppConfig;
use crate::infrastructure::services::PaymentService;
use crate::model::rpc::payment::payment_service_client::PaymentServiceClient;
use crate::model::rpc::payment::{FindPaymentByIdRequest, Payment};
use tonic::Status;
use tonic::transport::Channel;
use tracing::info;

pub struct PaymentServiceGrpcClient {
    config: AppConfig,
    payment_service_client: PaymentServiceClient<Channel>,
}

impl PaymentServiceGrpcClient {
    async fn new(config: AppConfig) -> Self {
        let channel = Channel::from_shared(
            format!(
                "http://{}:{}",
                config.payment_service_service_rpc_host, config.payment_service_service_rpc_port
            )
            .to_string(),
        )
        .expect("Failed to connect to user service")
        .connect()
        .await;

        let channel = match channel {
            Ok(chan) => {
                info!(
                    "connected to {} addr {}:{} ",
                    config.payment_service_service_name,
                    config.payment_service_service_rpc_host,
                    config.payment_service_service_rpc_port
                );
                chan
            }
            Err(e) => panic!("Failed to connect to user service: {}", e),
        };

        Self {
            config,
            payment_service_client: PaymentServiceClient::new(channel),
        }
    }
}

impl PaymentService for PaymentServiceGrpcClient {}
