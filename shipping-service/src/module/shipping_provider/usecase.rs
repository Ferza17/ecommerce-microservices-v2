use crate::model::diesel::orm::shipping_provider::{
    shipping_provider_to_proto, shipping_providers_to_proto,
};
use crate::model::rpc::shipping::list_shipping_providers_response::ListShippingProvidersResponseData;
use crate::model::rpc::shipping::{
    CreateShippingProviderRequest, CreateShippingProviderResponse, DeleteShippingProviderRequest,
    DeleteShippingProviderResponse, GetShippingProviderByIdRequest,
    GetShippingProviderByIdResponse, ListShippingProvidersRequest, ListShippingProvidersResponse,
    UpdateShippingProviderRequest, UpdateShippingProviderResponse,
};
use crate::module::shipping_provider::repository_postgres::{
    ShippingProviderPostgresRepository, ShippingProviderPostgresRepositoryImpl,
};
use tonic::{Request, Response, Status};
use tracing::{Level, event, instrument};

pub trait ShippingProviderUseCase {
    async fn create_shipping_provider(
        &self,
        request_id: String,
        request: Request<CreateShippingProviderRequest>,
    ) -> Result<Response<CreateShippingProviderResponse>, Status>;
    async fn get_shipping_provider_by_id(
        &self,
        request_id: String,
        request: Request<GetShippingProviderByIdRequest>,
    ) -> Result<Response<GetShippingProviderByIdResponse>, Status>;
    async fn update_shipping_provider(
        &self,
        request_id: &String,
        request: Request<UpdateShippingProviderRequest>,
    ) -> Result<Response<UpdateShippingProviderResponse>, Status>;
    async fn delete_shipping_provider(
        &self,
        request_id: String,
        request: Request<DeleteShippingProviderRequest>,
    ) -> Result<Response<DeleteShippingProviderResponse>, Status>;
    async fn list_shipping_providers(
        &self,
        request_id: String,
        request: Request<ListShippingProvidersRequest>,
    ) -> Result<Response<ListShippingProvidersResponse>, Status>;
}

#[derive(Debug, Clone)]
pub struct ShippingProviderUseCaseImpl {
    shipping_provider_repository: ShippingProviderPostgresRepositoryImpl,
}

impl ShippingProviderUseCaseImpl {
    pub fn new(shipping_provider_repository: ShippingProviderPostgresRepositoryImpl) -> Self {
        Self {
            shipping_provider_repository,
        }
    }
}

impl ShippingProviderUseCase for ShippingProviderUseCaseImpl {
    #[instrument("ShippingProviderUseCase.create_shipping_provider")]
    async fn create_shipping_provider(
        &self,
        request_id: String,
        request: Request<CreateShippingProviderRequest>,
    ) -> Result<Response<CreateShippingProviderResponse>, Status> {
        eprintln!("{:?}", request);
        eprintln!("{:?}", request_id);

        // TODO: Get shipping provider by id
        Ok(Response::new(CreateShippingProviderResponse {
            message: "".to_string(),
            status: "".to_string(),
            data: None,
        }))
    }

    #[instrument("ShippingProviderUseCase.get_shipping_provider_by_id")]
    async fn get_shipping_provider_by_id(
        &self,
        request_id: String,
        request: Request<GetShippingProviderByIdRequest>,
    ) -> Result<Response<GetShippingProviderByIdResponse>, Status> {
        event!(name: "ShippingProviderUseCase.get_shipping_provider_by_id", Level::INFO, request_id = request_id, request = ?request);

        let shipping_provider = self
            .shipping_provider_repository
            .get_shipping_provider_by_id(request_id.as_str(), request.into_inner().id.as_str())
            .await;

        match shipping_provider {
            Ok(shipping_provider) => Ok(Response::new(GetShippingProviderByIdResponse {
                message: "Get Shipping Provider By Id".to_string(),
                status: "success".to_string(),
                data: Option::from(shipping_provider_to_proto(shipping_provider)),
            })),

            Err(err) => {
                event!(
                    Level::ERROR,
                    request_id = request_id,
                    error = %err,
                    "Failed to get shipping provider by ID"
                );

                Err(Status::internal(format!(
                    "Failed to get shipping provider: {}",
                    err
                )))
            }
        }
    }

    #[instrument("ShippingProviderUseCase.update_shipping_provider")]
    async fn update_shipping_provider(
        &self,
        request_id: &String,
        request: Request<UpdateShippingProviderRequest>,
    ) -> Result<Response<UpdateShippingProviderResponse>, Status> {
        eprintln!("{:?}", request);
        eprintln!("{:?}", request_id);

        // TODO: Get shipping provider by id
        Ok(Response::new(UpdateShippingProviderResponse {
            message: "".to_string(),
            status: "".to_string(),
            data: None,
        }))
    }

    #[instrument("ShippingProviderUseCase.delete_shipping_provider")]
    async fn delete_shipping_provider(
        &self,
        request_id: String,
        request: Request<DeleteShippingProviderRequest>,
    ) -> Result<Response<DeleteShippingProviderResponse>, Status> {
        eprintln!("{:?}", request);
        eprintln!("{:?}", request_id);

        // TODO: Get shipping provider by id
        Ok(Response::new(DeleteShippingProviderResponse {
            message: "".to_string(),
            status: "".to_string(),
        }))
    }

    #[instrument("ShippingProviderUseCase.list_shipping_providers")]
    async fn list_shipping_providers(
        &self,
        request_id: String,
        request: Request<ListShippingProvidersRequest>,
    ) -> Result<Response<ListShippingProvidersResponse>, Status> {
        let request_inner = request.into_inner();
        let shipping_providers = self
            .shipping_provider_repository
            .list_shipping_providers(
                request_id.as_str(),
                &request_inner.page,
                &request_inner.limit,
            )
            .await;

        match shipping_providers {
            Ok(providers) => Ok(Response::new(ListShippingProvidersResponse {
                message: "List Shipping Providers".to_string(),
                status: "success".to_string(),
                data: Option::from(ListShippingProvidersResponseData {
                    shipping_providers: shipping_providers_to_proto(providers),
                    total_count: 0, // TODO: Get total count
                    page: request_inner.page,
                    limit: request_inner.limit,
                }),
            })),
            Err(err) => Err(Status::internal(format!(
                "Failed to list shipping providers: {}",
                err
            ))),
        }
    }
}
