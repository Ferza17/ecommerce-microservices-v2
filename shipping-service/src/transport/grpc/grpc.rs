use crate::config::config::AppConfig;
use crate::infrastructure::database::postgres::PostgresInfrastructure;
use crate::interceptor::request_id::grpc_interceptor::GrpcRequestIdInterceptor;
use crate::model::rpc::shipping::shipping_provider_service_server::ShippingProviderServiceServer;
use crate::module::shipping_provider::presenter::ShippingProviderPresenter;
use crate::module::shipping_provider::repository_postgres::ShippingProviderPostgresRepository;
use crate::module::shipping_provider::usecase::ShippingProviderUseCase;
use tonic::service::interceptor;
use tonic::transport::Server;

pub struct GrpcTransport {
    config: AppConfig,
}

impl GrpcTransport {
    pub fn new(cfg: AppConfig) -> Self {
        GrpcTransport { config: cfg }
    }

    pub async fn serve(&self) -> Result<(), Box<dyn std::error::Error>> {
        // tracing_subscriber::init();

        let addr = format!(
            "{}:{}",
            self.config.shipping_service_service_rpc_host,
            self.config.shipping_service_service_rpc_port
        )
        .to_string();

        eprintln!("Starting gRPC server on {}", addr);

        // Infrastructure Layer
        let postgres_infrastructure = PostgresInfrastructure::new(self.config.clone())
            .await
            .unwrap();

        // Repository Layer
        let shipping_provider_repository =
            ShippingProviderPostgresRepository::new(postgres_infrastructure);

        Server::builder()
            .layer(interceptor(GrpcRequestIdInterceptor))
            .add_service(ShippingProviderServiceServer::new(
                ShippingProviderPresenter::new(ShippingProviderUseCase::new(
                    shipping_provider_repository.clone(),
                )),
            ))
            .serve(addr.parse()?)
            .await?;

        Ok(())
    }
}
