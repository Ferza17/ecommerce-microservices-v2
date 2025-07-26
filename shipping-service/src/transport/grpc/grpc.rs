use crate::config::config::AppConfig;
use crate::infrastructure::database::async_postgres::get_connection;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::interceptor::auth::AuthLayer;
use crate::interceptor::logger::LoggerLayer;
use crate::interceptor::request_id::RequestIdLayer;
use crate::model::rpc::shipping::shipping_provider_service_server::ShippingProviderServiceServer;
use crate::module::shipping_provider::presenter_grpc::ShippingProviderGrpcPresenter;
use crate::module::shipping_provider::repository_postgres::ShippingProviderPostgresRepositoryImpl;
use crate::module::shipping_provider::usecase::ShippingProviderUseCaseImpl;
use tonic::transport::Server;
use tower::ServiceBuilder;

pub struct GrpcTransport {
    config: AppConfig,
}

impl GrpcTransport {
    pub fn new(cfg: AppConfig) -> Self {
        GrpcTransport { config: cfg }
    }

    pub async fn serve(&self) -> std::result::Result<(), Box<dyn std::error::Error>> {
        let addr = format!(
            "{}:{}",
            self.config.shipping_service_service_rpc_host,
            self.config.shipping_service_service_rpc_port
        )
        .to_string();

        eprintln!("GRPC Server is running on {}", addr);

        // Infrastructure Layer
        let postgres_pool = get_connection(&self.config.clone()).await;
        let user_service = UserServiceGrpcClient::new(self.config.clone()).await;

        // Repository Layer
        let shipping_provider_repository =
            ShippingProviderPostgresRepositoryImpl::new(postgres_pool);

        // REFLECTION
        let reflection_service = tonic_reflection::server::Builder::configure()
            .register_encoded_file_descriptor_set(include_bytes!("../../../descriptor.bin"))
            .build_v1alpha()
            .unwrap();

        let mut server = Server::builder()
            .layer(
                ServiceBuilder::new()
                    .layer(LoggerLayer)
                    .layer(RequestIdLayer)
                    .layer(AuthLayer::new(user_service.clone())),
            )
            .add_service(ShippingProviderServiceServer::new(
                ShippingProviderGrpcPresenter::new(
                    ShippingProviderUseCaseImpl::new(shipping_provider_repository),
                    user_service.clone(),
                ),
            ));

        if self.config.env != "production" {
            server = server.add_service(reflection_service);
        }

        server.serve(addr.parse()?).await?;
        Ok(())
    }
}
