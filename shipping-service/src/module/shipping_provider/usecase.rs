use crate::model::rpc::shipping::{
    CreateShippingProviderRequest, CreateShippingProviderResponse, DeleteShippingProviderRequest,
    DeleteShippingProviderResponse, GetShippingProviderByIdRequest,
    GetShippingProviderByIdResponse, ListShippingProvidersRequest, ListShippingProvidersResponse,
    UpdateShippingProviderRequest, UpdateShippingProviderResponse,
};
use crate::module::shipping_provider::repository_postgres::ShippingProviderPostgresRepository;
use tonic::{Request, Response, Status};
use tracing::{Level, event, instrument};

#[derive(Debug)]
pub struct ShippingProviderUseCase {
    shipping_provider_repository: ShippingProviderPostgresRepository,
}

impl ShippingProviderUseCase {
    pub fn new(shipping_provider_repository: ShippingProviderPostgresRepository) -> Self {
        Self {
            shipping_provider_repository,
        }
    }

    pub async fn create_shipping_provider(
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
    pub async fn get_shipping_provider_by_id(
        &self,
        request_id: String,
        request: Request<GetShippingProviderByIdRequest>,
    ) -> Result<Response<GetShippingProviderByIdResponse>, Status> {
        Ok(Response::new(GetShippingProviderByIdResponse {
            message: "Get Shipping Provider By Id ".to_string(),
            status: "success".to_string(),
            data: None,
        }))
    }

    pub async fn update_shipping_provider(
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

    pub async fn delete_shipping_provider(
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

    pub async fn list_shipping_providers(
        &self,
        request_id: String,
        request: Request<ListShippingProvidersRequest>,
    ) -> Result<Response<ListShippingProvidersResponse>, Status> {
        eprintln!("{:?}", request);
        eprintln!("{:?}", request_id);

        // TODO: Get shipping provider by id
        Ok(Response::new(ListShippingProvidersResponse {
            message: "".to_string(),
            status: "".to_string(),
            data: None,
        }))
    }
}
