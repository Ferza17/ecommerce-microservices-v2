use crate::model::rpc::shipping::shipping_provider_service_server::ShippingProviderService;
use crate::model::rpc::shipping::{
    CreateShippingProviderRequest, CreateShippingProviderResponse, DeleteShippingProviderRequest,
    DeleteShippingProviderResponse, GetShippingProviderByIdRequest,
    GetShippingProviderByIdResponse, ListShippingProvidersRequest, ListShippingProvidersResponse,
    UpdateShippingProviderRequest, UpdateShippingProviderResponse,
};
use crate::module::shipping_provider::usecase::ShippingProviderUseCase;
use crate::package::context::request_id::get_request_id_from_metadata;
use tonic::{Request, Response, Status};

pub struct ShippingProviderPresenter {
    shipping_provider_use_case: ShippingProviderUseCase,
}

impl ShippingProviderPresenter {
    pub fn new(shipping_provider_use_case: ShippingProviderUseCase) -> Self {
        Self {
            shipping_provider_use_case,
        }
    }

    // fn create_timestamp() -> Option<Timestamp> {
    //     let now = Utc::now();
    //     Some(Timestamp {
    //         seconds: now.timestamp(),
    //         nanos: now.timestamp_subsec_nanos() as i32,
    //     })
    // }
}

#[tonic::async_trait]
impl ShippingProviderService for ShippingProviderPresenter {
    async fn create_shipping_provider(
        &self,
        request: Request<CreateShippingProviderRequest>,
    ) -> Result<Response<CreateShippingProviderResponse>, Status> {
        self.shipping_provider_use_case
            .create_shipping_provider(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }

    async fn get_shipping_provider_by_id(
        &self,
        request: Request<GetShippingProviderByIdRequest>,
    ) -> Result<Response<GetShippingProviderByIdResponse>, Status> {
        self.shipping_provider_use_case
            .get_shipping_provider_by_id(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }

    async fn update_shipping_provider(
        &self,
        request: Request<UpdateShippingProviderRequest>,
    ) -> Result<Response<UpdateShippingProviderResponse>, Status> {
        self.shipping_provider_use_case
            .update_shipping_provider(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }

    async fn delete_shipping_provider(
        &self,
        request: Request<DeleteShippingProviderRequest>,
    ) -> Result<Response<DeleteShippingProviderResponse>, Status> {
        self.shipping_provider_use_case
            .delete_shipping_provider(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }

    async fn list_shipping_providers(
        &self,
        request: Request<ListShippingProvidersRequest>,
    ) -> Result<Response<ListShippingProvidersResponse>, Status> {
        self.shipping_provider_use_case
            .list_shipping_providers(get_request_id_from_metadata(request.metadata()), request)
            .await
            .map_err(|e| Status::internal(e.to_string()))
    }
}
