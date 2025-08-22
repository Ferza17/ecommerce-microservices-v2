use crate::config::config::AppConfig;
use crate::infrastructure::database::async_postgres::get_connection;
use crate::infrastructure::services::payment::PaymentServiceGrpcClient;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::interceptor::auth::AuthLayer;
use crate::interceptor::request_id::RequestIdLayer;
use crate::model::rpc::shipping::shipping_provider_service_server::ShippingProviderServiceServer;
use crate::model::rpc::shipping::shipping_service_server::ShippingServiceServer;
use crate::module::shipping::presenter_grpc::ShippingGrpcPresenter;
use crate::module::shipping::repository_postgres::ShippingPostgresRepositoryImpl;
use crate::module::shipping::usecase::ShippingUseCaseImpl;
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
        // Infrastructure Layer
        let postgres_pool = get_connection(&self.config.clone()).await;
        let user_service = UserServiceGrpcClient::new(self.config.clone()).await;
        let payment_service = PaymentServiceGrpcClient::new(self.config.clone()).await;

        // Repository Layer
        let shipping_provider_postgres_repository =
            ShippingProviderPostgresRepositoryImpl::new(postgres_pool.clone());
        let shipping_postgres_repository =
            ShippingPostgresRepositoryImpl::new(postgres_pool.clone());

        // UseCase Layer
        let shipping_provider_use_case =
            ShippingProviderUseCaseImpl::new(shipping_provider_postgres_repository.clone());
        let shipping_use_case = ShippingUseCaseImpl::new(
            shipping_postgres_repository.clone(),
            shipping_provider_postgres_repository.clone(),
            user_service.clone(),
            payment_service.clone(),
        );

        // Presenter Layer
        let shipping_provider_presenter =
            ShippingProviderGrpcPresenter::new(shipping_provider_use_case, user_service.clone());
        let shipping_presenter =
            ShippingGrpcPresenter::new(shipping_use_case.clone(), user_service.clone());

        // REFLECTION
        let reflection_service = tonic_reflection::server::Builder::configure()
            .register_encoded_file_descriptor_set(include_bytes!("../../../descriptor.bin"))
            .build_v1alpha()
            .unwrap();

        let addr = format!(
            "{}:{}",
            self.config.shipping_service_service_rpc_host,
            self.config.shipping_service_service_rpc_port
        )
        .parse()?;
        Server::builder()
            .layer(AuthLayer)
            .add_service(ShippingProviderServiceServer::new(
                shipping_provider_presenter,
            ))
            .add_service(ShippingServiceServer::new(shipping_presenter))
            .add_service(reflection_service)
            .serve(addr)
            .await?;

        Ok(())
    }
}
