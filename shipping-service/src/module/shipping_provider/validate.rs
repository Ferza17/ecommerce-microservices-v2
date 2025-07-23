use crate::model::rpc::shipping::{
    CreateShippingProviderRequest, DeleteShippingProviderRequest, GetShippingProviderByIdRequest,
    ListShippingProvidersRequest, UpdateShippingProviderRequest,
};
use tonic::{Request, Status};

pub fn validate_get_shipping_provider_by_id(
    request: &Request<GetShippingProviderByIdRequest>,
) -> Option<Status> {
    if request.get_ref().id.is_empty() {
        return Some(Status::invalid_argument("id"));
    }
    None
}

pub fn validate_create_shipping_provider(
    request: &Request<CreateShippingProviderRequest>,
) -> Option<Status> {
    if request.get_ref().name.is_empty() {
        return Some(Status::invalid_argument("name"));
    }
    None
}

pub fn validate_update_shipping_provider(
    request: &Request<UpdateShippingProviderRequest>,
) -> Option<Status> {
    if request.get_ref().id.is_empty() {
        return Some(Status::invalid_argument("id"));
    }
    if request.get_ref().name.is_none() {
        return Some(Status::invalid_argument("name"));
    }
    None
}

pub fn validate_delete_shipping_provider(
    request: &Request<DeleteShippingProviderRequest>,
) -> Option<Status> {
    if request.get_ref().id.is_empty() {
        return Some(Status::invalid_argument("id"));
    }
    None
}

pub fn validate_list_shipping_providers(
    request: &Request<ListShippingProvidersRequest>,
) -> Option<Status> {
    if request.get_ref().page <= 0 {
        return Some(Status::invalid_argument("page"));
    }
    if request.get_ref().limit <= 0 {
        return Some(Status::invalid_argument("limit"));
    }
    None
}
