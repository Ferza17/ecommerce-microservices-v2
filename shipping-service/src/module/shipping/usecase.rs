use crate::model::rpc::shipping::{
    CreateShippingRequest, CreateShippingResponse, DeleteShippingRequest, DeleteShippingResponse,
    GetShippingByIdRequest, GetShippingByIdResponse, ListShippingRequest, ListShippingResponse,
    UpdateShippingRequest, UpdateShippingResponse,
};
use crate::module::shipping::repository_postgres::{
    ShippingPostgresRepository, ShippingPostgresRepositoryImpl,
};
use tonic::{Request, Response, Status};
use tracing::instrument;

pub trait ShippingUseCase {
    async fn create_shipping(
        &self,
        request_id: String,
        request: Request<CreateShippingRequest>,
    ) -> Result<Response<CreateShippingResponse>, Status>;
    async fn get_shipping_by_id(
        &self,
        request_id: String,
        request: Request<GetShippingByIdRequest>,
    ) -> Result<Response<GetShippingByIdResponse>, Status>;
    async fn list_shipping(
        &self,
        request_id: String,
        request: Request<ListShippingRequest>,
    ) -> Result<Response<ListShippingResponse>, Status>;
    async fn update_shipping(
        &self,
        request_id: String,
        request: Request<UpdateShippingRequest>,
    ) -> Result<Response<UpdateShippingResponse>, Status>;
    async fn delete_shipping(
        &self,
        request_id: String,
        request: Request<DeleteShippingRequest>,
    ) -> Result<Response<DeleteShippingResponse>, Status>;
}

#[derive(Clone, Debug)]
pub struct ShippingUseCaseImpl {
    shipping_repository: ShippingPostgresRepositoryImpl,
}

impl ShippingUseCaseImpl {
    pub fn new(shipping_repository: ShippingPostgresRepositoryImpl) -> Self {
        Self {
            shipping_repository,
        }
    }
}

impl ShippingUseCase for ShippingUseCaseImpl {
    #[instrument("ShippingUseCase.create_shipping")]
    async fn create_shipping(
        &self,
        request_id: String,
        request: Request<CreateShippingRequest>,
    ) -> Result<Response<CreateShippingResponse>, Status> {
        todo!()
    }

    #[instrument("ShippingUseCase.get_shipping_by_id")]
    async fn get_shipping_by_id(
        &self,
        request_id: String,
        request: Request<GetShippingByIdRequest>,
    ) -> Result<Response<GetShippingByIdResponse>, Status> {
        todo!()
    }

    #[instrument("ShippingUseCase.list_shipping")]
    async fn list_shipping(
        &self,
        request_id: String,
        request: Request<ListShippingRequest>,
    ) -> Result<Response<ListShippingResponse>, Status> {
        todo!()
    }

    #[instrument("ShippingUseCase.update_shipping")]
    async fn update_shipping(
        &self,
        request_id: String,
        request: Request<UpdateShippingRequest>,
    ) -> Result<Response<UpdateShippingResponse>, Status> {
        todo!()
    }

    #[instrument("ShippingUseCase.delete_shipping")]
    async fn delete_shipping(
        &self,
        request_id: String,
        request: Request<DeleteShippingRequest>,
    ) -> Result<Response<DeleteShippingResponse>, Status> {
        todo!()
    }
}
